package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	userName := fetchUser()
	likes := fetchUserLikes(userName)
	match := fetchUserMatch(userName)

	fmt.Println("Likes: ", likes)
	fmt.Println("Match: ", match)

	fmt.Println("Took: ", time.Since(start), " ms.")
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return "Abhi"
}

func fetchUserLikes(username string) int {
	time.Sleep(time.Millisecond * 180)
	return 102
}

func fetchUserMatch(username string) string {
	time.Sleep(time.Millisecond * 100)
	return "ANA"
}
