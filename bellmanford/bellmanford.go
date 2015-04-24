package main

import (
    "github.com/sourcedelica/algorithms-go/graph"
    "os"
    "fmt"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s filename\n", os.Args[0])
        os.Exit(1)
    }
    filename := os.Args[1]
    ewdGraph := graph.ReadEWDGraph(filename)

    bf := ewdGraph.BellmanFord(0)

    if bf.NegativeCycle != nil {
        fmt.Printf("Negative cycle: %v\n", bf.NegativeCycle)
        os.Exit(1)
    }

    fmt.Printf("Result: %v\n", bf)
}