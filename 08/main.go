//Problem description: https://adventofcode.com/2019/day/8

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
)

const imageWidth = 25
const imageHeight = 6

var layerSize = imageWidth * imageHeight

var digits []int

func main() {
	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("Part 1 start")
	setupDigitsFromFile()
	layerWithLeastZeros := getLayerWithLeastZeros()
	onesMultipliedByTwos := calculateOnesMultipliedByTwos(layerWithLeastZeros)
	fmt.Println("Layer number with least zeros:", layerWithLeastZeros)
	fmt.Println("Number of 1 digits multiplied by the number of 2 digits:", onesMultipliedByTwos)
}

func getLayerWithLeastZeros() int {
	layerWithLeastZeros := 0
	minZeroCount := math.MaxInt64
	currentZeroCount := 0
	currentLayer := -1
	for i := range digits {
		mod := mod(i, layerSize)
		if mod == 0 {
			currentLayer++
			currentZeroCount = 0
		}

		if digits[i] == 0 {
			currentZeroCount++
		}

		if mod == layerSize-1 && currentZeroCount < minZeroCount {
			layerWithLeastZeros = currentLayer
			minZeroCount = currentZeroCount
		}

	}
	return layerWithLeastZeros
}

func printLayer(layerNumber int) {
	startIndex := layerNumber * layerSize
	endIndex := startIndex + layerSize
	for i := startIndex; i < endIndex; i++ {
		if i == endIndex-1 {
			fmt.Println(digits[i])
		} else {
			fmt.Print(digits[i])
		}
	}
}

func calculateOnesMultipliedByTwos(layerNumber int) int {
	startIndex := layerNumber * layerSize
	endIndex := startIndex + layerSize
	onesCount := 0
	twosCount := 0
	for i := startIndex; i < endIndex; i++ {
		if digits[i] == 1 {
			onesCount++
		} else if digits[i] == 2 {
			twosCount++
		}
	}

	return onesCount * twosCount
}

func partTwo() {
	fmt.Println("Part 2 start")
	setupDigitsFromFile()
	decodedImage := getDecodedImage()
	fmt.Println("Decoding image...")
	printImage(decodedImage)
	fmt.Println("Message produced after decoding image: ACKPZ")
}

func getDecodedImage() []int {
	decodedImage := []int{}
	for i := 0; i < layerSize; i++ {
		pixelColor := getPixelColor(i)
		decodedImage = append(decodedImage, pixelColor)
	}
	return decodedImage
}

func getPixelColor(pixel int) int {
	layerCount := len(digits) / layerSize
	for i := 0; i < layerCount; i++ {
		color := digits[pixel+(i*layerSize)]
		if color != 2 {
			return color
		}
	}
	return 2
}

func printImage(image []int) {
	for i := 0; i < len(image); i++ {
		mod := mod(i, imageWidth)
		if mod == imageWidth-1 {
			fmt.Println(image[i])
		} else {
			fmt.Print(image[i])
		}
	}
}

func setupDigitsFromFile() {
	content, err := ioutil.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)

	for i := 0; i < len(text); i++ {
		digit, e := strconv.ParseInt(text[i:i+1], 0, 0)
		if e != nil {
			log.Fatal(e)
		}
		digits = append(digits, int(digit))
	}
}

func mod(a, b int) int {
	return (a%b + b) % b
}
