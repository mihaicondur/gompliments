package main

import wnr "github.com/lloyd/wnram"

func Templates() [][]Segment {
	return [][]Segment{
		[]Segment{Literal{Text: "You are "}, Expansion{POS: wnr.Adjective, Matching: "beautiful"}},
	}
}
