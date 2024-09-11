package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CF971D(vertices [][]int) int {
	total := 0
	// Y coordinate can only be 0 or 1!!!
	mid := 0

	for i := 0; i < len(vertices); i++ {
		for j := 0; j < len(vertices); j++ {
			if i == j {
				continue
			}

			if vertices[i][1] == vertices[j][1] {
				// Horizontal line
				for k := 0; k < len(vertices); k++ {
					if vertices[k][0] == vertices[i][0] && vertices[k][1] != vertices[i][1] {
						total++
					}
					if (float64(vertices[k][0]) == float64(vertices[i][0]+vertices[j][0])/2.0) && vertices[k][1] != vertices[i][1] {
						mid++
					}
				}

			}
		}
	}
	return total + mid/2
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
		numVertices, _ := strconv.Atoi(line)
		// Convert the split strings to integers
		vertices := [][]int{}
		for j := 0; j < numVertices; j++ {
			line, _ = reader.ReadString('\n')
			line = strings.TrimSpace(line)
			parts := strings.Split(line, " ")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			vertices = append(vertices, []int{x, y})
		}
		// Output or process the integers as needed
		writer.WriteString(fmt.Sprintf("%d\n", CF971D(vertices)))
	}
}
