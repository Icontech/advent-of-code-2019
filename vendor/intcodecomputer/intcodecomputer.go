package intcodecomputer

import (
	"fmt"
	"log"
	"strconv"
)

//IntCodeComputer struct
type IntCodeComputer struct {
	name                   string
	instructions           []int
	address                int
	inputs                 []int
	currentInputIndex      int
	output                 int
	shouldPauseAfterOutput bool
	isPaused               bool
	isHalted               bool
}

//NewIntCodeComputer creates a new IntCodeComputer
func NewIntCodeComputer(instructions []int, shouldPauseAfterOutput bool, name string) *IntCodeComputer {
	icc := IntCodeComputer{
		inputs:                 []int{0},
		instructions:           instructions,
		shouldPauseAfterOutput: shouldPauseAfterOutput,
		name:                   name}
	return &icc
}

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

var operations = map[int]func(*IntCodeComputer, []int){
	1: (*IntCodeComputer).runAdd,
	2: (*IntCodeComputer).runMultiply,
	3: (*IntCodeComputer).runInput,
	4: (*IntCodeComputer).runOutput,
	5: (*IntCodeComputer).runJumpIfTrue,
	6: (*IntCodeComputer).runJumpIfFalse,
	7: (*IntCodeComputer).runLessThan,
	8: (*IntCodeComputer).runEquals,
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

//Run runs the program initialized with the Init func.
func (icc *IntCodeComputer) Run() {
	icc.runInstruction()
}

//UpdateInstructions updates the instructions used by the program and sets the address to 0.
func (icc *IntCodeComputer) UpdateInstructions(instr []int) {
	icc.instructions = instr
	icc.address = 0
}

//Reset resets all variables to their initial state.
func (icc *IntCodeComputer) Reset() {
	icc.output = 0
	icc.instructions = []int{0}
	icc.address = 0
	icc.instructions = nil
	icc.isHalted = false
	icc.isPaused = false
}

//UpdateInputs adds new values to be used for the input operation. For each input operation, the index of the array will be incremented by 1.
func (icc *IntCodeComputer) UpdateInputs(inputs []int) {
	icc.inputs = inputs
	icc.currentInputIndex = 0
}

// GetOutput returns the current value of the output variable
func (icc *IntCodeComputer) GetOutput() int {
	return icc.output
}

//Pause pauses the currently running program.
func (icc *IntCodeComputer) Pause() {
	if !icc.isPaused {
		fmt.Println(icc.name, "paused")
		icc.isPaused = true
	}
}

//Resume resumes the currently paused program.
func (icc *IntCodeComputer) Resume() {
	if icc.isPaused {
		icc.isPaused = false
		fmt.Println(icc.name, "resumed")
		icc.runInstruction()
	}
}

//IsPaused returns true if the program has been paused and false otherwise.
func (icc *IntCodeComputer) IsPaused() bool {
	return icc.isPaused
}

//IsHalted returns true if the program has been halted and false otherwise.
func (icc *IntCodeComputer) IsHalted() bool {
	return icc.isHalted
}

func (icc *IntCodeComputer) getInput() int {
	input := icc.inputs[icc.currentInputIndex]
	icc.currentInputIndex++
	if icc.currentInputIndex == len(icc.inputs) {
		icc.currentInputIndex = 0
	}
	return input
}

func (icc *IntCodeComputer) updateOutput(value int) {
	icc.output = value
}

func (icc *IntCodeComputer) runInstruction() {
	if icc.isPaused {
		return
	}

	if icc.instructions[icc.address] == 99 {
		icc.isHalted = true
		fmt.Println(icc.name, "halted")
		return
	}

	ocpm := createOpCodeAndParamModes(icc.instructions[icc.address])
	operation := operations[ocpm.opCode]
	icc.address++
	operation(icc, ocpm.paramModes)
	icc.runInstruction()
}

func (icc *IntCodeComputer) runAdd(paramModes []int) {
	params := icc.getParams(paramModes, true)
	result := add(params[0], params[1])
	icc.instructions[params[2]] = result
	icc.address += len(paramModes)
}

func (icc *IntCodeComputer) runMultiply(paramModes []int) {
	params := icc.getParams(paramModes, true)
	result := multiply(params[0], params[1])
	icc.instructions[params[2]] = result
	icc.address += len(paramModes)
}

func (icc *IntCodeComputer) runInput(paramModes []int) {
	params := icc.getParams(paramModes, true)
	input := icc.getInput()
	fmt.Println(icc.name, "input", input)
	icc.instructions[params[0]] = input
	icc.address += len(paramModes)
}

func (icc *IntCodeComputer) runOutput(paramModes []int) {
	params := icc.getParams(paramModes, false)
	icc.output = params[0]
	fmt.Println(icc.name, "output:", icc.output)
	icc.address += len(paramModes)
	if icc.shouldPauseAfterOutput {
		icc.Pause()
	}
}

func (icc *IntCodeComputer) runJumpIfTrue(paramModes []int) {
	params := icc.getParams(paramModes, false)
	if params[0] != 0 {
		icc.address = params[1]
	} else {
		icc.address += len(paramModes)
	}
}

func (icc *IntCodeComputer) runJumpIfFalse(paramModes []int) {
	params := icc.getParams(paramModes, false)
	if params[0] == 0 {
		icc.address = params[1]
	} else {
		icc.address += len(paramModes)
	}
}

func (icc *IntCodeComputer) runLessThan(paramModes []int) {
	params := icc.getParams(paramModes, true)
	if params[0] < params[1] {
		icc.instructions[params[2]] = 1
	} else {
		icc.instructions[params[2]] = 0
	}
	icc.address += len(paramModes)
}

func (icc *IntCodeComputer) runEquals(paramModes []int) {
	params := icc.getParams(paramModes, true)
	if params[0] == params[1] {
		icc.instructions[params[2]] = 1
	} else {
		icc.instructions[params[2]] = 0
	}
	icc.address += len(paramModes)
}

func (icc *IntCodeComputer) getParams(paramModes []int, willWriteToAddress bool) []int {
	address := icc.address
	params := make([]int, len(paramModes))
	for i := 0; i < len(paramModes); i++ {
		if willWriteToAddress && i == len(paramModes)-1 {
			params[i] = icc.instructions[address]
		} else {
			params[i] = getParam(address, icc.instructions, paramModes[i])
		}
		address++
	}
	return params
}

func getParam(i int, instructions []int, paramMode int) int {
	if paramMode == 1 {
		return instructions[i]
	}
	address := instructions[i]
	return instructions[address]
}

func createOpCodeAndParamModes(instruction int) opCodeAndParamModes {
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
