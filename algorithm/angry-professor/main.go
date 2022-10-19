package main

import "log"

func angryProfessor(k int32, a []int32) string {
	const (
		cancel    string = "YES"
		notCancel string = "NO"
	)

	var onTime []int32
	for _, v := range a {
		if v <= 0 {
			onTime = append(onTime, v)
		}
	}

	onTimeTotal := len(onTime)
	if int32(onTimeTotal) > k || int32(onTimeTotal) == k {
		return notCancel
	} else {
		return cancel
	}
}

func main() {
	k := 2
	a := []int32{0, -1, 2, 1}
	log.Println(angryProfessor(int32(k), a))
}
