package main

import (
    "github.com/sourcedelica/algorithms-go/graph"
    "os"
    "fmt"
    "github.com/sourcedelica/algorithms-go/util"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s filename start\n", os.Args[0])
        os.Exit(1)
    }
    filename := os.Args[1]
    edgeList := graph.ReadEdgeList(filename)
    start := util.Atoi(os.Args[2])

    bf := edgeList.BellmanFordDP(start)

    if bf.NegativeCycle != nil {
        fmt.Printf("Negative cycle includes: %v\n", bf.NegativeCycle)
        os.Exit(1)
    }

    fmt.Printf("Distances: %v\n", bf.Dists)
    fmt.Printf("Edges: %v\n", bf.Edges)
}