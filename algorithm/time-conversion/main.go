package main

import (
	"fmt"
	"log"
	"strconv"
)

func timeConversion(s string) string {

	minute := s[3:5]
	second := s[6:8]
	zone := s[8:]
	hourTMP, _ := strconv.Atoi(s[:2])
	var hour string

	if zone == "PM" {
		if hourTMP == 12 {
			hour = "12"
		} else {
			res := hourTMP + 12
			hour = strconv.Itoa(res)
		}
		return fmt.Sprintf("%s:%s:%s", hour, minute, second)
	}

	hour = s[:2]
	if hour == "12" {
		hour = "00"
	}
	return fmt.Sprintf("%s:%s:%s", hour, minute, second)

}

func main() {
	s := "12:00:00AM"
	log.Println(timeConversion(s))
}
