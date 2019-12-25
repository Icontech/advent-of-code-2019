package intcodecomputer

import (
	"fmt"
	"log"
	"strconv"
)

type opCodeAndParamModes struct {
	opCode     int
	paramModes []int
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

//Init initializes the intcode computer with a program (an array of ints). Must be called before Run func.
func Init(program []int) {
	instructions = nil
	originalInstructions = nil
	for i := range program {
		instructions = append(instructions, program[i])
		originalInstructions = append(originalInstructions, program[i])
	}
}

//Run runs the program initialized with the Init func.
func Run() {
	address := new(int)
	runInstruction(address)
}

//ResetInstructions resets the instructions to their initial state but does not change the inputs/output values.
func ResetInstructions() {
	for i := range originalInstructions {
		instructions[i] = originalInstructions[i]
	}
}

//ResetProgram resets the instructions, inputs and output to their initial state.
func ResetProgram() {
	ResetInstructions()
	output = 0
	inputs = []int{0}
}

//UpdateInputs adds new values to be used for the input operation. For each input operation, the index of the array will be incremented by 1.
func UpdateInputs(ints []int) {
	inputs = nil
	for i := range ints {
		inputs = append(inputs, ints[i])
	}
	currentInputIndex = 0
}

// GetOutput returns the current value of the output variable
func GetOutput() int {
	return output
}

func getInput() int {
	input := inputs[currentInputIndex]
	currentInputIndex++
	if currentInputIndex == len(inputs) {
		currentInputIndex = 0
	}
	return input
}

func updateOutput(value int) {
	output = value
}

func runInstruction(address *int) {
	if *address >= len(instructions)-1 || instructions[*address] == 99 {
		return
	}

	ocpm := createopCodeAndParamModes(instructions[*address])
	operation := operations[ocpm.opCode]
	*address++
	operation(address, ocpm.paramModes)
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

func createopCodeAndParamModes(instruction int) opCodeAndParamModes {
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

	ocpm := opCodeAndParamModes{
		int(opCodeInt64),
		paramModes,
	}
	return ocpm
}
