package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	now := time.Now()
	userId := 10
	reschp := make(chan string, 64)

	wg := &sync.WaitGroup{}

	go fetchUserData(userId, reschp, wg)
	wg.Add(1)
	go fetchUserRecommendations(userId, reschp, wg)
	wg.Add(1)
	go fetchUserLikes(userId, reschp, wg)
	wg.Add(1)
	wg.Wait()

	close(reschp) // need to close this...
	for resp := range reschp {
		fmt.Println(resp)
	}
	fmt.Println(time.Since(now))
}

func fetchUserData(userID int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(80 * time.Millisecond)
	respch <- "user data"
	wg.Done()
}

func fetchUserRecommendations(userId int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(120 * time.Millisecond)
	respch <- "user recommendation"
	wg.Done()
}

func fetchUserLikes(userId int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(50 * time.Millisecond)
	respch <- "user likes"
	wg.Done()
}
