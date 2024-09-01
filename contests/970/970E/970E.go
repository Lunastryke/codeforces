package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func CF970E(n int, s string) int {
	if n == 2 {
		return 0
	}
	operations := 0
	if n%2 != 0 {
		// We need to choose 1 character to remove
		// Find the character that has the lowest count and remove it
		charCount := make([]int, 26)
		for i := 0; i < n; i++ {
			charCount[s[i]-'a']++
		}
		for i := 0; i < 26; i++ {
			if charCount[i] == 0 {
				charCount[i] = math.MaxInt32
			}
		}
		// Find the character with the lowest odd count
		lowestOddCount := math.MaxInt32
		lowestEvenCount := math.MaxInt32
		for i := 0; i < 26; i++ {
			if charCount[i]%2 != 0 && charCount[i] < lowestOddCount {
				lowestOddCount = charCount[i]
			}
			if charCount[i]%2 == 0 && charCount[i] < lowestEvenCount {
				lowestEvenCount = charCount[i]
			}
		}
		if lowestOddCount != math.MaxInt32 {
			// remove the character
			for i := 0; i < n; i++ {
				if charCount[int(s[i]-'a')] == lowestOddCount {
					s = s[:i] + s[i+1:]
					break
				}
			}
		} else {
			for i := 0; i < n; i++ {
				if charCount[int(s[i]-'a')] == lowestEvenCount {
					s = s[:i] + s[i+1:]
					break
				}
			}
		}

		operations++
		n--
	}
	if len(s) == 0 || len(s) == 2 {
		return operations
	}
	// we try doing the even case first
	// if n is even, we need to partition the array into odd and even characters and make counts for each of them
	evenCharCount := make([]int, 26)
	oddCharCount := make([]int, 26)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			evenCharCount[s[i]-'a']++
		} else {
			oddCharCount[s[i]-'a']++
		}
	}
	// we need to find the character that has the most count in the even and odd characters
	maxEvenChar := 0
	maxOddChar := 0
	for i := 0; i < 26; i++ {
		if evenCharCount[i] > evenCharCount[maxEvenChar] {
			maxEvenChar = i
		}
		if oddCharCount[i] > oddCharCount[maxOddChar] {
			maxOddChar = i
		}
	}
	// Operations needed
	operations += (n/2 - evenCharCount[maxEvenChar]) + (n/2 - oddCharCount[maxOddChar])

	return operations
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
		n, _ := strconv.Atoi(line)
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		s := line
		writer.WriteString(strconv.Itoa(CF970E(n, s)) + "\n")
	}
}
