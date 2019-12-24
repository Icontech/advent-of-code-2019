//Problem description: https://adventofcode.com/2019/day/7

package main

import (
	"fmt"
)

func main() {
	partOne()
}

func partOne() {
	fmt.Println("Part 1 start")
	permutations := createAllPhaseSettingPermutations()
	fmt.Println(permutations)
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
