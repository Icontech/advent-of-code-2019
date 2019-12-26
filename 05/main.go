//Problem description: https://adventofcode.com/2019/day/5

package main

import (
	"fmt"
	"intcodecomputer"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("Part 1 start")
	instructions := getInstructionsFromFile()
	computer := intcodecomputer.NewIntCodeComputer(instructions, false)
	computer.UpdateInputs([]int{1})
	computer.Run()
}

func partTwo() {
	fmt.Println("Part 2 start")
	instructions := getInstructionsFromFile()
	computer := intcodecomputer.NewIntCodeComputer(instructions, false)
	computer.UpdateInputs([]int{5})
	computer.Run()
}

func getInstructionsFromFile() []int {
	var instructions []int
	content, err := ioutil.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	inputAsStrings := strings.Split(text, ",")

	for i := range inputAsStrings {
		input, e := strconv.ParseInt(inputAsStrings[i], 0, 0)
		if e != nil {
			log.Fatal(e)
		}
		instructions = append(instructions, int(input))
	}

	return instructions
}
