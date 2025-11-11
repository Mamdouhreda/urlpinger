package fetch

import (
	"io"
	"net/http"
	"time"
)

// Fetch returns the full response body and headers for a URL.
func Fetch(url string) (int, http.Header, string, time.Duration ,error) {
	start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        return 0, nil, "", 0, err
    }
    defer resp.Body.Close()

    b, err := io.ReadAll(resp.Body)
	//caculate the duration
	dur := time.Since(start)
    if err != nil {
        return resp.StatusCode, resp.Header, "", dur ,err
    }
    return resp.StatusCode, resp.Header, string(b), dur ,nil
}
