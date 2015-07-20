package main
import (
	"os"
	"fmt"
	"github.com/sourcedelica/algorithms-go/util"
)

// Rod-cutting problem
// Given a rod of length n inches and prices of all integer sizes from 1 to n,
// determine the maximum value obtainable by cutting up the rod and selling the pieces.

// For example, given a rod of length 8, the prices of each integral size are given as follows:
// length 1   2   3   4   5   6   7   8
// price  1   5   8   9  10  17  17  20
// The maximum value for this rod is 22 (dividing the rod into lengths 2 and 6)

// TODO - add reconstruction for recursive

func main() {
	argc := len(os.Args)

	if (argc < 2 || (os.Args[1] != "recursive" && os.Args[1] != "dp")) {
		fmt.Printf("Usage: %s recursive|dp price...\n", os.Args[0])
		os.Exit(1)
	}

	prices := make([]int, argc - 2)
	for i := 2; i < argc; i++ {
		prices[i - 2] = util.Atoi(os.Args[i])
	}

	var maxValue int
	if os.Args[1] == "recursive" {
		maxValue = CutRecursive(prices)
	} else {
		maxValue = CutDP(prices)
	}
	fmt.Printf("Max value: %d\n", maxValue)
}

func CutRecursive(prices []int) int {
	maxValues := make([]int, len(prices))
	return cutRecursive(prices, maxValues, len(prices))
}

func cutRecursive(prices []int, maxValues []int, n int) int {
	if n < 1 {
		return 0
	} else if n == 1 {
		return prices[0]
	} else if maxValues[n - 1] > 0 {
		return maxValues[n - 1]
	}
	maxValue := -1

	for i := 0; i < n; i++ {
		value := prices[i] + cutRecursive(prices, maxValues, n - i - 1)
		if value > maxValue {
			maxValue = value
		}
	}

	maxValues[n - 1] = maxValue
	return maxValue
}

func CutDP(prices []int) int {
	n := len(prices)
 	maxValues := make([]int, n + 1)
	pieces := make([]int, n + 1)

	for i := 1; i <= n; i++ {
		maxValue := -1

		for j := 1; j <= i; j++ {
			piece := i - j
			if (prices[j - 1] + maxValues[piece] > maxValue) {
				maxValue = prices[j - 1] + maxValues[piece]
				pieces[i] = j
			}
		}

		maxValues[i] = maxValue
	}

	/*lengths := */reconstruct(pieces, n)  // TODO - after recursive reconstruction return lengths to main
	return maxValues[n]
}

func reconstruct(pieces []int, n int) []int {
	var lengths []int

	for ; n > 0; {
		lengths = append(lengths, pieces[n])
		n -= pieces[n]
	}
fmt.Printf("%v\n", lengths)   // TODO - print price too
	return lengths
}