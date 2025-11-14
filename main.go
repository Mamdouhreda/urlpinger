package main

import (
	"fmt"
	"sync"
	"time"
	"urlpinger/data"
	loadevent "urlpinger/load"
)

const MaxWorkers = 5 // Maximum number of concurrent workers
func main() {
    start := time.Now()
	var wg sync.WaitGroup
    sem := make(chan struct{}, MaxWorkers) // create the channel with the limit of the workers
    for _, site := range data.Sites {
       wg.Add(1)
       sem <- struct{}{} // sending the token 
        url := site.URL
        id := site.ID

        go func(url string, id int) {
            defer wg.Done()
            defer func() { <-sem }() // Release the token 
            navTiming, err := loadevent.LoadEventMS(url)
            if err != nil {
                fmt.Printf("Error fetching metrics for %s: %v\n", url, err)
                return
            }

            fmt.Println("URL:", url)
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