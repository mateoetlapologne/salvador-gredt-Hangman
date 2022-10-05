package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type HangManData struct {
	Word             string
	ToFind           string
	Attempts         int
	HangmanPositions [10]string
	KnownLetters     []string
}

func main() {
	h := HangManData{}
	h.init()
}

func (h *HangManData) init() { //func to initialize the game
	h.Attempts = 10
	h.ToFind = randomWord()
	for _, v := range h.ToFind {
		h.Word += "_"
		_ = v //to avoid the error
	}
	n := (len(h.ToFind) / 2) - 1
	fmt.Println(h.ToFind)
	fmt.Println("Bienvenue au Jeu du Pendu")
	h.KnownLetters = append(h.KnownLetters, string(h.ToFind[n]))
	h.updateword()
	h.game()
}
func (h *HangManData) game() {
	//print a txt file
	url := "src/ascii-art/pos" + strconv.Itoa(h.Attempts-1) + ".txt"
	os.OpenFile(url, os.O_RDONLY, 0666)
	if h.Word == h.ToFind {
		fmt.Println("Vous avez gagné")
	} else if h.Attempts == 1 {
		fmt.Println("Vous avez perdu")
	} else {
		var entry string
		fmt.Println("Il te reste ", h.Attempts, " essai(s) \nVoila les lettres que tu a trouvée(s) : ", h.Word, "\nVeuillez entrer une lettre : ")
		fmt.Scan(&entry)
		if isintheword(h.ToFind, entry) {
			if AlreadyKnown(h, entry) {
				fmt.Println("Tu as déjà trouvé cette lettre")
			} else {
				h.KnownLetters = append(h.KnownLetters, entry)
				h.updateword()
				fmt.Println("Bien joué !")
			}
		} else {
			h.Attempts--
			fmt.Println("La lettre ", entry, " n'est pas dans le mot")
		}
		h.game()
	}
}
func isintheword(word string, letter string) bool { //func to check if the letter is in the word
	for _, v := range word {
		if string(v) == letter {
			return true
		}
	}
	return false
}
func AlreadyKnown(h *HangManData, letter string) bool {
	for _, v := range h.KnownLetters {
		if v == letter {
			return true
		}
	}
	return false
}
func (h *HangManData) updateword() { //func to update the display
	h.Word = ""
	for _, v := range h.ToFind {
		if AlreadyKnown(h, string(v)) {
			h.Word += string(v)
		} else {
			h.Word += "_"
		}
	}
}

func randomWord() string { //func to get a random word from the file.txt
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
