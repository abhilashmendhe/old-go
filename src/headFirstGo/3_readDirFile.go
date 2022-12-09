package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func readFileDir(path string) {

	fmt.Println("Path:", path)
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		filepath := filepath.Join(path, file.Name())
		// filepath := path + "/" + file.Name()
		if file.IsDir() {
			readFileDir(filepath)
		} else {
			fmt.Println(filepath)
		}
	}
}

func main() {

	readFileDir("./")
}
