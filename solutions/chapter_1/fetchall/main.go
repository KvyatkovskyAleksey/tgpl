// concurrent load for urls list and log about download time and response size for each

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, site_url := range os.Args[1:] {
		go fetch(site_url, ch) // run coroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receiving from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(site_url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(site_url)
	if err != nil {
		ch <- fmt.Sprintf("Error on get url %s: %v\n", site_url, err)
		return
	}
	defer resp.Body.Close()
	writeToFile(site_url, resp, ch)
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // need to close for don't use resourcses anymore
	if err != nil {
		ch <- fmt.Sprintf("Error on count bytes %s, %v\n", site_url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, site_url)
}

func writeToFile(site_url string, resp *http.Response, ch chan<- string) {
	filename, err := url.Parse(site_url)
	if err != nil {
		ch <- fmt.Sprintf("Can't parse url %s, %v", site_url, err)
		return
	}

	//read body into a buffer
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error on read bytes %s: %v", site_url, err)
	}

	// write to file
	err = os.WriteFile(filename.Host, bytes, 0644)
	if err != nil {
		ch <- fmt.Sprintf("Error on file write (%s): %v", site_url, err)
	}
}
