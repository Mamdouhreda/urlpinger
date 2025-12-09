package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	loadevent "urlpinger/load"
)

func SubmitMultiURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Make response streaming-friendly (Server-Sent Events)
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Limit to 5 concurrent workers
	sem := make(chan struct{}, 5)

	// Channel for goroutines to send results
	resultChan := make(chan map[string]interface{})

	// Read URLs
	raw := r.FormValue("multi-url")
	if raw == "" {
		http.Error(w, "At least one URL is required", http.StatusBadRequest)
		return
	}

	lines := strings.Split(raw, "\n")
	var urls []string
	for _, line := range lines {
		u := strings.TrimSpace(line)
		if u != "" {
			urls = append(urls, u)
			if len(urls) == 10 {
				break
			}
		}
	}

	if len(urls) == 0 {
		http.Error(w, "No valid URLs provided", http.StatusBadRequest)
		return
	}

	// Launch goroutines
	for _, url := range urls {
		sem <- struct{}{} // acquire worker slot

		go func(url string) {
			defer func() { <-sem }() // release worker slot

			navTiming, err := loadevent.LoadEventMS(url)

			if err != nil {
				resultChan <- map[string]interface{}{
					"url":   url,
					"error": err.Error(),
				}
				return
			}
			//print in the terminal the url 
			fmt.Printf("[multi] fetched metrics for %s\n", url)
			
			resultChan <- map[string]interface{}{
				"url":        url,
				"loadEvent":  navTiming.LoadEvent,
				"ttfb":       navTiming.TTFB,
				"dns":        navTiming.DNS,
				"tls":        navTiming.TLS,
				"slowImages": navTiming.SlowImages,
			}
		}(url)
	}

	// Read results as soon as they finish and send as SSE
	for i := 0; i < len(urls); i++ {
		result := <-resultChan
		jsonData, _ := json.Marshal(result)
		fmt.Fprintf(w, "data: %s\n\n", jsonData)
		if flusher, ok := w.(http.Flusher); ok {
			flusher.Flush()
		}
	}
}
