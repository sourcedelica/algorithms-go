package main

import (
    "fmt"
    "github.com/sourcedelica/algorithms-go/knapsack"
    "github.com/sourcedelica/algorithms-go/util"
    "os"
)

// Solve for a sack of weight W1, then for the remaining items, solve for weight W2
// Uses top-down recursion
func main() {
    sack := knapsack.Load()
    w1 := util.Atoi(os.Args[2])
    sack.SetW(w1)

    fmt.Printf("Sack size: %d, max weight: %d\n", sack.Size(), sack.W())

    result := sack.Find()

    fmt.Printf("Items: %v\n", result.Items)
    fmt.Printf("Weight: %d\n", result.Weight)
    fmt.Printf("Value: %d\n", result.Value)

    w2 := util.Atoi(os.Args[3])
    var items2 []knapsack.Item
    resultMap := make(map[int]bool)
    for _, i := range result.Items {
        resultMap[i] = true
    }
    for i := 1; i <= sack.Size(); i++ {
        if (!resultMap[i]) {
            items2 = append(items2, sack.Item(i))
        }
    }
    sack2 := knapsack.New(w2, items2)

    result = sack2.Find()

    fmt.Printf("Sack size: %d, max weight: %d\n", sack2.Size(), sack2.W())
    fmt.Printf("Items: %v\n", result.Items)
    fmt.Printf("Weight: %d\n", result.Weight)
    fmt.Printf("Value: %d\n", result.Value)
}