package main

import (
    "github.com/sourcedelica/algorithms-go/graph"
    "os"
    "fmt"
)

func main() {
    if len(os.Args) < 1 {
        fmt.Fprintf(os.Stderr, "Usage: %s filename\n", os.Args[0])
        os.Exit(1)
    }
    filename := os.Args[1]

    // Note: vertices must start at 1
    ewdGraph := graph.ReadEWDGraph(filename)

    for v, _ := range ewdGraph.Nodes {
        ewdGraph.AddEdge(0, v, 0)
    }

    bf := ewdGraph.BellmanFord(0)
    delete(ewdGraph.Nodes, 0)

    if bf.NegativeCycle != nil {
        fmt.Printf("Negative cycle: %v\n", bf.NegativeCycle)
        os.Exit(1)
    }

    for _, node := range ewdGraph.Nodes {
        for _, edge := range node.Edges {
            edge.Weight = edge.Weight + bf.Dists[edge.U] - bf.Dists[edge.V]
fmt.Printf("edge: %v\n", edge)
        }
    }
}