// minimal echo server

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Please provide port")
		return
	}

	port := os.Args[1]
	http.HandleFunc("/", handler) // call handler on each request
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}

// return path for each request
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
