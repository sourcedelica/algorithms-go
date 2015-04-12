package main

import (
    "github.com/sourcedelica/algorithms-go/knapsack"
    "fmt"
    "github.com/sourcedelica/algorithms-go/util"
)

// Solve Knapsack problem using iteration
func main() {
    sack := knapsack.Load()

    n := sack.Size()
    W := sack.W()
    A := make([][]int, n + 1)
    A[0] = make([]int, W + 1)

    for i := 1; i <= n; i++ {
        A[i] = make([]int, W + 1)
        for x := 0; x <= W; x++ {
            var aix int
            w := sack.Weight(i)
            if w > x {
                aix = A[i - 1][x]
            } else {
                v := sack.Value(i)
                aix = util.Max(A[i - 1][x], A[i - 1][x - w] + v)
            }
            A[i][x] = aix
        }
    }

    fmt.Printf("Value of sack: %d\n", A[n][W])
}