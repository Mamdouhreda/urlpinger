package fetch
import (
    "io"
    "net/http"
)

// Fetch returns the full response body and headers for a URL.
func Fetch(url string) (int, http.Header, string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return 0, nil, "", err
    }
    defer resp.Body.Close()

    b, err := io.ReadAll(resp.Body)
    if err != nil {
        return resp.StatusCode, resp.Header, "", err
    }
    return resp.StatusCode, resp.Header, string(b), nil
}
