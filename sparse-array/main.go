package main

import "fmt"

func matchingString(strings []string, queries []string) []int32 {

	res := make(map[string]int32)
	for _, v := range queries {
		res[v] = 0
	}

	for _, value := range strings {
		for _, v1 := range queries {
			if value == v1 {
				res[v1]++
			}
		}
	}

	final := make([]int32, 0)
	for _, v := range res {
		final = append(final, v)
	}

	return final

}

func main() {
	myString := []string{"Reza", "abdas", "Reza", "abdas"}
	myQuery := []string{"Reza", "lmn", "abdas"}
	fmt.Println(matchingString(myString, myQuery))

}
