//Problem description: https://adventofcode.com/2019/day/2

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func add(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

var operations = map[int](func(int, int) int){
	1: add,
	2: multiply,
}

var originalInput []int

func partOne() {
	fmt.Println("Part 1 start")

	inputs := getInputFromFile()
	inputs[1] = 12
	inputs[2] = 2
	runGravityAssistProgram(inputs)

	fmt.Println("Value in position 0 after program halts:", inputs[0])
}

func runGravityAssistProgram(inputs []int) {
	i := 0
	roof := len(inputs) - 4
	for i < roof {
		opcode := inputs[i]
		if opcode == 99 {
			break
		}
		operation, ok := operations[opcode]
		if !ok {
			fmt.Println("No operation found for", opcode)
			return
		}
		inPos1 := inputs[i+1]
		inPos2 := inputs[i+2]
		outPos := inputs[i+3]
		inputs[outPos] = operation(inputs[inPos1], inputs[inPos2])
		i += 4
	}
}

func partTwo() {
	fmt.Println("Part 2 start")

	inputs := setupInputs()
	expectedOutput := 19690720
	found, noun, verb := findNounAndVerb(inputs, expectedOutput)

	if !found {
		log.Fatal("No noun and verb found for the expected output")
	}

	fmt.Println("Noun", noun, "and verb", verb, "produce output", expectedOutput)
	result := 100*noun + verb
	fmt.Println("100 *", noun, "+", verb, "=", result)
}

func findNounAndVerb(inputs []int, output int) (bool, int, int) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			resetInputs(inputs)
			inputs[1] = noun
			inputs[2] = verb
			runGravityAssistProgram(inputs)
			if inputs[0] == output {
				return true, noun, verb
			}
		}
	}
	return false, 0, 0
}

func setupInputs() []int {
	originalInput = getInputFromFile()
	var inputs []int
	for i := range originalInput {
		inputs = append(inputs, originalInput[i])
	}
	return inputs
}

func resetInputs(inputs []int) {
	if len(originalInput) != len(inputs) {
		log.Fatal("length of inputs not matching")
	}
	for i := range originalInput {
		inputs[i] = originalInput[i]
	}
}

func getInputFromFile() []int {
	content, err := ioutil.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	inputAsStrings := strings.Split(text, ",")
	var inputAsInts []int

	for _, str := range inputAsStrings {
		input, e := strconv.ParseInt(str, 0, 0)
		if e != nil {
			log.Fatal(err)
		}
		inputAsInts = append(inputAsInts, int(input))
	}

	return inputAsInts
}
