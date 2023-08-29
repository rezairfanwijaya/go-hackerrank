package main

import (
	"log"
	"strings"
	"unicode"
)

func CaesarCipher(s string, k int32) string {
	// init originial alphabet
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	// pemisahan alphabet
	front := alphabet[k:]
	back := alphabet[:k]

	// penggabungan alphabet
	var hashAlphabet []string
	hashAlphabet = append(hashAlphabet, front...)
	hashAlphabet = append(hashAlphabet, back...)
	pairAlphabet := make(map[string]string)
	for n := range alphabet {
		pairAlphabet[alphabet[n]] = hashAlphabet[n]
	}

	// mencari kata berdasarkan inputan dan mencocokan dengan aplhabet yang sudah di acak
	var tmp []string
	for n := range s {
		if string(s[n]) == "-" {
			tmp = append(tmp, "-")
		}

		// cek uppercase
		if unicode.IsUpper(rune(s[n])) {
			// kira lowerkan dulu
			lowerString := strings.ToLower(string(s[n]))

			// cek ke pairing map
			pairing := pairAlphabet[lowerString]
			tmp = append(tmp, strings.ToUpper(pairing))
		} else {
			tmp = append(tmp, pairAlphabet[string(s[n])])
		}

	}

	// hasil
	result := strings.Join(tmp, "")
	return result
}

func main() {
	s := "middle-Outz"
	log.Println(CaesarCipher(s, 2))
}
