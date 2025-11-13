package main

import (
	"fmt"
	"sync"
	"urlpinger/data"
	loadevent "urlpinger/load"
)

func main() {
	var wg sync.WaitGroup
    for _, site := range data.Sites {
        wg.Add(1)
        go func(url string) {
            defer wg.Done()
        navTiming, err := loadevent.LoadEventMS(site.URL)
        if err != nil {
            fmt.Printf("Error fetching metrics for %s: %v\n", url, err)
            return
        }
        fmt.Println("URL:", site.URL)
        fmt.Println("LoadEvent:", navTiming.LoadEvent, "seconds")
        fmt.Println("TTFB:", navTiming.TTFB, "milliseconds")
        fmt.Println("DNS:", navTiming.DNS, "milliseconds")
        fmt.Println("TLS:", navTiming.TLS, "milliseconds")
        fmt.Println("ID:", site.ID)
    }(site.URL)
}
    wg.Wait()
}