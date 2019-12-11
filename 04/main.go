//Problem description: https://adventofcode.com/2019/day/4

package main

import (
	"fmt"
	"strconv"
)

func main() {
	partOne()
}

func partOne() {
	fmt.Println("Part 1 start")
	startRange := 197487
	endRange := 673251
	passwordCount := 0
	for i := startRange; i <= endRange; i++ {
		if isValidPassword(strconv.Itoa(i), 0, false) {
			passwordCount++
		}
	}

	fmt.Println("Number of different passwords:", passwordCount)
}

func isValidPassword(number string, index int, doubleFound bool) bool {
	if index == len(number)-1 {
		return doubleFound
	}

	digit1, _ := strconv.ParseInt(number[index:index+1], 0, 0)
	digit2, _ := strconv.ParseInt(number[index+1:index+2], 0, 0)

	if digit1 > digit2 {
		return false
	}

	if digit1 == digit2 {
		doubleFound = true
	}

	index++
	return isValidPassword(number, index, doubleFound)
}
