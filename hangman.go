package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

type HangManData struct {
	Word         string
	ToFind       string
	Attempts     int
	KnownLetters []string
	TriedLetters []string
}

func main() {
	h := HangManData{}
	h.init()
}

func (h *HangManData) init() { //func to initialize the game
	path := os.Args[1]
	h.Attempts = 10
	h.ToFind = randomWord(path)
	for _, v := range h.ToFind {
		h.Word += "_"
		_ = v //to avoid the error
	}
	n := (len(h.ToFind) / 2) - 1
	fmt.Println("Bienvenue au Jeu du Pendu")
	h.KnownLetters = append(h.KnownLetters, string(h.ToFind[n]))
	h.updateword()
	h.game()
}

func (h *HangManData) game() { //func to play the game
	if h.Attempts != 10 {
		h.Displayhangman()
	}
	if h.Word == h.ToFind {
		fmt.Println("Vous avez gagné\n le mot était ", h.ToFind)
	} else if h.Attempts == 0 {
		fmt.Println("Vous avez perdu,le mot était ", h.ToFind)
	} else {
		var entry string
		fmt.Println("Il te reste ", h.Attempts, " essai(s) \nVoila les lettres que tu a trouvée(s) : ", h.Word, "\nVeuillez entrer une lettre : ")
		fmt.Scan(&entry)
		if len(entry) == 1 {
			if isintheword(h.ToFind, entry) {
				if AlreadyKnown(h, entry) {
					fmt.Println("Tu as déjà trouvé cette lettre")
					h.game()
				} else {
					h.KnownLetters = append(h.KnownLetters, entry)
					h.updateword()
					fmt.Println("Bien joué !")
					h.game()
				}
			} else if !alreadytried(h, entry) {
				h.Attempts--
				fmt.Println("La lettre ", entry, " n'est pas dans le mot")
				h.TriedLetters = append(h.TriedLetters, entry)
				h.game()
			} else {
				fmt.Println("Tu as déjà essayé cette lettre")
				h.game()
			}
		} else if len(entry) >= 2 {
			if entry == h.ToFind {
				fmt.Println("Vous avez gagné\n le mot était ", h.ToFind)
			} else {
				h.Attempts--
				h.Attempts--
				fmt.Println("Ce n'est pas le mot, tu perds 2 essais")
				h.game()
			}
		}
	}
}

func (h *HangManData) Displayhangman() { //func to display the hangman
	file, err := ioutil.ReadFile("src/hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(file)[-(h.Attempts-9)*71 : -(h.Attempts-9)*71+71])
}
func isintheword(word string, letter string) bool { //func to check if the letter is in the word
	for _, v := range word {
		if string(v) == letter {
			return true
		}
	}
	return false
}
func alreadytried(h *HangManData, letter string) bool { //func to check if the letter is already tried
	for _, v := range h.TriedLetters {
		if v == letter {
			return true
		}
	}
	return false
}
func AlreadyKnown(h *HangManData, letter string) bool { //func to check if the letter is already known
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
func randomWord(arg string) string { //func to get a random word from the file.txt
	//read the file
	file, err := os.Open(arg)
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
