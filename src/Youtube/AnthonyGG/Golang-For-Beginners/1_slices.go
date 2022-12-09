package main

import "fmt"

var users = []string{}

func addusers(users []string) {
	for _, user := range users {
		fmt.Println(user)
	}
}

func addUser(user ...string) {
	users = append(users, user...)
}

func main() {

	addusers([]string{"aa", "bb", "dd"})

	addUser("abhi", "weqr", "asdf")
	fmt.Println(users)

	fmt.Println(append(users[:1], users[2:]...))
}
