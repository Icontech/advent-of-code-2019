//Problem description: https://adventofcode.com/2019/day/5

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	partOne()
}

type Object struct {
	Name     string
	Children []int
}

var objects []Object

func partOne() {
	fmt.Println("Part 1 start")
	createObjectsFromFile()
	totalOrbitCount := runOrbitCalculation()
	fmt.Println("Total number of direct and indirect orbits:", totalOrbitCount)
}

func createObjectsFromFile() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parentAndChild := strings.Split(scanner.Text(), ")")
		addOrUpdateObject(parentAndChild[0], parentAndChild[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func addOrUpdateObject(name string, childName string) {
	objIndex := getObjectIndex(name)
	if objIndex == -1 {
		objects = append(objects, Object{name, []int{}})
		objIndex = len(objects) - 1
	}

	childIndex := getObjectIndex(childName)
	if childIndex == -1 {
		objects = append(objects, Object{childName, []int{}})
		childIndex = len(objects) - 1
	}
	objects[objIndex].Children = append(objects[objIndex].Children, childIndex)
}

func getObjectIndex(name string) int {
	for i := range objects {
		if objects[i].Name == name {
			return i
		}
	}
	return -1
}

func runOrbitCalculation() int {
	comIndex := getObjectIndex("COM")
	if comIndex == -1 {
		log.Fatal("COM object not found")
	}

	startObj := objects[comIndex]
	totalOrbitCount := new(int)
	for i := range startObj.Children {
		calculateObjectOrbits(startObj.Children[i], totalOrbitCount, 0)
	}
	return *totalOrbitCount
}

func calculateObjectOrbits(index int, totalOrbitCount *int, count int) {
	count++
	*totalOrbitCount += count

	object := objects[index]

	for i := range object.Children {
		calculateObjectOrbits(object.Children[i], totalOrbitCount, count)
	}
}
