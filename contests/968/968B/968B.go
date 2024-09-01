package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// func CF968B(n int, input []int) int {

// 	// Turtle goes first and tries to maximise the value of last number
// 	// Piggie goes next and tries to minimise the value of last number

// 	// So Turtle will always pick the number before the smallest number
// 	// Piggie will always pick the largest number

// 	// First number in the input can never be removed, it can only be replaced with the number after it
// 	turtleTurn := true
// 	for len(input) > 1 {
// 		if turtleTurn {
// 			// Get index of smallest number
// 			minIndex := 0
// 			for i := 1; i < len(input); i++ {
// 				if input[i] < input[minIndex] {
// 					minIndex = i
// 				}
// 			}
// 			// Remove the number before the smallest number
// 			input = append(input[:minIndex], input[minIndex+1:]...)
// 		} else {
// 			// Get index of largest number
// 			maxIndex := 0
// 			for i := 1; i < len(input); i++ {
// 				if input[i] > input[maxIndex] {
// 					maxIndex = i
// 				}
// 			}
// 			// Remove the largest number
// 			input = append(input[:maxIndex], input[maxIndex+1:]...)
// 		}
// 		turtleTurn = !turtleTurn
// 	}
// 	return input[0]
// }

func solveGame(a []int) int {
	m := len(a)

	for m > 1 {
		if len(a)%2 == 1 { // Turtle's turn
			a[0] = max(a[0], a[1])
		} else { // Piggy's turn
			a[0] = min(a[0], a[1])
		}
		a = a[:1+copy(a[1:], a[2:])]
		m--
	}
	return a[0]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	// First line is the number of test cases
	n, _ := strconv.Atoi(line)
	for i := 0; i < n; i++ {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		// Second line is the input
		n, _ := strconv.Atoi(line)
		line, _ = reader.ReadString('\n')
		// Split line into slice of ints
		input := strings.Fields(line)
		intInput := make([]int, n)
		for i := 0; i < n; i++ {
			intInput[i], _ = strconv.Atoi(input[i])
		}

		writer.WriteString(strconv.Itoa(solveGame(intInput)) + "\n")
	}

}
