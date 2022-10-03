package main

import "log"

func gradingStudents(grades []int32) []int32 {
	const MULTIPLE int32 = 5
	/*
		rules :
			1. bulatkan grades ke kelipatan 5 terdekat
			2. kurangi pembulatan tsb dengan grades
			3. jika kurang dari 3 maka replace grades dengan pembulatan
			4. jika kelipatan 5 terdekat kurang dari 40 maka jangan ada pembulatan dan langsung return grades
	*/

	var newGrades []int32
	rounded := 0

	for _, grade := range grades {
		// cek apakah grade kelipatan 5
		if grade%MULTIPLE == 0 {
			// cek selisih
			difference := (grade + MULTIPLE) - grade
			if difference < 3 {
				rounded = int(grade + MULTIPLE)
				newGrades = append(newGrades, int32(rounded))
			} else {
				rounded = int(grade)
				newGrades = append(newGrades, int32(rounded))
			}
		} else {
			// mencari bilangan kelipatan 5 terdekat
			roundedtmp := 0
			for i := 1; i <= 5; i++ {
				tmp := grade + int32(i)
				if tmp%5 == 0 {
					roundedtmp = int(tmp)
				}
			}

			// jika kurang dari 40, jangan lakukan pembulatan
			if roundedtmp < 40 {
				newGrades = append(newGrades, grade)
			} else {
				difference := roundedtmp - int(grade)
				if difference < 3 {
					rounded = roundedtmp
					newGrades = append(newGrades, int32(rounded))
				} else {
					rounded = int(grade)
					newGrades = append(newGrades, int32(rounded))
				}
			}
		}
	}

	return newGrades
}

func main() {
	grades := []int32{73,67,38,33}
	log.Println(gradingStudents(grades))
}
