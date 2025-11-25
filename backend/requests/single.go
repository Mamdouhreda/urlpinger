package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	loadevent "urlpinger/load"
)

// SubmitSingleURL handles post request for single url submission
func SubmitSingleURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	url := r.FormValue("single-url")
	if url == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	// Log the received URL
	fmt.Printf("[single] received URL: %s\n", url)

	// Get performance metrics using the load event function
	navTiming, err := loadevent.LoadEventMS(url)
	if err != nil {
		fmt.Printf("Error fetching metrics for %s: %v\n", url, err)
		http.Error(w, fmt.Sprintf("Error fetching metrics: %v", err), http.StatusInternalServerError)
		return
	}

	// Print results to console
	// fmt.Println("URL:", url)
	// fmt.Printf("Slow Images: %v\n", navTiming.SlowImages)
	// fmt.Println("LoadEvent:", navTiming.LoadEvent, "seconds")
	// fmt.Println("TTFB:", navTiming.TTFB, "milliseconds")
	// fmt.Println("DNS:", navTiming.DNS, "milliseconds")
	// fmt.Println("TLS:", navTiming.TLS, "milliseconds")
	// fmt.Println()

	// Send JSON response back to client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"url":        url,
		"loadEvent":  navTiming.LoadEvent,
		"ttfb":       navTiming.TTFB,
		"dns":        navTiming.DNS,
		"tls":        navTiming.TLS,
		"slowImages": navTiming.SlowImages,
	})
}