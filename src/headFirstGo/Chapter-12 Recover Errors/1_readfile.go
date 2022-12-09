package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(filename string) (*os.File, error) {
	fmt.Println("Opening file....")
	return os.Open(filename)
}

func CloseFile(file *os.File) {
	fmt.Println("Closing file...")
	file.Close()
}

func GetFloats(filename string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(filename)
	if err != nil {
		return nil, errors.New("Not able to open the file...")
	}
	defer CloseFile(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, errors.New("Can't parse to float...")
		}
		numbers = append(numbers, number)
	}
	// CloseFile(file) /// it won't close if there persist an error above the some called funcs.
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}

func main() {

	filename := os.Args[1]

	numbers, err := GetFloats(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(numbers)
}
