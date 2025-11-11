package main

import (
    "fmt"
    "urlpinger/data"
    "urlpinger/fetch"
)





func main() {
    code, headers, body, dur , err := fetch.Fetch(data.First.URL)
    if err != nil {
        fmt.Println("error:", err)
        return
    }
    fmt.Println("status code:", code)
    fmt.Println("headers:", headers)
	fmt.Println("body bytes:", len(body))
	fmt.Println("durtion:", dur)
}
