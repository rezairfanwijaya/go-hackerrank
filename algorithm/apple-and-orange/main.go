package main

import (
	"fmt"
)

func appleAndOrange(s, t, a, b int32, apples, oranges []int32) {
	var appleBucket = make(map[string]int)
	var orangeBucket = make(map[string]int)

	for _, apple := range apples {
		distance := a + apple
		if distance >= s && distance <= t {
			appleBucket["apple"] += 1
		}
	}

	for _, orange := range oranges {
		distance := b + orange
		if distance >= s && distance <= t {
			orangeBucket["orange"] += 1
		}
	}

	fmt.Println("apple", appleBucket["apple"])
	fmt.Println("orange", orangeBucket["orange"])

}

func main() {
	var (
		s int32 = 7
		t int32 = 10
		a int32 = 4
		b int32 = 12
	)

	apples := []int32{2, 3, -4}
	oranges := []int32{3, -2, -4}

	appleAndOrange(s, t, a, b, apples, oranges)
}
