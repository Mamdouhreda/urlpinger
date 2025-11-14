package main

import (
	"fmt"
	"sync"
	"time"
	"urlpinger/data"
	loadevent "urlpinger/load"
)

func main() {
    start := time.Now()
	var wg sync.WaitGroup

    for _, site := range data.Sites {
        wg.Add(1)

        // capture values
        url := site.URL
        id := site.ID

        go func(url string, id int) {
            defer wg.Done()

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