package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
	"time"
	"urlpinger/data"
	loadevent "urlpinger/load"
)

const MaxWorkers = 5 // Maximum number of concurrent workers

// home handles the root route.
func home(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "index.html")
}

// renderTemplate parses and executes the named template from the frontend directory.
func renderTemplate(w http.ResponseWriter, tmpl string) {
    // Get the path relative to the project root (go up one level from backend)
    templatePath := filepath.Join("..", "frontend", tmpl)
    t, err := template.ParseFiles(templatePath)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    // register the http handler for root and index.html
    http.HandleFunc("/", home)

    // start the server
    go func() {
        fmt.Println("Starting server at :8080")
        err := http.ListenAndServe(":8080", nil)
        if err != nil {
            fmt.Println("HTTP server error:", err)
        }
    }()
    start := time.Now()
    var wg sync.WaitGroup
    sem := make(chan struct{}, MaxWorkers) // create the channel with the limit of the workers
    for _, site := range data.Sites {
        wg.Add(1)
        sem <- struct{}{} // sending the token
        url := site.URL
        id := site.ID
        // start the goroutine and send the url and id to the function
        go func(url string, id int) {
            defer wg.Done()
            defer func() { <-sem }() // Release the token
            navTiming, err := loadevent.LoadEventMS(url)
            if err != nil {
                fmt.Printf("Error fetching metrics for %s: %v\n", url, err)
                return
            }

            fmt.Println("URL:", url)
            fmt.Printf("Slow Images: %v\n", navTiming.SlowImages)
            fmt.Println("LoadEvent:", navTiming.LoadEvent, "seconds")
            fmt.Println("TTFB:", navTiming.TTFB, "milliseconds")
            fmt.Println("DNS:", navTiming.DNS, "milliseconds")
            fmt.Println("TLS:", navTiming.TLS, "milliseconds")
            fmt.Println("ID:", id)
            fmt.Println()
        }(url, id)
    }
    wg.Wait()
    end := time.Now()
    fmt.Println("==== FINISHED ====")
    fmt.Printf("Time: %.2f seconds\n", end.Sub(start).Seconds())
}