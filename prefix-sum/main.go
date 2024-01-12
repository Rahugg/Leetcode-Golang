package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	array := make([]int64, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		x, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		array[i] = x
	}

	prefixArray := findPrefix(array)

	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(prefixArray); i++ {
		fmt.Fprint(writer, prefixArray[i], " ")
	}
	writer.Flush()
}

func findPrefix(array []int64) []int64 {
	prefixArray := make([]int64, len(array)+1)

	for i := 0; i < len(array); i++ {
		prefixArray[i+1] += prefixArray[i] + array[i]
	}

	return prefixArray
}
