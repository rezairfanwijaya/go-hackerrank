package main

func ChocolateFeast(n int32, c int32, m int32) int32 {
	initialChoho := n / c
	var chocoExchange int32
	var chocoEaten int32 = initialChoho

	for {
		initialChoho = initialChoho - m
		if initialChoho < 0 {
			initialChoho = (n / c) - chocoExchange
			if initialChoho < m {
				break
			}
			continue
		}

		chocoEaten++
		chocoExchange++
	}

	return chocoEaten
}
