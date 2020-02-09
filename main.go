package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello\n")
}

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":8090", nil)
}
