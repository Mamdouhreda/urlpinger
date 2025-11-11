package data

// URLCheck represents a single URL to check and related metadata.
type URLCheck struct {
    ID  int
    URL string
}

// First is a sample record you can use in main.
var First = URLCheck{
    ID:  1,
    URL: "https://www.google.com/",
}

//second sample 

var Second = URLCheck{
	ID: 2,
	URL: "https://mamdouh.co.uk/",
}

