package main

import "log"

func pageCount(n, p int32) int32 {
	if n == p {
		return 0
	}

	start := 1
	counterFromFront := 0
	counterFromBack := 0

	for i := 1; i <= int(n); i++ {
		if p == 1 {
			break
		}

		left := start + 1
		right := left + 1
		start = right
		counterFromFront++

		if left == int(p) || right == int(p) {
			start = int(n)
			break
		}

		if left == int(n) || right == int(n) {
			start = int(n)
			break
		}
	}

	for i := n; i >= 1; i-- {
		if n%2 != 0 {
			if n == p || n-1 == p {
				break
			}

			left := n
			right := left - 1
			start = int(right)
			counterFromBack++

			if left == p || right == p {
				start = 1
				break
			}

			if left == n || right == n {
				start = 1
				break
			}
		} else {
			left := start - 1
			right := left - 1
			start = right
			counterFromBack++

			if left == int(p) || right == int(p) {
				start = 1
				break
			}

			if left == int(n) || right == int(n) {
				start = 1
				break
			}
		}

	}

	if counterFromFront < counterFromBack {
		return int32(counterFromFront)
	} else {
		return int32(counterFromBack)
	}
}

func main() {
	n := 6
	p := 5
	log.Println(pageCount(int32(n), int32(p)))
}
