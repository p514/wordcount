package main

import (
	"errors"
	"fmt"
//	"flag"
	"os"
	"strings"
	"unicode"
)

func count(src string) int {
	f := func(c rune) bool {
		return !unicode.IsSpace(c)
	}
	return len(strings.FieldsFunc(src, f))
}

func main() {
	// input := " Welcome aboard!"
	// words := strings.Fields(input)
	// fmt.Println(words, len(words))

	src, err := readInput()
	if err != nil {
		fail(err)
	}
	n := count(src)

	fmt.Println(n)
}

func readInput() (src string, err error) {
	src = os.Args[1]
	if src == "" {
		return src, errors.New("missing string")
	}
	return src, nil
}
		
func fail(err error) {
	fmt.Println("wordcount:", err)
	os.Exit(1)
}