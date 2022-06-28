package main

import (
	"fmt"
	"net/http"
	"time"
)

// In go you can use Go Routine to harness the power of multicore processor, so each go routine
// will be executed at the same time
// If you don't use channels the main routine will be blocked by the calling and wait for the routine to be finalized
// The "go" keyword creates a new thread
// In single core processors the thread runs in Concurrency, in multi core they are executed in parallel

//  channel <- 5 : 		  	send the value in the Channel
//  myNumber <- channel : 	Wait for a value to be sent into the channel.  When we get one, assign the value to 'myNumber'
// fmt.Println(<- channel): Wait for a value to be sent into the channel.  When we get one, log it out immediately

//In this program, it receives a list of link and check every 4 seconds, if the link is up
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	//For each link in the slice, start a new go routine passing the channel to receive the result back
	for _, link := range links {
		go checkLink(link, c)
	}

	// Function Literal, it's like a Lambda Function in c#
	for l := range c {
		go func(link string) {
			time.Sleep(4 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}
