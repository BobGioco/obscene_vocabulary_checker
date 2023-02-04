package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var fileName, sentence string

func checkSentence(bannedWords []string, sentence string) string {
	char := "*"
	words := strings.Fields(sentence)
	var outputSentence, checkWord string

	for _, word := range words {
		checkWord = word
		for _, bannedWord := range bannedWords {
			if wordComparison(bannedWord, word) {
				checkWord = strings.Repeat(char, len(word))
				break
			}
		}
		outputSentence += " " + checkWord
	}
	return outputSentence
}

func wordComparison(bannedWord string, word string) bool {
	word = strings.ReplaceAll(word, ",", "")
	word = strings.ReplaceAll(word, ".", "")
	if strings.ToLower(word) == strings.ToLower(bannedWord) {
		return true
	} else {
		return false
	}
}

func main() {

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

	fmt.Scan(&sentence)
	for sentence != "exit" {
		fmt.Println(checkSentence(bannedWords, sentence))
		fmt.Scan(&sentence)
	}

	fmt.Println("Bye!")

}
