package main

import (
	"fmt"

	"urlpinger/data"
	loadevent "urlpinger/load"
)

func main() {
    for _, site := range data.Sites {
        navTiming, err := loadevent.LoadEventMS(site.URL)
        if err != nil {
            fmt.Printf("Error fetching metrics for %s: %v\n", site.URL, err)
            continue
        }

        fmt.Println("URL:", site.URL)
        fmt.Println("LoadEvent:", navTiming.LoadEvent, "seconds")
        fmt.Println("TTFB:", navTiming.TTFB, "milliseconds")
        fmt.Println("DNS:", navTiming.DNS, "milliseconds")
        fmt.Println("TLS:", navTiming.TLS, "milliseconds")
        fmt.Println()
    }
}
