package cutthesticks

func CutTheSticks(arr []int32) []int32 {
	loop := true
	var result []int32

	for loop {
		length := len(arr)

		// loop akan berhenti ketika length arr = 0
		if length == 0 {
			loop = false
		} else {
			// lakukan penguarangan pada array
			newArr := MinusArrayWithMin(arr)
			// timpa array lama dengan array baru
			arr = newArr
			// append length
			result = append(result, int32(length))
		}
	}

	return result
}

func FindMinElement(arr []int32) int32 {
	min := arr[0]

	for _, v := range arr {
		if v < min {
			min = v
		}
	}

	return min
}

func RemoveElementInMiddleArray(arr []int32, index int32) []int32 {
	var merged []int32

	start := arr[:index]
	end := arr[index+1:]
	merged = append(merged, start...)
	merged = append(merged, end...)

	return merged
}

func MinusArrayWithMin(arr []int32) []int32 {
	min := FindMinElement(arr)
	var tmp []int32
	var result []int32

	for i := 0; i < len(arr); i++ {
		minus := arr[i] - min
		tmp = append(tmp, minus)
	}

	// hilangkan angka 0
	for _, v := range tmp {
		if v == 0 {
			continue
		} else {
			result = append(result, v)
		}
	}

	return result
}
