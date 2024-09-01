package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func CF970B(n int, matrix string) bool {
	// if n cannot be square rooted, return false

	if !hasIntegerSquareRoot(n) {
		return false
	}
	// Check if the matrix is beautiful
	// Calculate the r and c
	r := 0
	for i := 0; i < n; i++ {
		if matrix[i] != '1' {
			break
		}
		r++
	}
	if r == n {
		return n <= 4
	}
	r--
	for i := n - 1; i >= n-1-r; i-- {
		if matrix[i] != '1' {
			return false
		}
	}
	// Both top and bottom rows must have the same number of 1s
	// Check the middle rows for 100001 pattern
	for i := r; i < n-1-r; i = i + r {
		row := matrix[i:(i + r)]
		if !checkRowSurrounded10(row) {
			return false
		}
	}

	return true
}

func checkRowSurrounded10(row string) bool {
	if row[0] != '1' || row[len(row)-1] != '1' {
		return false
	}
	for i := 1; i < len(row)-1; i++ {
		if row[i] != '0' {
			return false
		}
	}
	return true
}

// hasIntegerSquareRoot checks if a number has an integer square root
func hasIntegerSquareRoot(n int) bool {
	if n < 0 {
		return false // Negative numbers do not have real square roots
	}
	// Compute the integer square root
	sqrt := int(math.Sqrt(float64(n)))
	// Check if squaring the result gives the original number
	return sqrt*sqrt == n
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
		// first line is number of elements in the matrix
		numEle, _ := strconv.Atoi(line)
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		// second line is the matrix
		matrix := line
		if CF970B(numEle, matrix) {
			writer.WriteString("YES\n")
		} else {
			writer.WriteString("NO\n")
		}
	}
}
