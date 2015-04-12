package main

import (
    "github.com/sourcedelica/algorithms-go/knapsack"
    "fmt"
    "github.com/sourcedelica/algorithms-go/util"
)

// Solve Knapsack problem using recursion and memoization
func main() {
    sack := knapsack.Load()
    n := sack.Size()
    W := sack.W()
    memo := make(map[string]int)

    optimalValue := find(n, W, &sack, memo)

    fmt.Printf("Value of sack: %d\n", optimalValue)
}

func find(i int, x int, sack *knapsack.Knapsack, memo map[string]int) int {
    if i == 0 {
        return 0
    }

    k := key(i, x)
    if aix, ok := memo[k]; ok {
        return aix
    }

    w := sack.Weight(i)
    v := sack.Value(i)

    var aix int
    if w > x {
        aix = find(i - 1, x, sack, memo)
    } else {
        aix = util.Max(find(i - 1, x, sack, memo), find(i - 1, x - w, sack, memo) + v)
    }
    memo[k] = aix
    return aix
}

func key(i int, x int) string {
    return fmt.Sprintf("%d,%d", i, x)
}