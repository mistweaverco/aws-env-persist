package main

import (
	"fmt"
	"os"
	"strings"
)

func complete() {
	partial := ""
	if len(os.Args) > 2 {
		partial = os.Args[2]
	}
	if _, ok := os.LookupEnv("COMP_LINE"); !ok {
		return
	}
	for _, word := range allowedArgs {
		if partial == "" || strings.HasPrefix(word, partial) {
			fmt.Println(word)
		}
	}

	os.Exit(0)
}
