package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var fileName, word string

func main() {
	char := "*"
	fmt.Scan(&fileName)
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var bannedWords []string
	for scanner.Scan() {
		bannedWords = append(bannedWords, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Scan(&word)
	for word != "exit" {

		for _, bannedWord := range bannedWords {
			if strings.ToLower(word) == strings.ToLower(bannedWord) {
				word = strings.Repeat(char, len(word))
				break
			}
		}
		fmt.Println(word)
		fmt.Scan(&word)
	}

	fmt.Println("Bye!")

}
