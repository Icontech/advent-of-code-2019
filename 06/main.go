//Problem description: https://adventofcode.com/2019/day/6

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

type Node struct {
	Name     string
	Children []*Node
	Parent   *Node
}

var tree map[string]*Node

var distancesByName map[string]int

func partOne() {
	fmt.Println("Part 1 start")
	createTreeFromFile()
	totalOrbitCount := runOrbitCalculation()
	fmt.Println("Total number of direct and indirect orbits:", totalOrbitCount)
}

func runOrbitCalculation() int {
	startObj := getNode("COM")
	totalOrbitCount := new(int)
	for i := range startObj.Children {
		calculateObjectOrbits(startObj.Children[i], totalOrbitCount, 0)
	}
	return *totalOrbitCount
}

func calculateObjectOrbits(node *Node, totalOrbitCount *int, count int) {
	count++
	*totalOrbitCount += count

	for i := range node.Children {
		calculateObjectOrbits(node.Children[i], totalOrbitCount, count)
	}
}

func partTwo() {
	fmt.Println("Part 2 start")
	createTreeFromFile()
	shortestDistance := runOrbitTransferCalculation()
	if shortestDistance == -1 {
		fmt.Println("No minimum number of orbital transfers found")
	} else {
		fmt.Println("Minimum number of orbital transfers:", shortestDistance)
	}
}

//Dijkstra's algorithm to find shortest distance from source tot target
func runOrbitTransferCalculation() int {
	source := getNode("YOU").Parent
	target := getNode("SAN").Parent

	distancesByName = make(map[string]int)
	unfinishedNodes := make(map[string]*Node)
	for key := range tree {
		distancesByName[key] = math.MaxInt64
		unfinishedNodes[key] = tree[key]
	}
	distancesByName[source.Name] = 0

	for len(unfinishedNodes) > 0 {
		node := getMinDistanceNode(unfinishedNodes)
		delete(unfinishedNodes, node.Name)

		if node.Name == target.Name {
			return distancesByName[target.Name]
		}

		for i := range node.Children {
			alt := distancesByName[node.Name] + 1
			if alt < distancesByName[node.Children[i].Name] {
				distancesByName[node.Children[i].Name] = alt
			}
		}

		if node.Parent != nil {
			alt := distancesByName[node.Name] + 1
			if alt < distancesByName[node.Parent.Name] {
				distancesByName[node.Parent.Name] = alt
			}
		}
	}

	return -1
}

func getMinDistanceNode(nodes map[string]*Node) *Node {
	minDist := math.MaxInt64
	var minDistNode *Node = nil
	for key := range nodes {
		if distancesByName[key] < minDist {
			minDist = distancesByName[key]
			minDistNode = nodes[key]
		}
	}
	return minDistNode
}

func createTreeFromFile() {
	tree = make(map[string]*Node)
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parentAndChild := strings.Split(scanner.Text(), ")")
		addOrUpdateNode(parentAndChild[0], parentAndChild[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func addOrUpdateNode(name string, childName string) {
	node := getNode(name)
	if node == nil {
		node = addNode(name, Node{name, nil, nil})
	}

	child := getNode(childName)
	if child == nil {
		child = addNode(childName, Node{childName, nil, node})
	}
	child.Parent = node
	node.Children = append(node.Children, child)
}

func getNode(name string) *Node {
	node, ok := tree[name]
	if ok {
		return node
	}
	return nil
}

func addNode(name string, node Node) *Node {
	tree[name] = &node
	return &node
}

func printTree(node *Node) {
	fmt.Println("NAME", node.Name)
	if node.Parent != nil {
		fmt.Println("PARENT", node.Parent.Name)
	} else {
		fmt.Println("NO PARENT")
	}
	if len(node.Children) > 0 {
		fmt.Println("CHILDREN")
		for i := range node.Children {
			fmt.Println(node.Children[i].Name)
		}

		for i := range node.Children {
			printTree(node.Children[i])
		}
	} else {
		fmt.Println("NO CHILDREN")
	}
}
