package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func CF968C(n int, input string) string {
	return ""
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
		output := CF968C(n, input)
		writer.WriteString(output + "\n")
	}

}
