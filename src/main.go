package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s string
	s = "06:01:00AM"
	var eachParts []string
	eachParts = strings.Split(s, ":")
	var hours int
	hours, _ = strconv.Atoi(eachParts[0])
	var secondLeftPart string
	secondLeftPart = eachParts[2][:2]
	var secondRightPart string
	secondRightPart = eachParts[2][2:]
	if secondRightPart == "PM" {
		if hours != 12 {
			hours += 12
		}
	}
	if secondRightPart == "AM" {
		if hours == 12 {
			hours = 0
		}
	}
	fmt.Printf("%02d:%s:%s", hours, eachParts[1], secondLeftPart)
}
