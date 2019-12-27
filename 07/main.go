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

var amplifiers []*intcodecomputer.IntCodeComputer
var instructions []int

func main() {
	partOne()
	//partTwo()
}

func partOne() {
	fmt.Println("Part 1 start")
	instructions = getInstructionsFromFile()
	phaseSettings := []int{4, 3, 2, 1, 0}
	setupAmplifiers(false, len(phaseSettings))
	permutations := createAllPhaseSettingPermutations(phaseSettings)
	maxOutput, maxPhaseSettings := findLargestThrustersOutputSignal(permutations)
	fmt.Println("Largest output signal sent to thrusters:", maxOutput, "with phaseSettings:", maxPhaseSettings)
}

func setupAmplifiers(shouldPauseOnOutput bool, numOfAmplifiers int) {
	for i := 0; i < numOfAmplifiers; i++ {
		icc := intcodecomputer.NewIntCodeComputer(instructions, shouldPauseOnOutput, "amp"+strconv.Itoa(i))
		amplifiers = append(amplifiers, icc)
	}
}

func partTwo() {
	fmt.Println("Part 2 start")
	instructions = getInstructionsFromFile()
	phaseSettings := []int{4, 3, 2, 1, 0}
	setupAmplifiers(false, len(phaseSettings))
	permutations := createAllPhaseSettingPermutations(phaseSettings)
	maxOutput, maxPhaseSettings := findLargestThrustersOutputSignal(permutations)
	fmt.Println("Largest output signal sent to thrusters:", maxOutput, "with phaseSettings:", maxPhaseSettings)
}

func createAllPhaseSettingPermutations(phaseSettings []int) *[][]int {
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
		resetAllAmplifiers()
		runAllAmplifiersOnce(phaseSettings)
		output := amplifiers[len(amplifiers)-1].GetOutput()
		if output > maxOutputSignal {
			fmt.Println("MAX", output)
			maxOutputSignal = output
			maxPhaseSettings = phaseSettings
		}
	}

	return maxOutputSignal, maxPhaseSettings
}

func runAllAmplifiersOnce(phaseSettings []int) {
	for i, phase := range phaseSettings {
		prevAmpIndex := mod(i-1, len(phaseSettings))
		amplifiers[i].UpdateInputs([]int{phase, amplifiers[prevAmpIndex].GetOutput()})
		amplifiers[i].Run()
	}
}

func runAllAmplifiersWithFeedbackLoop(phaseSettings []int) {
	for i, phase := range phaseSettings {
		prevAmpIndex := mod(i-1, len(phaseSettings))
		amplifiers[i].UpdateInputs([]int{phase, amplifiers[prevAmpIndex].GetOutput()})
		amplifiers[i].Run()
	}

	for i := range amplifiers {
		fmt.Println(amplifiers[i].IsPaused())
	}

}

func resetAllAmplifiers() {
	for i := range amplifiers {
		amplifiers[i].Reset()
		amplifiers[i].UpdateInstructions(instructions)
	}
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

func mod(a, b int) int {
	return (a%b + b) % b
}
