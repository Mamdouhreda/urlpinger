package loadevent

import (
	"context"
	"github.com/chromedp/chromedp"
)

// LoadEventMS navigates to the URL and returns loadEventEnd in milliseconds.
func LoadEventMS(url string)(float64, error){
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var loadEventEnd float64

	// Run Chrome and measure loadEventEnd
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		// Wait for the page to finish loading and get performance timing
		chromedp.Evaluate(`window.performance.timing.loadEventEnd - window.performance.timing.navigationStart`, &loadEventEnd),
	)
	if err != nil {
		return 0, err
	}

	return loadEventEnd, nil
}

