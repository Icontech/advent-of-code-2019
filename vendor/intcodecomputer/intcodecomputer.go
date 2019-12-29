package intcodecomputer

import (
	"fmt"
	"log"
	"strconv"
)

//IntCodeComputer struct
type IntCodeComputer struct {
	name                   string
	instructions           []int64
	address                int
	inputs                 []int64
	currentInputIndex      int
	output                 int64
	shouldPauseAfterOutput bool
	isPaused               bool
	isHalted               bool
}

//NewIntCodeComputer creates a new IntCodeComputer
func NewIntCodeComputer(instructions []int64, shouldPauseAfterOutput bool, name string) *IntCodeComputer {
	icc := IntCodeComputer{
		inputs:                 []int64{0},
		instructions:           instructions,
		shouldPauseAfterOutput: shouldPauseAfterOutput,
		name:                   name}
	return &icc
}

type opCodeAndParamModes struct {
	opCode     int
	paramModes []int
}

func add(a int64, b int64) int64 {
	return a + b
}

func multiply(a int64, b int64) int64 {
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
func (icc *IntCodeComputer) UpdateInstructions(instr []int64) {
	icc.instructions = nil
	for i := 0; i < len(instr); i++ {
		icc.instructions = append(icc.instructions, instr[i])
	}
	icc.address = 0
}

//Reset resets all variables to their initial state.
func (icc *IntCodeComputer) Reset() {
	icc.output = 0
	icc.instructions = []int64{0}
	icc.address = 0
	icc.instructions = nil
	icc.isHalted = false
	icc.isPaused = false
}

//UpdateInputs adds new values to be used for the input operation. For each input operation, the index of the array will be incremented by 1.
func (icc *IntCodeComputer) UpdateInputs(inputs []int64) {
	icc.inputs = inputs
	icc.currentInputIndex = 0
}

// GetOutput returns the current value of the output variable
func (icc *IntCodeComputer) GetOutput() int64 {
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

//GetInstruction returns the instruction in the provided address and a true value, if the address is within range. Otherwise returns false and 0.
func (icc *IntCodeComputer) GetInstruction(address int) (bool, int64) {
	if icc.instructions == nil {
		return false, 0
	}

	if 0 <= address && address < len(icc.instructions) {
		return true, icc.instructions[address]
	}

	return false, 0
}

func (icc *IntCodeComputer) getInput() int64 {
	input := icc.inputs[icc.currentInputIndex]
	icc.currentInputIndex++
	if icc.currentInputIndex == len(icc.inputs) {
		icc.currentInputIndex = 0
	}
	return input
}

func (icc *IntCodeComputer) updateOutput(value int64) {
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
	fmt.Println(icc.name, "input:", input)
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
		icc.address = int(params[1])
	} else {
		icc.address += len(paramModes)
	}
}

func (icc *IntCodeComputer) runJumpIfFalse(paramModes []int) {
	params := icc.getParams(paramModes, false)
	if params[0] == 0 {
		icc.address = int(params[1])
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

func (icc *IntCodeComputer) getParams(paramModes []int, willWriteToAddress bool) []int64 {
	address := icc.address
	params := make([]int64, len(paramModes))
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

func getParam(i int, instructions []int64, paramMode int) int64 {
	if paramMode == 1 {
		return instructions[i]
	}
	address := instructions[i]
	return instructions[address]
}

func createOpCodeAndParamModes(instruction int64) opCodeAndParamModes {
	inst := strconv.FormatInt(instruction, 10)
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
