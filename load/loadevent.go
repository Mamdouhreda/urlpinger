package loadevent

import (
	"context"
	"github.com/chromedp/chromedp"
	"time"
)

// NavTiming represents various performance timing metrics for a page
type NavTiming struct {
    DNS                float64 // DNS: Time spent resolving the domain name (ms) — typical good range: ~20-120 ms (20-120) :contentReference[oaicite:0]{index=0}
    TLS                float64 // TLS: Time spent in SSL/TLS handshake (ms) — typical overhead noted: ~30-50 ms extra on good connections :contentReference[oaicite:1]{index=1}
    TTFB               float64 // TTFB: Time to First Byte (ms) — recommended <200 ms is “excellent”, <800 ms “good”/acceptable :contentReference[oaicite:2]{index=2}
    LoadEvent          float64 // LoadEvent: When all resources finish loading (ms) — typical load time benchmark: ≤2000-3000 ms (2-3 s) is acceptable, under 2000 ms is better :contentReference[oaicite:4]{index=4}
}

// LoadEventMS navigates to the URL and returns loadEventEnd in milliseconds.
func LoadEventMS(url string)(NavTiming, error){
	// Create context with timeout (50 seconds)
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	// Create new Chrome instance
	ctx, cancelBrowser := chromedp.NewContext(ctx)
	defer cancelBrowser()

	  // all vairable to get back
	var (
		DNS              float64
		TLS              float64
		TTFB             float64
		loadEventEnd        float64
	)
	// Run Chrome and measure loadEventEnd
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		// Wait for the page to finish loading and get performance timing
		chromedp.Evaluate(`window.performance.timing.loadEventEnd - window.performance.timing.navigationStart`, &loadEventEnd),
		chromedp.Evaluate(`window.performance.timing.responseStart - window.performance.timing.requestStart`, &TTFB),
		// DNS lookup time
		chromedp.Evaluate(`window.performance.timing.domainLookupEnd - window.performance.timing.domainLookupStart`, &DNS),
		// TLS handshake time
		chromedp.Evaluate(`window.performance.timing.connectEnd - window.performance.timing.secureConnectionStart`, &TLS),		

	)
	if err != nil {
		return NavTiming{}, err
	}
	//convert from millsecond to seconds
	loadEventEnd = loadEventEnd / 1000
	
	return NavTiming{
		LoadEvent: loadEventEnd,
		TTFB : TTFB,
		TLS: TLS,
		DNS: DNS,
		}, nil
}

