package main

import (
	"net/http"
)

func main() {
	// Define the URL to send the GET request to
	url := "http://alert.keg.wang/alert"
	http.Get(url)
}

