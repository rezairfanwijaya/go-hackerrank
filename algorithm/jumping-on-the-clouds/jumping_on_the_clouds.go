package jumpingontheclouds

import "log"

func JumpingOnClouds(c []int32) int32 {
	// ? varible untuk menampung step yang dilakui
	// ? untuk melewati cloud
	var step []int32
	triple := false
	for i := 0; i < len(c); i++ {
		log.Println("index", i)
		// ? ambil 3 deret angka
		indexSecond := i + 1
		indexThird := i + 2
		length := len(c)

		if length < 3 {
			if c[i] != 1 {
				step = append(step, c[i])
			}
		}

		if indexSecond < length && indexThird < length {
			second := int(c[indexSecond])
			third := int(c[indexThird])

			if c[i] == 0 && second == 0 && third == 0 {
				if !triple {
					step = append(step, c[i], int32(second))
					triple = true
				} else {
					step = append(step, int32(third))
				}
				log.Println("triple", i)
				i++
				log.Println(c)
			} else {
				if c[i] != 1 {
					step = append(step, c[i])
				}
			}
		} else {
			log.Println("masuk ke ", i)
			if c[i] != 1 {
				step = append(step, c[i])
			}
		}
	}

	log.Println(step)
	log.Println(len(step) - 1)
	return 0
}
