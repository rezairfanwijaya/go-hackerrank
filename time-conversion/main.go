package main

import (
	"fmt"
	"strconv"
)

func timeConversion(s string) string {

	hour, min, sec := s[:2], s[3:5], s[6:8]
	hourInt, _ := strconv.Atoi(hour)

	if s[8] == 'P' {
		if hour != "12" {
			hour = fmt.Sprintf("%d", (12 + (hourInt)))
		}
	} else {
		if hour == "12" {
			hour = "00"
		}
	}
	return hour + ":" + min + ":" + sec

	// fmt.Println("hour", hour)
	// fmt.Println("min", min)
	// fmt.Println("sec", sec)

}

func main() {
	jam := timeConversion("11:59:59PM")
	fmt.Println(jam)
}
