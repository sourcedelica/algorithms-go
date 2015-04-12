package main

import (
    "github.com/sourcedelica/algorithms-go/knapsack"
    "fmt"
)

// Solve Knapsack problem using recursion and memoization
func main() {
    sack := knapsack.Load()

    n := sack.Size()
    W := sack.W()
    memo := make(map[string]int)

    optValue := solve(n, W, &sack, memo)

    fmt.Printf("Value of sack: %d\n", optValue)
}

func solve(i int, x int, sack *knapsack.Knapsack, memo map[string]int) int {
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
        aix = solve(i - 1, x, sack, memo)
    } else {
        aix = Max(solve(i - 1, x, sack, memo), solve(i - 1, x - w, sack, memo) + v)
    }
    memo[k] = aix
    return aix
}

func key(i int, x int) string {
    return fmt.Sprintf("%d,%d", i, x)
}

func Max(a int, b int) int {
    if (a > b) {
        return a
    } else {
        return b
    }
}
