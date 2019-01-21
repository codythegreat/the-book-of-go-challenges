// Fetches URLs concurrently and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	// notes: we will use goroutines to create concurrent execution. we use a channel so
	// 				that the goroutines can communicate with the execution of the main function
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // starts goroutine
	}
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Couldn't create file")
	}

	defer file.Close()

	for range os.Args[1:] {
		_, err := file.WriteString(<-ch + "\n") // receive from channel ch and write to file
		if err != nil {
			fmt.Println("failed to write to file")
		}
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // prevents resources from leaking
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
