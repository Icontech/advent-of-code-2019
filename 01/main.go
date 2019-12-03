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
}

func partOne() {
	fmt.Println("Calculating fuel requirement...")
	masses := getMassesFromFile()
	fuelRequired := int64(0)
	for _, mass := range masses {
		fuelRequired += calculateFuel(mass)
	}

	fmt.Println("Done!")
	fmt.Println("The fuel requirement is", fuelRequired)
}

func calculateFuel(mass int64) int64 {
	x := float64(mass / 3)
	return int64(math.Floor(x) - 2)
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
