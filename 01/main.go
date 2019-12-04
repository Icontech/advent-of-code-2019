package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("Part 1 start")

	masses := getMassesFromFile()
	fuelRequired := int64(0)
	for _, mass := range masses {
		fuelRequired += calculateFuelPart1(mass)
	}

	fmt.Println("Fuel requirement:", fuelRequired)
}

func calculateFuelPart1(mass int64) int64 {
	x := float64(mass / 3)
	return int64(math.Floor(x) - 2)
}

func partTwo() {
	fmt.Println("Part 2 start")

	masses := getMassesFromFile()
	fuelRequired := int64(0)
	for _, mass := range masses {
		fuelRequired += calculateFuelPart2(mass, 0)
	}

	fmt.Println("Fuel requirement:", fuelRequired)
}

func calculateFuelPart2(mass int64, fuelRequirement int64) int64 {
	x := float64(mass / 3)
	mass = int64(math.Floor(x) - 2)

	if mass < 1 {
		return fuelRequirement
	}
	fuelRequirement += mass

	return calculateFuelPart2(mass, fuelRequirement)
}

func getMassesFromFile() []int64 {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var masses []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, e := strconv.ParseInt(scanner.Text(), 0, 0)
		if e != nil {
			log.Fatal(err)
		}
		masses = append(masses, mass)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return masses
}
