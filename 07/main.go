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

func main() {
	partOne()
}

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
		intcodecomputer.ResetProgram()
		runAllAmplifiers(phaseSettings)
		output := intcodecomputer.GetOutput()
		if intcodecomputer.GetOutput() > maxOutputSignal {
			fmt.Println("MAX", output)
			maxOutputSignal = output
			maxPhaseSettings = phaseSettings
		}
	}

	return maxOutputSignal, maxPhaseSettings
}

func runAllAmplifiers(phaseSettings []int) {
	for _, phase := range phaseSettings {
		intcodecomputer.ResetInstructions()
		intcodecomputer.UpdateInputs([]int{phase, intcodecomputer.GetOutput()})
		runSingleAmplifier()
	}
}

func runSingleAmplifier() {
	intcodecomputer.Run()
}

func setupInstructionsFromFile() {
	content, err := ioutil.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	inputAsStrings := strings.Split(text, ",")

	var instructions []int
	for i := range inputAsStrings {
		input, e := strconv.ParseInt(inputAsStrings[i], 0, 0)
		if e != nil {
			log.Fatal(e)
		}
		instructions = append(instructions, int(input))
	}

	intcodecomputer.Init(instructions)
}
