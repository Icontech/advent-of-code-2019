//Problem description: https://adventofcode.com/2019/day/7

package main

import (
	"fmt"
	"intcodecomputer"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var computers []*intcodecomputer.IntCodeComputer
var instructions []int

func main() {
	partOne()
	//partTwo()
}

func partOne() {
	fmt.Println("Part 1 start")
	instructions = getInstructionsFromFile()
	fmt.Println(instructions)
	comp := intcodecomputer.NewIntCodeComputer(instructions)
	computers = append(computers, comp)
	permutations := createAllPhaseSettingPermutations()
	maxOutput, maxPhaseSettings := findLargestThrustersOutputSignal(permutations)
	fmt.Println("Largest output signal sent to thrusters:", maxOutput, "with phaseSettings:", maxPhaseSettings)
}

func partTwo() {
	fmt.Println("Part 2 start")
	instructions = getInstructionsFromFile()
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
		computers[0].UpdateInstructions(instructions)
		computers[0].Reset()
		runAllAmplifiers(phaseSettings)
		output := computers[0].GetOutput()
		if output > maxOutputSignal {
			fmt.Println("MAX", output)
			maxOutputSignal = output
			maxPhaseSettings = phaseSettings
		}
	}

	return maxOutputSignal, maxPhaseSettings
}

func runAllAmplifiers(phaseSettings []int) {
	for _, phase := range phaseSettings {
		computers[0].UpdateInstructions(instructions)
		computers[0].UpdateInputs([]int{phase, computers[0].GetOutput()})
		runSingleAmplifier()
	}
}

func runSingleAmplifier() {
	computers[0].Run()
}

func getInstructionsFromFile() []int {
	content, err := ioutil.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	inputAsStrings := strings.Split(text, ",")

	var instr []int
	for i := range inputAsStrings {
		input, e := strconv.ParseInt(inputAsStrings[i], 0, 0)
		if e != nil {
			log.Fatal(e)
		}
		instr = append(instr, int(input))
	}

	return instr
}
