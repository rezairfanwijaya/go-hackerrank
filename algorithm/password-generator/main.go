package main

import (
	"log"
	"strings"
)

func PasswordGenerator(rawPassword string) string {
	var reversePassword []string
	vokal := map[string]string{
		"a": "a",
		"i": "i",
		"u": "u",
		"e": "e",
		"o": "o",
		"A": "A",
		"I": "I",
		"U": "U",
		"E": "E",
		"O": "O",
	}

	consonant := map[string]string{
		"b": "b",
		"c": "c",
		"d": "d",
		"f": "f",
		"g": "g",
		"h": "h",
		"j": "j",
		"k": "k",
		"l": "l",
		"m": "m",
		"n": "n",
		"p": "p",
		"q": "q",
		"r": "r",
		"s": "s",
		"t": "t",
		"v": "v",
		"w": "w",
		"x": "x",
		"y": "y",
		"z": "z",
	}

	// reverse password
	for i := len(rawPassword) - 1; i >= 0; i-- {
		reversePassword = append(reversePassword, string(rawPassword[i]))
	}

	var swtichVocal []string

	// proses pertukaran huruf vokal dan consonant menjadi uppper atau lowwer
	for _, v := range reversePassword {
		//  hapus spasi
		if v == " " {
			continue
		}
		values, isExsist := vokal[v]
		if isExsist {
			switch values {
			case "a":
				swtichVocal = append(swtichVocal, "E")
			case "i":
				swtichVocal = append(swtichVocal, "O")
			case "u":
				swtichVocal = append(swtichVocal, "A")
			case "e":
				swtichVocal = append(swtichVocal, "I")
			case "o":
				swtichVocal = append(swtichVocal, "U")
			case "A":
				swtichVocal = append(swtichVocal, "e")
			case "I":
				swtichVocal = append(swtichVocal, "o")
			case "U":
				swtichVocal = append(swtichVocal, "a")
			case "E":
				swtichVocal = append(swtichVocal, "i")
			case "O":
				swtichVocal = append(swtichVocal, "u")
			}
		} else {
			values, isExsist := consonant[v]
			if isExsist {
				upper := strings.ToUpper(values)
				swtichVocal = append(swtichVocal, upper)
			} else {
				lower := strings.ToLower(v)
				swtichVocal = append(swtichVocal, lower)
			}
		}
	}

	res := strings.Join(swtichVocal, "")
	return res
}

func main() {
	log.Println(PasswordGenerator("Semangat Pagi 12!#"))
}
