//Problem description: https://adventofcode.com/2019/day/2

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

var instructions []int64

func partOne() {
	fmt.Println("Part 1 start")

	setupInstructionsFromFile()
	instructions[1] = 12
	instructions[2] = 2
	icc := intcodecomputer.NewIntCodeComputer(instructions, false, "computer")
	icc.Run()
	ok, value := icc.GetInstruction(0)
	if !ok {
		log.Fatal("0 out of range")
	}
	fmt.Println("Value in position 0 after program halts:", value)
}

func partTwo() {
	fmt.Println("Part 2 start")

	setupInstructionsFromFile()
	expectedOutput := int64(19690720)
	found, noun, verb := findNounAndVerb(expectedOutput)

	if !found {
		log.Fatal("No noun and verb found for the expected output")
	}

	fmt.Println("Noun", noun, "and verb", verb, "produce output", expectedOutput)
	result := 100*noun + verb
	fmt.Println("100 *", noun, "+", verb, "=", result)
}

func findNounAndVerb(output int64) (bool, int, int) {
	icc := intcodecomputer.NewIntCodeComputer(instructions, false, "computer")
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			icc.Reset()
			instructions[1] = int64(noun)
			instructions[2] = int64(verb)
			icc.UpdateInstructions(instructions)
			icc.Run()

			ok, value := icc.GetInstruction(0)
			if !ok {
				log.Fatal("0 out of range")
			}

			if value == output {
				return true, noun, verb
			}
		}
	}
	return false, 0, 0
}

func setupInstructionsFromFile() {
	instructions = nil
	content, err := ioutil.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	inputAsStrings := strings.Split(text, ",")

	for _, str := range inputAsStrings {
		input, e := strconv.ParseInt(str, 0, 0)
		if e != nil {
			log.Fatal(err)
		}
		instructions = append(instructions, input)
	}
}
