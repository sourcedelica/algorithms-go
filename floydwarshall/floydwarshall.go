package main

import (
    "os"
    "fmt"
    "github.com/sourcedelica/algorithms-go/graph"
    "math"
)

type FWShortestPaths struct {
    NegativeCycle bool  // true if a negative cycle was detected
    Paths []graph.Edge  // Shortest s->t paths with their lengths
    edges [][]int       // edges[s][t] contains the "from" vertex of the edge to t
                        // along the shortest s->t path
}

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s filename\n", os.Args[0])
        os.Exit(1)
    }
    filename := os.Args[1]
    // Note - vertices must be numbered 1 or greater
    ewdGraph := graph.ReadEWDGraph(filename)

    sp := ewdGraph.FloydWarshall()

    if sp.NegativeCycle {
        fmt.Println("Negative cycle")
        os.Exit(1)
    }

    printShortestPath(&sp)
}

// Prints the shortest path found in the graph
func printShortestPath(sp *graph.FWShortestPaths) {
    min := math.Inf(1)
    var from, to int

    for _, edge := range sp.Paths {
        if edge.Weight < min {
            min = edge.Weight
            from = edge.From()
            to = edge.To()
        }
    }
    fmt.Printf("Minimum path %d->%d, length=%f\n", from, to, min)

    for _, i := range sp.Path(from, to) {
        fmt.Printf("%d", i)
        if i != to {
            fmt.Printf("->")
        }
    }
    fmt.Println()
}