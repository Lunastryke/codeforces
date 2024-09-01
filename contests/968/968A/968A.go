package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func CF968A(n int, input string) bool {
	return input[0] != input[n-1]
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
		input := strings.TrimSpace(line)
		if CF968A(n, input) {
			writer.WriteString("YES\n")
		} else {
			writer.WriteString("NO\n")
		}
	}

}
