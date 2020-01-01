//Problem description: https://adventofcode.com/2019/day/9

package main

import (
	"fmt"
	"intcodecomputer"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var amplifiers []*intcodecomputer.IntCodeComputer
var instructions []int64

func main() {
	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("Part 1 start")
	setupInstructionsFromFile()
	runBoostProgramTest()
}

func partTwo() {
	fmt.Println("Part 2 start")
	setupInstructionsFromFile()
	runBoostProgramInSensorBoostMode()
}

func runBoostProgramTest() {
	icc := intcodecomputer.NewIntCodeComputer(instructions, false, "computer")
	icc.UpdateInputs([]int64{1})
	icc.Run()
}

func runBoostProgramInSensorBoostMode() {
	icc := intcodecomputer.NewIntCodeComputer(instructions, false, "computer")
	icc.UpdateInputs([]int64{2})
	icc.Run()
}

func setupInstructionsFromFile() {
	content, err := ioutil.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	inputAsStrings := strings.Split(text, ",")

	instructions = nil
	for i := range inputAsStrings {
		input, e := strconv.ParseInt(inputAsStrings[i], 0, 0)
		if e != nil {
			log.Fatal(e)
		}
		instructions = append(instructions, input)
	}
}
