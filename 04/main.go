//Problem description: https://adventofcode.com/2019/day/4

package main

import (
	"fmt"
	"strconv"
)

func main() {
	partOne()
	partTwo()
}

const startRange int = 197487
const endRange int = 673251

func partOne() {
	fmt.Println("Part 1 start")
	passwordCount := 0
	for i := startRange; i <= endRange; i++ {
		if isValidPasswordPartOne(strconv.Itoa(i), 0, false) {
			passwordCount++
		}
	}

	fmt.Println("Number of different passwords:", passwordCount)
}

func isValidPasswordPartOne(number string, index int, doubleFound bool) bool {
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
	return isValidPasswordPartOne(number, index, doubleFound)
}

func partTwo() {
	fmt.Println("Part 2 start")
	passwordCount := 0
	for i := startRange; i <= endRange; i++ {
		if isValidPasswordPartTwo(strconv.Itoa(i)) {
			passwordCount++
		}
	}

	fmt.Println("Number of different passwords:", passwordCount)
}

type DigitTracker struct {
	Digit, Count int64
}

func isValidPasswordPartTwo(number string) bool {
	var digitTracker DigitTracker
	doubleFound := false
	for i := 0; i < len(number); i++ {
		if i == 0 {
			firstDigit, _ := strconv.ParseInt(number[0:1], 0, 0)
			digitTracker.Digit = firstDigit
			digitTracker.Count = 1
		} else {
			currentDigit, _ := strconv.ParseInt(number[i:i+1], 0, 0)

			if digitTracker.Digit > currentDigit {
				return false
			}

			if doubleFound {
				digitTracker.Digit = currentDigit
			} else {
				if digitTracker.Digit == currentDigit {
					digitTracker.Count++
					if i == len(number)-1 && digitTracker.Count == 2 {
						doubleFound = true
					}
				} else {
					if digitTracker.Count == 2 {
						doubleFound = true
					}
					digitTracker.Digit = currentDigit
					digitTracker.Count = 1
				}
			}
		}
	}
	return doubleFound
}
