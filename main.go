package main

import (
    "fmt"
    "urlpinger/data"
    "urlpinger/load"
)


func main() {
 
	NavTiming, err := loadevent.LoadEventMS(data.First.URL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("LoadEvent:", NavTiming.LoadEvent, "seconds")
	fmt.Println("TTFB:", NavTiming.TTFB, "milliseconds")
	fmt.Println("DNS:", NavTiming.DNS, "milliseconds")
	fmt.Println("TLS:", NavTiming.TLS, "milliseconds")

}
