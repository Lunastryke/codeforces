package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CF71A(words []string) []string {
	abbv := []string{}
	for i := 0; i < len(words); i++ {
		n := len(words[i])
		if n > 10 {
			// We need to abbreviate
			abbv = append(abbv, string(words[i][0])+fmt.Sprint(n-2)+string(words[i][n-1]))
		} else {
			abbv = append(abbv, words[i])
		}
	}
	return abbv
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	// First line is the number of words
	n, _ := strconv.Atoi(line)
	words := make([]string, n)
	for i := 0; i < n; i++ {
		line, _ = reader.ReadString('\n')
		words[i] = strings.TrimSpace(line)
	}
	abbv := CF71A(words)
	for _, a := range abbv {
		writer.WriteString(a + "\n")
	}
}
