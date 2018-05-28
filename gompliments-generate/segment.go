package main

import (
	"log"
	"math/rand"

	wnr "github.com/lloyd/wnram"
)

type Segment interface {
	ToText(wordnet *wnr.Handle) string
}

type Literal struct {
	Text string
}

type Expansion struct {
	POS      wnr.PartOfSpeech
	Matching string
}

func (lit Literal) ToText(wordnet *wnr.Handle) string {
	return lit.Text
}

func (exp Expansion) ToText(wordnet *wnr.Handle) string {
	words, err := wordnet.Lookup(wnr.Criteria{Matching: exp.Matching, POS: []wnr.PartOfSpeech{exp.POS}})
	if err != nil {
		log.Fatal(err)
	}

	myWord := words[rand.Intn(len(words))]
	similar := myWord.Related(wnr.SimilarTo)
	return similar[rand.Intn(len(similar))].Word()
}
