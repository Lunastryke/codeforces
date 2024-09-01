package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Fprintf(writer, "Read line: %s", line)
	}
}
