package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type HangManData struct {
	Word             string
	ToFind           string
	Attempts         int
	HangmanPositions [10]string
}

func main() {
	h := HangManData{}
	h.init()
}

func (h *HangManData) init() {
	h.ToFind = randomWord()
	n := (len(h.ToFind) / 2) - 1
	h.Word = string(h.ToFind[n])
	var pword string
	fmt.Println(h.ToFind)
	for _, v := range h.ToFind {
		if string(v) == h.Word {
			pword += string(v)
			h.Word = ""
		} else {
			pword += "_"
		}
	}
	fmt.Println(pword)
}

func randomWord() string {
	//read the file
	file, err := os.Open("src/words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//make a slice of words
	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	//return a random word
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}
