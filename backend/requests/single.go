package requests

import (
	"fmt"
	"net/http"
)

// SubmitSingleURL handles post request for single url submission
func SubmitSingleURL(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    url := r.FormValue("single-url")
    if url == "" {
        http.Error(w, "URL is required", http.StatusBadRequest)
        return
    }
	// Log and echo the URL for quick debugging
	fmt.Printf("[single] received URL: %s\n", url)
	fmt.Fprintf(w, "Received URL: %s", url)
}