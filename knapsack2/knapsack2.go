package main

import (
    "fmt"
    "github.com/sourcedelica/algorithms-go/knapsack"
)

// Solve Knapsack problem using top-down recursion
func main() {
    sack := knapsack.Load()

    result := sack.Find()

    fmt.Printf("Items: %v\n", result.Items)
    fmt.Printf("Weight: %d\n", result.Weight)
    fmt.Printf("Value: %d\n", result.Value)
}