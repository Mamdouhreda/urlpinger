package main

import (
    "fmt"
    "urlpinger/data"
    "urlpinger/load"
)
// NavTiming represents various performance timing metrics for a page
type NavTiming struct {
    DNS                float64 // DNS: Time spent resolving the domain name (ms) — typical good range: ~20-120 ms (20-120) :contentReference[oaicite:0]{index=0}
    Connect            float64 // Connect: Time to establish TCP connection (ms) — typical good range: ~~ <100-200 ms (no exact number found) 
    TLS                float64 // TLS: Time spent in SSL/TLS handshake (ms) — typical overhead noted: ~30-50 ms extra on good connections :contentReference[oaicite:1]{index=1}
    TTFB               float64 // TTFB: Time to First Byte (ms) — recommended <200 ms is “excellent”, <800 ms “good”/acceptable :contentReference[oaicite:2]{index=2}
    DOMContentLoaded   float64 // DOMContentLoaded: When HTML is fully parsed (ms) — no widely-quoted average found; aim: <1000-1500 ms for good experience :contentReference[oaicite:3]{index=3}
    LoadEvent          float64 // LoadEvent: When all resources finish loading (ms) — typical load time benchmark: ≤2000-3000 ms (2-3 s) is acceptable, under 2000 ms is better :contentReference[oaicite:4]{index=4}
    FCP                float64 // FCP: First Contentful Paint (ms) — goal often <1000 ms on desktop, ~1500-2000 ms mobile; but exact average varies widely 
}

func main() {
 
	loadEventEnd, err := loadevent.LoadEventMS(data.Second.URL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("LoadEvent:", loadEventEnd)

}
