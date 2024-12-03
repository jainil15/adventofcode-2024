package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFile(filename string) string {
	inputBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln("Error reading the file")
	}
	return string(inputBytes)
}

func getInput(input string) ([]int, []int) {
	rows := strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	left := make([]int, len(rows))
	right := make([]int, len(rows))
	for i, row := range rows {
		s := strings.Split(row, "  ")
		leftN, err := strconv.Atoi(strings.Trim(s[0], " "))
		rightN, err := strconv.Atoi(strings.Trim(s[1], " "))
		if err != nil {
			log.Fatalln("Error converting number", s[0], "Herer")
		}
		left[i] = leftN
		right[i] = rightN
	}
	return left, right
}

func calcPart1(left, right []int) {
	slices.Sort(left)
	slices.Sort(right)
	sum := 0
	for i := range left {
		s := left[i] - right[i]
		if s < 0 {
			sum -= s
		} else {
			sum += s
		}
	}
	fmt.Println(sum)
}

func calcPart2(left, right []int) {
	count := make(map[int]int)
	for _, r := range right {
		count[r]++
	}
	sum := 0
	for _, l := range left {
		sum += l * count[l]
	}
	log.Println(sum)
}

func main() {
	filename := os.Args[1]
	inpt := readFile(filename)
	left, right := getInput(inpt)
	log.Println(left, right)
	calcPart2(left, right)
	log.Println("Here")
}
