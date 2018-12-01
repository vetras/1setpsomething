package main

import (
	"io"
	"fmt"
	"net/http"
)

var count int = 0

func version(w http.ResponseWriter, r *http.Request) {
	count = count + 1
	io.WriteString(w, fmt.Sprintf(" GO HTTP echo reply v1 -- %d", count))
}

func main() {
	
	http.HandleFunc("/", version)
	http.HandleFunc("/version", version)

	fmt.Println("Listening on port 8000 ...")
	http.ListenAndServe(":8000", nil)
}
