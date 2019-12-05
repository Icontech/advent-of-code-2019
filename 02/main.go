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
}

func add(a int64, b int64) int64 {
	return a + b
}

func multiply(a int64, b int64) int64 {
	return a * b
}

var operations = map[int64](func(int64, int64) int64){
	1: add,
	2: multiply,
}

func partOne() {
	fmt.Println("Part 1 start")

	inputs := getInputFromFile()
	inputs[1] = 12
	inputs[2] = 2
	runGravityAssistProgram(inputs)

	fmt.Println("Value in position 0 after program halts:", inputs[0])
}

func runGravityAssistProgram(inputs []int64) {
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

func getInputFromFile() []int64 {
	content, err := ioutil.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	inputAsStrings := strings.Split(text, ",")
	var inputAsInts []int64

	for _, str := range inputAsStrings {
		input, e := strconv.ParseInt(str, 0, 0)
		if e != nil {
			log.Fatal(err)
		}
		inputAsInts = append(inputAsInts, input)
	}

	return inputAsInts
}
