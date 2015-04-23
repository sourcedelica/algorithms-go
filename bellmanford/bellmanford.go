package main

import (
    "github.com/sourcedelica/algorithms-go/graph"
    "math"
    "github.com/sourcedelica/algorithms-go/util"
    "os"
    "fmt"
)

type bellmanFord struct {
    graph         *graph.AdjacencyList
    queue         *util.IntQueue
    onQueue       []bool
    dists         []float64
    edges         []graph.Edge
    pass          *int
    negativeCycle *[]graph.Edge
}

type BellmanFord struct {
    NegativeCycle []graph.Edge
    Dists         []float64
    Edges         []graph.Edge
}

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s filename\n", os.Args[0])
        os.Exit(1)
    }
    filename := os.Args[1]
    ewdGraph := graph.ReadEWDGraph(filename)

    bf := BellmanFordSP(ewdGraph, 0)

    if bf.NegativeCycle != nil {
        fmt.Printf("Negative cycle: %v\n", bf.NegativeCycle)
        os.Exit(1)
    }

    fmt.Printf("Result: %v\n", bf)
}

func BellmanFordSP(ewdGraph *graph.AdjacencyList, start int) BellmanFord {
    n := ewdGraph.Size() + 1    // Allowing vertices to start from 1
    dists := make([]float64, n)
    edges := make([]graph.Edge, n)
    onQueue := make([]bool, n)
    var queue util.IntQueue
    var negativeCycle []graph.Edge
    var pass int

    for k, _ := range ewdGraph.Nodes {
        dists[k] = math.Inf(1)
    }
    dists[start] = 0
    queue.Enqueue(start)
    onQueue[start] = true

    bf := bellmanFord{ewdGraph, &queue, onQueue, dists, edges, &pass, &negativeCycle}

    // Relax all edges until no more edges have been relaxed
    for !queue.Empty() && len(*bf.negativeCycle) == 0 {
        v := queue.Dequeue()
        onQueue[v] = false
        bf.relax(v)
    }

    return BellmanFord{Dists: dists, Edges: edges, NegativeCycle: negativeCycle}
}

func (bf *bellmanFord) relax(v int) {
    // Relax all edges out of v
    for _, edge := range bf.graph.Nodes[v].Edges {
        to := edge.To()
        if bf.dists[to] > bf.dists[v] + edge.Weight {
            bf.dists[to] = bf.dists[v] + edge.Weight
            bf.edges[to] = edge
            // If an edge was relaxed, queue its `to` vertex for relaxing next iteration
            if !bf.onQueue[to] {
                bf.queue.Enqueue(to)
                bf.onQueue[to] = true
            }
        }

        // Check for negative cycle
        *bf.pass += 1
        if *bf.pass % bf.graph.Size() == 0 {
            cycle := bf.findNegativeCycle()
            if (cycle != nil) {
                fmt.Printf("found cycle %v\n", cycle)
                *bf.negativeCycle = cycle
                break
            }
        }
    }
}

func (bf *bellmanFord) findNegativeCycle() []graph.Edge {
    n := len(bf.edges)
    g := graph.NewGraph(n)

    for v := 0; v < n; v++ {
        if (bf.edges[v] != graph.Edge{}) {
            g.AddEdge(bf.edges[v].U, bf.edges[v].V, bf.edges[v].Weight)
        }
    }

    return g.FindNegativeCycle()
}