package main

import (
	"fmt"
	"net/http"
	"urlpinger/requests"
)

const MaxWorkers = 5 // Maximum number of concurrent workers



func main() {
	// register the http handler for root, index.html, and form submission
	http.HandleFunc("/", requests.Home)
	http.HandleFunc("/submit-single-url", requests.SubmitSingleURL)

	// start the server
	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("HTTP server error:", err)
	}
    // var wg sync.WaitGroup
    // sem := make(chan struct{}, MaxWorkers) // create the channel with the limit of the workers
    // for _, site := range data.Sites {
    //     wg.Add(1)
    //     sem <- struct{}{} // sending the token
    //     url := site.URL
    //     id := site.ID
    //     // start the goroutine and send the url and id to the function
    //     go func(url string, id int) {
    //         defer wg.Done()
    //         defer func() { <-sem }() // Release the token
    //         navTiming, err := loadevent.LoadEventMS(url)
    //         if err != nil {
    //             fmt.Printf("Error fetching metrics for %s: %v\n", url, err)
    //             return
    //         }

    //         fmt.Println("URL:", url)
    //         fmt.Printf("Slow Images: %v\n", navTiming.SlowImages)
    //         fmt.Println("LoadEvent:", navTiming.LoadEvent, "seconds")
    //         fmt.Println("TTFB:", navTiming.TTFB, "milliseconds")
    //         fmt.Println("DNS:", navTiming.DNS, "milliseconds")
    //         fmt.Println("TLS:", navTiming.TLS, "milliseconds")
    //         fmt.Println("ID:", id)
    //         fmt.Println()
    //     }(url, id)
    // }
    // wg.Wait()
}