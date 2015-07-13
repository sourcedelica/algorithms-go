package main
import (
	"os"
	"fmt"
	"github.com/sourcedelica/algorithms-go/util"
)

func main() {
	argc := len(os.Args)

	if (argc < 2 || (os.Args[1] != "recursive" && os.Args[1] != "dp")) {
		fmt.Printf("Usage: %s %v recursive|dp price...\n", os.Args[0])
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

	for i := 1; i <= n; i++ {
		maxValue := -1

		for j := 0; j < i; j++ {
			maxValue = util.Max(maxValue, prices[j] + maxValues[i - j - 1])
		}

		maxValues[i] = maxValue
	}

	return maxValues[n]
}