package main

import (
	"fmt"
	"log"
)

func stairCase(n int32) {
	var i, j int32

	for i = 1; i <= n; i++ { // baris

		for j = 1; j <= n; j++ { // kolom
			str := "#"

			log.Println("i", i)
			log.Println("n-j",n-j)
			if n-j >= i {
				str = " "
			}

			fmt.Print(str)
		}

		fmt.Println("")
	}
}

func main() {
	stairCase(5)
}
