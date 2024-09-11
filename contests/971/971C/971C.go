package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CF971C(a int, b int, c int) int {
	if a == 0 && b == 0 {
		return 0
	}
	x := a / c
	y := b / c
	if a%c > 0 {
		x++
	}
	if b%c > 0 {
		y++
	}
	totalMoves := 0
	if x < y {
		totalMoves = x*2 + (y-x)*2
	} else if x == y {
		totalMoves = x * 2
	} else {
		totalMoves = y*2 + (x-y)*2 - 1
	}
	return totalMoves
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// Read the number of lines (test cases)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	numLines, _ := strconv.Atoi(line)

	// Process each line of input
	for i := 0; i < numLines; i++ {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")

		// Convert the split strings to integers
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		c, _ := strconv.Atoi(parts[2])

		// Output or process the integers as needed
		writer.WriteString(fmt.Sprintf("%d\n", CF971C(a, b, c)))
	}
}
