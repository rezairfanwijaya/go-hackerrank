package main

import "log"

func kangaroo(x1, v1, x2, v2 int32) string {
	if v2 > v1 && x2 > x1 {
		return "NO"
	}

	stepKangaroo1 := int(x1)
	stepKangaroo2 := int(x2)

	stop := true

	for stop {
		stepKangaroo1 += int(v1)
		stepKangaroo2 += int(v2)

		if stepKangaroo1 == stepKangaroo2 {
			stop = false
		}
	}
	return "YES"
}

func main() {
	var (
		x1 int32 = 0
		v1 int32 = 3
		x2 int32 = 4
		v2 int32 = 2
	)

	log.Println(kangaroo(x1, v1, x2, v2))
}
