package main

import (
	"fmt"
)

func parentheses(s string) bool {
	// cek panjang string
	// jika ganjil maka return false
	if len(s)%2 != 0 {
		return false
	}

	// buat stack dan map
	stack := []rune{}
	openCharacters := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	// loop string
	for _, v := range s {
		// jika karater match dengan key di map opencharacter
		// maka masukan value dari key tadi kedalam stack
		if closeCharacter, isExist := openCharacters[v]; isExist {
			stack = append(stack, closeCharacter)
			continue
		}

		// ambil index stack
		// pastikan panjang stack lebih dari 0 dan close character di stack terkahir
		// tidak sama dengan string sekarang (double closer single opener)
		stackIndex := len(stack) - 1
		if stackIndex < 0 || v != stack[stackIndex] {
			return false
		}

		// kosongkan stack untuk menampung next closer
		stack = stack[:stackIndex]
	}

	return true
}

func main() {
	fmt.Println(parentheses("{[]{}}"))
}
