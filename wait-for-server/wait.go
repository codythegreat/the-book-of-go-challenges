// waitForServer attempts to contact the server of a URL
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("Server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	url := "http://www.codymaxie.com"
	if err := WaitForServer(url); err != nil {
		log.Fatalf("Site is down: %v\n", err)
	} else {
		fmt.Println("site was available")
	}
}
