//Problem description: https://adventofcode.com/2019/day/7

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

var operations = map[int]func(*int, []int){
	1: runAdd,
	2: runMultiply,
	3: runInput,
	4: runOutput,
	5: runJumpIfTrue,
	6: runJumpIfFalse,
	7: runLessThan,
	8: runEquals,
}

var numOfParametersByOpCode = map[string]int{
	"1": 3,
	"2": 3,
	"3": 1,
	"4": 1,
	"5": 2,
	"6": 2,
	"7": 3,
	"8": 3,
}

var instructions = []int{}
var originalInstructions = []int{}
var inputs = []int{0}
var currentInputIndex = 0
var output int = 0

func partOne() {
	fmt.Println("Part 1 start")
	setupInstructionsFromFile()
	permutations := createAllPhaseSettingPermutations()
	maxOutput, maxPhaseSettings := findLargestThrustersOutputSignal(permutations)
	fmt.Println("Largest output signal sent to thrusters:", maxOutput, "with phaseSettings:", maxPhaseSettings)
}

func createAllPhaseSettingPermutations() *[][]int {
	phaseSettings := []int{4, 3, 2, 1, 0}
	permutations := [][]int{}
	permute(phaseSettings, 0, &permutations)
	return &permutations
}

func permute(phases []int, index int, permutations *[][]int) {
	if index == len(phases)-1 {
		addPermutation(permutations, phases)
	} else {
		for i := index; i < len(phases); i++ {
			a := phases[index]
			b := phases[i]
			phases[index] = b
			phases[i] = a

			permute(phases, index+1, permutations)

			phases[index] = a
			phases[i] = b
		}
	}
}

func addPermutation(permutations *[][]int, phases []int) {
	phaseCopy := make([]int, len(phases))
	for i, v := range phases {
		phaseCopy[i] = v
	}
	*permutations = append(*permutations, phaseCopy)
}

func findLargestThrustersOutputSignal(permutations *[][]int) (int, []int) {
	maxOutputSignal := 0
	maxPhaseSettings := []int{}
	for _, phaseSettings := range *permutations {
		runAllAmplifiers(phaseSettings)
		if output > maxOutputSignal {
			fmt.Println("MAX", output)
			maxOutputSignal = output
			maxPhaseSettings = phaseSettings
		}
	}

	return maxOutputSignal, maxPhaseSettings
}

func runAllAmplifiers(phaseSettings []int) int {
	updateOutput(0)
	for _, phase := range phaseSettings {
		resetInstructions()
		updateInputs([]int{phase, output})
		runSingleAmplifier()
	}
	return output
}

func runSingleAmplifier() {
	address := new(int)
	runInstruction(address)
}

func updateInputs(ints []int) {
	inputs = nil
	for i := range ints {
		inputs = append(inputs, ints[i])
	}
	currentInputIndex = 0
}

func getInput() int {
	input := inputs[currentInputIndex]
	currentInputIndex++
	if currentInputIndex == len(inputs) {
		currentInputIndex = 0
	}
	return input
}

func resetInstructions() {
	for i := range originalInstructions {
		instructions[i] = originalInstructions[i]
	}
}

func updateOutput(value int) {
	output = value
}

func runInstruction(address *int) {
	if *address >= len(instructions)-1 || instructions[*address] == 99 {
		return
	}

	ocpm := createOpCodeAndParamModes(instructions[*address])
	operation := operations[ocpm.OpCode]
	*address++
	operation(address, ocpm.ParamModes)
	runInstruction(address)
}

func runAdd(address *int, paramModes []int) {
	params := getParams(*address, paramModes, true)
	result := add(params[0], params[1])
	instructions[params[2]] = result
	*address += len(paramModes)
}

func runMultiply(address *int, paramModes []int) {
	params := getParams(*address, paramModes, true)
	result := multiply(params[0], params[1])
	instructions[params[2]] = result
	*address += len(paramModes)
}

func runInput(address *int, paramModes []int) {
	params := getParams(*address, paramModes, true)
	instructions[params[0]] = getInput()
	*address += len(paramModes)
}

func runOutput(address *int, paramModes []int) {
	params := getParams(*address, paramModes, false)
	output = params[0]
	fmt.Println("output", output)
	*address += len(paramModes)
}

func runJumpIfTrue(address *int, paramModes []int) {
	params := getParams(*address, paramModes, false)
	if params[0] != 0 {
		*address = params[1]
	} else {
		*address += len(paramModes)
	}
}

func runJumpIfFalse(address *int, paramModes []int) {
	params := getParams(*address, paramModes, false)
	if params[0] == 0 {
		*address = params[1]
	} else {
		*address += len(paramModes)
	}
}

func runLessThan(address *int, paramModes []int) {
	params := getParams(*address, paramModes, true)
	if params[0] < params[1] {
		instructions[params[2]] = 1
	} else {
		instructions[params[2]] = 0
	}
	*address += len(paramModes)
}

func runEquals(address *int, paramModes []int) {
	params := getParams(*address, paramModes, true)
	if params[0] == params[1] {
		instructions[params[2]] = 1
	} else {
		instructions[params[2]] = 0
	}
	*address += len(paramModes)
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
	instructions = nil
	originalInstructions = nil
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
		originalInstructions = append(originalInstructions, int(input))
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
