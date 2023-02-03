package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var file_name, word string

func check_word(word string, banned_word string) bool {
	if strings.ToLower(word) == banned_word {
		return true
	}
	return false
}

func main() {
	flag := false
	fmt.Scan(&file_name, &word)
	file, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if check_word(word, scanner.Text()) {
			flag = true
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(flag)
}
