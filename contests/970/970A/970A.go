package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func CF970A(num1 int, num2 int) bool {
	if num1%2 == 0 {
		if num2%2 == 0 {
			return true
		} else if num1 >= 2 {
			return true
		}
	}
	return false
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

		// Output or process the integers as needed
		if CF970A(a, b) {
			writer.WriteString("YES\n")
		} else {
			writer.WriteString("NO\n")
		}
	}
}
