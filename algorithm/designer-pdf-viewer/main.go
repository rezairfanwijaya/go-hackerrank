package main

import "log"

func designerPdfViewer(h []int32, word string) int32 {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "q", "y", "z"}
	pairAlphabet := make(map[string]int32)
	var letterHeight int32 = 0
	lenWord := len(word)

	// pairing alphabet with value of h params
	for index, v := range h {
		pairAlphabet[alphabet[index]] = v
	}

	// sync word with pair alphabet
	for _, v := range word {
		key := string(v)
		height := pairAlphabet[key]
		if height > letterHeight {
			letterHeight = height
		}
	}

	selectionArea := lenWord * int(letterHeight)
	return int32(selectionArea)
}

func main() {
	h := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}
	word := "zaba"
	log.Println(designerPdfViewer(h, word))
}
