package main

import (
    "os"
    "fmt"
    "github.com/sourcedelica/algorithms-go/graph"
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

    sp.PrintShortestPath()
}