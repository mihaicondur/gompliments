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

	// for _, w := range words {
	// 	w.Dump()
	// }

	myWord := words[1]
	var similar []string
	for _, r := range myWord.Related(wnr.SimilarTo) {
		//r.Dump()
		similar = append(similar, r.Word())
		for _, s := range r.Synonyms() {
			similar = append(similar, s)
		}
	}
	r := rand.Intn(len(similar))
	return similar[r]
}
