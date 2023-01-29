package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var fileName string

func main() {
	_, err := fmt.Scan(&fileName)
	if err != nil {
		return
	}
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
