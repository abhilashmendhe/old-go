package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	userName := fetchUser()
	// likes := fetchUserLikes(userName)
	// match := fetchUserMatch(userName)
	respch := make(chan any, 2) // any value in this channel, with size 2
	wg := &sync.WaitGroup{}

	wg.Add(2)
	// it is little faster to fetch values,because of goroutines
	go fetchUserLikes(userName, respch, wg)
	go fetchUserMatch(userName, respch, wg)
	wg.Wait() // will block till we get all our values in respch(channel)
	//  with for you get deadlock fatal error,, it is ranging over the channel, we have to close the channel
	close(respch) // it is not printing anything, we are closing the channel, but does not
	// gaurantee if the above two function are done.. seems like we are closing the channel before
	// we read anything from the channel.How to synchronize? We will use sync.waitgroup..
	for resp := range respch {
		fmt.Println("resp: ", resp)
	}
	fmt.Println("Took: ", time.Since(start), " ms.")
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return "Abhi"
}

func fetchUserLikes(username string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 180)

	respch <- 102
	wg.Done() // it subtracts the value by 1 or pop the vals from line 15

}

func fetchUserMatch(username string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	respch <- "Ana"
	wg.Done() // it subtracts the value by 1 or pop the vals from line 15
}
