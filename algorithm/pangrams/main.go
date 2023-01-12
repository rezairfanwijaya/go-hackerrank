package main

import (
	"fmt"
	"strings"
)

func pangrams(s string) string {
	// library alphabet
	alphabet := map[string]bool{
		"a": false,
		"b": false,
		"c": false,
		"d": false,
		"e": false,
		"f": false,
		"g": false,
		"h": false,
		"i": false,
		"j": false,
		"k": false,
		"l": false,
		"m": false,
		"n": false,
		"o": false,
		"p": false,
		"q": false,
		"r": false,
		"s": false,
		"t": false,
		"u": false,
		"v": false,
		"w": false,
		"x": false,
		"y": false,
		"z": false,
		" ": false,
	}

	for _, v := range s {
		// jadikan semua inputan menjadi lower
		huruf := strings.ToLower(string(v))

		// cek input ke library
		_, isExist := alphabet[huruf]

		// jika ada ubah value menjadi true
		if isExist {
			alphabet[huruf] = true
		}
	}

	// jika ada yang false artinya no pangrams
	for _, v := range alphabet {
		if !v {
			return "not pangram"
		}
	}

	return "pangram"
}

func main() {
	fmt.Println(pangrams("We promptly judged antique ivory buckles for the  prize"))
}
