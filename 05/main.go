//Problem description: https://adventofcode.com/2019/day/5

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

type OpCodeAndParamModes struct {
	OpCode     int
	ParamModes []int
}

func add(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

var operations = map[int]func(int, []int){
	1: runAdd,
	2: runMultiply,
	3: runInput,
	4: runOutput,
}

var numOfParametersByOpCode = map[string]int{
	"1": 3,
	"2": 3,
	"3": 1,
	"4": 1,
}

var instructions = []int{}

func partOne() {
	fmt.Println("Part 1 start")
	setupInstructionsFromFile()
	runInstruction(0)
	//fmt.Println(instructions)
}

func runInstruction(address int) {
	if address >= len(instructions)-1 || instructions[address] == 99 {
		return
	}

	ocpm := createOpCodeAndParamModes(instructions[address])
	operation := operations[ocpm.OpCode]
	operation(address+1, ocpm.ParamModes)

	address += len(ocpm.ParamModes) + 1
	runInstruction(address)
}

func runAdd(address int, paramModes []int) {
	params := getParams(address, paramModes, true)
	result := add(params[0], params[1])
	instructions[params[2]] = result
}

func runMultiply(address int, paramModes []int) {
	params := getParams(address, paramModes, true)
	result := multiply(params[0], params[1])
	instructions[params[2]] = result
}

func runInput(address int, paramModes []int) {
	params := getParams(address, paramModes, true)
	input := 1
	instructions[params[0]] = input
}

func runOutput(address int, paramModes []int) {
	params := getParams(address, paramModes, false)
	fmt.Println("output", params[0])
}

func getParams(address int, paramModes []int, willWriteToAddress bool) []int {
	params := make([]int, len(paramModes))
	for i := 0; i < len(paramModes); i++ {
		if willWriteToAddress && i == len(paramModes)-1 {
			params[i] = instructions[address]
		} else {
			params[i] = getParam(address, paramModes[i])
		}
		address++
	}
	return params
}

func getParam(i int, paramMode int) int {
	if paramMode == 1 {
		return instructions[i]
	}
	address := instructions[i]
	return instructions[address]
}

func setupInstructionsFromFile() {
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
}

func createOpCodeAndParamModes(instruction int) OpCodeAndParamModes {
	inst := strconv.Itoa(instruction)
	length := len(inst)
	opCode := inst[length-1 : length]
	pmIndex := length - 3
	numOfParams := numOfParametersByOpCode[opCode]

	var paramModes []int
	for i := 0; i < numOfParams; i++ {
		if pmIndex < 0 {
			paramModes = append(paramModes, 0)
		} else {
			pm, e := strconv.ParseInt(inst[pmIndex:pmIndex+1], 0, 0)
			if e != nil {
				log.Fatal(e)
			}
			paramModes = append(paramModes, int(pm))
		}
		pmIndex--
	}

	opCodeInt64, e := strconv.ParseInt(opCode, 0, 0)
	if e != nil {
		log.Fatal(e)
	}

	ocpm := OpCodeAndParamModes{
		int(opCodeInt64),
		paramModes,
	}
	return ocpm
}
