package repeatedstring

func RepetedString(s string, n int64) int64 {
	// counter for letter a
	count := 0

	// slice string for all letter
	var tmp []string

	// condition for infinite loop
	loop := true

	for loop {
		// open all element in string
		for i := 0; i < len(s); i++ {
			// check if letter == a increment counter
			letter := string(s[i])
			if letter == "a" {
				count++
			}

			tmp = append(tmp, letter)
			len := len(tmp)

			// if length tmp == n stop loop
			if int64(len) == n {
				loop = false
				break
			}
		}
	}

	return int64(count)
}
