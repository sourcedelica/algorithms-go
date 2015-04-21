package main

import (
    "os"
    "fmt"
    "github.com/sourcedelica/algorithms-go/graph"
    "math"
)

type ShortestPaths struct {
    negativeCycle bool  // true if a negative cycle was detected
    paths []graph.Edge  // Shortest s->t paths with their lengths
    edges [][]int       // edges[s][t] contains the "from" vertex of the edge to t along the shortest s->t path
}

// Reconstructs the from->to shortest path from the edges matrix
func (sp *ShortestPaths) Path(from int, to int) []int {
    var path []int

    if sp.edges[from][to] == 0 {
        return nil
    }

    path = append(path, from)
    for from != to {
        from = sp.edges[from][to]
        path = append(path, from)
    }

    return path
}

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s filename\n", os.Args[0])
        os.Exit(1)
    }
    filename := os.Args[1]
    ewdGraph := graph.ReadEWDGraph(filename)

    sp := findShortestPaths(&ewdGraph)

    if sp.negativeCycle {
        fmt.Println("Negative cycle")
        os.Exit(1)
    }

    printShortestPath(&sp)
}

// Prints the shortest path found in the graph
func printShortestPath(sp *ShortestPaths) {
    min := math.Inf(1)
    var from, to int
    for _, edge := range(sp.paths) {
        if edge.Weight < min {
            min = edge.Weight
            from = edge.From()
            to = edge.To()
        }
    }
    fmt.Printf("Minimum path %d->%d, length=%f\n", from, to, min)

    for _, i := range(sp.Path(from, to)) {
        fmt.Printf("%d", i)
        if i != to {
            fmt.Printf("->")
        }
    }
    fmt.Println()
}

// All-pairs shortest paths using Floyd-Warshall algorithm
// http://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm
func findShortestPaths(ewdGraph *graph.AdjacencyList) ShortestPaths {
    n := ewdGraph.Size()

    dists := make([][]float64, n + 1)
    edges := make([][]int, n + 1)

    // Initialization
    for i := 1; i <= n; i++ {
        dists[i] = make([]float64, n + 1)
        edges[i] = make([]int, n + 1)
        for j := 1; j <= n; j++ {
            dists[i][j] = math.Inf(1)
        }
    }

    for i := 1; i <= n; i++ {
        for _, edge := range(ewdGraph.Nodes[i].Edges) {
            dists[i][edge.To()] = edge.Weight
            edges[i][edge.To()] = edge.To()
        }
    }

    for k := 1; k <= n; k++ {
        for i := 1; i <= n; i++ {
            for j := 1; j <= n; j++ {
                left := dists[i][j]
                right := dists[i][k] + dists[k][j]
                if (left > right) {
                    dists[i][j] = right
                    edges[i][j] = edges[i][k]
                }
            }
            if dists[i][i] < 0 {
                return ShortestPaths{negativeCycle: true}
            }
        }
    }

    var paths []graph.Edge

    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            if i != j && !math.IsInf(dists[i][j], 1) {
                paths = append(paths, graph.Edge{i, j, dists[i][j]})
            }
        }
    }

    return ShortestPaths{paths: paths, edges: edges}
}