package loadevent

import "testing"


func Test_LoadEventMS(t *testing.T) {
	// create a test for the LoadEventMS function 
	result, err := LoadEventMS("https://mamdouh.co.uk")
	if err != nil {
		t.Fatalf("LoadEventMS returned error: %v", err)
	}
	// Optionally, add assertions here to validate the performance data.
	t.Logf("LoadEvent: %.2f seconds", result.LoadEvent)
	t.Logf("TTFB: %.2f ms", result.TTFB)
	t.Logf("DNS: %.2f ms", result.DNS)
	t.Logf("TLS: %.2f ms", result.TLS)
	if len(result.SlowImages) > 0 {
	t.Logf("SlowImages: %v", result.SlowImages)
	}
}