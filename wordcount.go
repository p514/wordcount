package main

import (
	"fmt"
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
	src := readInput()
	n := count(src)
	fmt.Println(n)
}

func readInput() string {
	return os.Args[1]	
}
