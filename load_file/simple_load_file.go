package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lineCounts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		lineCounts[input.Text()]++
	}

	for line, count := range lineCounts {
		fmt.Printf("%s:%d", line, count)
	}
}
