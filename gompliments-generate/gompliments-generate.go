package main

import (
	"fmt"

	wnr "github.com/lloyd/wnram"

	"os"
)

func main() {
	wnpath := os.Args[1]
	wn, err := wnr.New(wnpath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Searching")
	if found, err := wn.Lookup(wnr.Criteria{
		Matching: "good",
		POS:      []wnr.PartOfSpeech{wnr.Adjective}}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Found:")
		fmt.Println(found)
		fmt.Println(err)
		for _, f := range found {
			fmt.Println(f)
			f.Dump()
		}
	}
}
