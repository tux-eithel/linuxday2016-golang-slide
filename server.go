package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	flag.Parse()
	http.HandleFunc("/", func(h http.ResponseWriter, r *http.Request) {
		fmt.Fprint(h, "hello from server")

	})
	http.ListenAndServe(":8080", nil)
}
