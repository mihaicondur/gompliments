package main

import (
	"fmt"

	wnr "github.com/lloyd/wnram"

	"os"

	"log"

	"math/rand"
)

func main() {
	wordnet := loadWordnet()
	templates := Templates()
	template := templates[rand.Intn(len(templates))]
	fmt.Println()
	for _, segment := range template {
		fmt.Print(segment.ToText(wordnet))
	}
	fmt.Println()
}

func loadWordnet() *wnr.Handle {
	wnpath := os.Args[1]
	wn, err := wnr.New(wnpath)
	if err != nil {
		log.Fatal(err)
	}
	return wn
}
