package main

import (
	"log"
	"strings"
)

func Palindrome(s string) bool {
	// palindrome
	// aaa = aaa = true
	// ada = ada = true

	// non palindrome
	// aku = uka = false
	// kuki = ikuk = false

	var tmp []string
	for i := len(s) - 1; i >= 0; i-- {
		tmp = append(tmp, string(s[i]))
	}

	checker := strings.Join(tmp, "")

	if s == checker {
		return true
	} else {
		return false
	}

}

func main() {
	s := "aa"
	log.Println(Palindrome(s))
}
