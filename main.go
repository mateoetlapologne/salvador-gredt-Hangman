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

func (h *HangManData) init() { //func to initialize the game
	h.Attempts = 10
	h.ToFind = randomWord()
	n := (len(h.ToFind) / 2) - 1
	fmt.Println(h.ToFind)
	h.updateword(string(h.ToFind[n]))
	fmt.Println("Bienvenue au Jeu du Pendu, Il te reste ", h.Attempts, " essai(s) \nVoila les lettres que tu a trouvée(s) : ", h.Word)
}
func isintheword(word string, letter string) bool { //func to check if the letter is in the word
	for _, v := range word {
		if string(v) == letter {
			return true
		}
	}
	return false
}
func (h *HangManData) updateword(letter string) { //func to update the display
	for _, v := range h.ToFind {
		if string(v) == letter {
			h.Word += string(v)
		} else {
			h.Word += "_"
		}
	}
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
