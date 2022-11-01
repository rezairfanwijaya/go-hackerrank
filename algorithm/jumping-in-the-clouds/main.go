package main

import "log"

func jumpingOnClouds(c []int32, k int32) int32 {
	// initial power
	power := 100

	// if no action to jump
	if k == 0 {
		return int32(power)
	}

	// initial cloud
	const (
		THUNDER = 1
	)

	// initial jump step
	jump := k

	// initial loop
	loop := true

	for loop {
		// if jump > length cloud, return to cloud[0]
		if int(jump) > len(c)-1 {
			// current cloud
			cloud := c[0]
			// decrease power
			power = decresePower(cloud, power, THUNDER)
			// stop loop
			loop = false
		} else {
			// current cloud
			cloud := c[jump]
			// decrease power
			power = decresePower(cloud, power, THUNDER)
			// increase jump
			jump += k
		}
	}

	return int32(power)
}

func decresePower(cloud int32, power int, THUNDER int) int {
	if int(cloud) == THUNDER {
		return power - 1 - 2
	} else {
		return power - 1
	}
}

func main() {
	c := []int32{1, 1, 1, 0, 1, 1, 0, 0, 0, 0}
	k := 3
	log.Println(jumpingOnClouds(c, int32(k)))
}
