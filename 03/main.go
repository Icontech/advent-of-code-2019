//Problem description: https://adventofcode.com/2019/day/3

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
}

type Point struct {
	X, Y int64
}

func right(startPoint Point) Point {
	return Point{
		startPoint.X + 1,
		startPoint.Y,
	}
}

func left(startPoint Point) Point {
	return Point{
		startPoint.X - 1,
		startPoint.Y,
	}
}

func up(startPoint Point) Point {
	return Point{
		startPoint.X,
		startPoint.Y + 1,
	}
}

func down(startPoint Point) Point {
	return Point{
		startPoint.X,
		startPoint.Y - 1,
	}
}

var movements = map[string](func(Point) Point){
	"R": right,
	"L": left,
	"U": up,
	"D": down,
}

func move(instructions []string) []Point {
	var allPoints []Point
	allPoints = append(allPoints, Point{0, 0})
	startPoint := allPoints[0]
	for i := range instructions {
		addedPointsAfterMovement := moveOnce(startPoint, instructions[i])
		allPoints = append(allPoints, addedPointsAfterMovement...)
		startPoint = allPoints[len(allPoints)-1]
	}

	return allPoints
}

func moveOnce(startPoint Point, input string) []Point {
	direction := string(input[0])
	steps, e := strconv.ParseInt(input[1:], 0, 0)
	if e != nil {
		log.Fatal(e)
	}
	movement := movements[direction]

	var newPoints []Point
	for i := 0; i < int(steps); i++ {
		newPoints = append(newPoints, movement(startPoint))
		startPoint = newPoints[len(newPoints)-1]
	}

	return newPoints
}

func partOne() {
	fmt.Println("Part 1 start")

	wire1Instructions, wire2Instructions := getInstructionsFromFile()
	wire1Points := move(wire1Instructions)
	wire2Points := move(wire2Instructions)

	intersections := getIntersections(wire1Points, wire2Points)
	minimumDistance := calculateMinimumDistanceToZeroPoint(intersections)
	fmt.Println("Manhattan distance from central port to closest intersection:", minimumDistance)
}

func getIntersections(w1 []Point, w2 []Point) []Point {
	var intersections []Point
	var start []Point
	var end []Point
	if len(w1) >= len(w2) {
		start = w1
		end = w2
	} else {
		start = w2
		end = w1
	}

	for i := range start {
		for j := range end {
			if start[i] == end[j] && !(start[i].X == 0 && start[i].Y == 0) {
				intersections = append(intersections, start[i])
			}
		}
	}

	return intersections
}

func calculateMinimumDistanceToZeroPoint(points []Point) int {
	minDistance := math.MaxInt64
	for i := range points {
		distance := int(math.Abs(float64(points[i].X)) + math.Abs(float64(points[i].Y)))
		if distance < minDistance {
			minDistance = distance
		}
	}
	return minDistance
}

func getInstructionsFromFile() ([]string, []string) {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), ",")
		lines = append(lines, splitLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(lines) != 2 {
		log.Fatal("Too many lines read from input")
	}

	return lines[0], lines[1]
}
