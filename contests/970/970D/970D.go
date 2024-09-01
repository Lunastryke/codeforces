package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CF970D(n int, p []int, color string) []int {
	// DFS? see which
	// make P a 1 indexed array
	p = append([]int{0}, p...)
	// make the color a 1 indexed array
	colorArr := []bool{}
	for _, c := range color {
		colorArr = append(colorArr, c == '0')
	}
	colorArr = append([]bool{false}, colorArr...)
	// Finding cycles, each cycle is a group of elements that are connected	to each other
	// result of every element in the cycle is the same
	visited := make([]bool, n+1)
	colorMap := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		curr := i
		currCycleVisited := []int{curr}
		blackCount := 0
		if colorArr[curr] {
			blackCount++
		}
		for !visited[p[curr]] {
			curr = p[curr]
			visited[curr] = true
			if colorArr[curr] {
				blackCount++
			}
			currCycleVisited = append(currCycleVisited, curr)
		}
		for _, v := range currCycleVisited {
			colorMap[v] = blackCount
		}
	}
	return colorMap[1:]
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
		// second line is the space-separated integers
		parts := strings.Split(line, " ")
		p := make([]int, numEle)
		for i, part := range parts {
			p[i], _ = strconv.Atoi(part)
		}
		// third line is the string of colors
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		color := line
		result := []string{}
		for _, r := range CF970D(numEle, p, color) {
			result = append(result, fmt.Sprintf("%d", r))
		}
		writer.WriteString(strings.Join(result, " ") + "\n")
	}
}
