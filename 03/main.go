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
	//partOne()
	partTwo()
}

type Point struct {
	X, Y int64
}

type Movement struct {
	Operation (func(Point) Point)
	Steps     int
}

type Distance struct {
	DistancesByPoint map[Point]int
	CurrentDistance  *int
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

var operations = map[string](func(Point) Point){
	"R": right,
	"L": left,
	"U": up,
	"D": down,
}

func move(movements []Movement) ([]Point, Distance) {
	distance := Distance{make(map[Point]int), new(int)}
	var allPoints []Point
	startPoint := Point{0, 0}
	allPoints = append(allPoints, startPoint)
	distance.DistancesByPoint[startPoint] = 0
	for i := range movements {
		addedPointsAfterMovement := moveOnce(startPoint, movements[i], distance)
		allPoints = append(allPoints, addedPointsAfterMovement...)
		startPoint = allPoints[len(allPoints)-1]
	}

	return allPoints, distance
}

func moveOnce(startPoint Point, movement Movement, distance Distance) []Point {
	var newPoints []Point
	for i := 0; i < movement.Steps; i++ {
		*distance.CurrentDistance++
		reachedPoint := movement.Operation(startPoint)
		dist, ok := distance.DistancesByPoint[reachedPoint]
		if ok {
			distance.CurrentDistance = &dist
		} else {
			distance.DistancesByPoint[reachedPoint] = *distance.CurrentDistance
		}
		newPoints = append(newPoints, reachedPoint)
		startPoint = newPoints[len(newPoints)-1]
	}

	return newPoints
}

func createMovement(instruction string) Movement {
	direction := string(instruction[0])
	steps, e := strconv.ParseInt(instruction[1:], 0, 0)
	if e != nil {
		log.Fatal(e)
	}
	operation := operations[direction]
	return Movement{operation, int(steps)}
}

func partOne() {
	fmt.Println("Part 1 start")

	wire1Movements, wire2Movements := getMovementsFromFile()
	wire1Points, _ := move(wire1Movements)
	wire2Points, _ := move(wire2Movements)

	intersections := getIntersections(wire1Points, wire2Points)
	minimumDistance := calculateMinimumDistanceToZeroPoint(intersections)
	fmt.Println("Manhattan distance from central port to closest intersection:", minimumDistance)
}

func partTwo() {
	fmt.Println("Part 2 start")

	wire1Movements, wire2Movements := getMovementsFromFile()
	wire1Points, w1Dist := move(wire1Movements)
	wire2Points, w2Dist := move(wire2Movements)

	intersections := getIntersections(wire1Points, wire2Points)

	minimumTotalStepsToIntersection := getShortestCombinedStepsToIntersection(intersections, w1Dist.DistancesByPoint, w2Dist.DistancesByPoint)
	fmt.Println("Fewest combined steps the wires must take to reach an intersection:", minimumTotalStepsToIntersection)
}

func getShortestCombinedStepsToIntersection(intersections []Point, w1Distances map[Point]int, w2Distances map[Point]int) int {
	minSteps := math.MaxInt64
	for i := range intersections {
		fmt.Println(intersections[i], w1Distances[intersections[i]], w2Distances[intersections[i]])
		combinedSteps := w1Distances[intersections[i]] + w2Distances[intersections[i]]
		if combinedSteps < minSteps {
			minSteps = combinedSteps
		}
	}
	return minSteps
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

func getMovementsFromFile() ([]Movement, []Movement) {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var wireMovements [][]Movement
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), ",")
		var movements []Movement
		for i := range splitLine {
			movements = append(movements, createMovement(splitLine[i]))
		}
		wireMovements = append(wireMovements, movements)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(wireMovements) != 2 {
		log.Fatal("Too many lines read from input")
	}

	return wireMovements[0], wireMovements[1]
}
