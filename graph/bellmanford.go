package graph

import (
    "github.com/sourcedelica/algorithms-go/util"
    "math"
)

type bellmanFord struct {
    graph         *AdjacencyList
    queue         *util.IntQueue
    onQueue       []bool
    dists         []float64
    edges         []Edge
    pass          int
    negativeCycle *[]Edge
}

type BFShortestPaths struct {
    NegativeCycle []Edge
    Dists         []float64
    Edges         []Edge
}

// Single-source shortest path using the Bellman-Ford algorithm
// http://en.wikipedia.org/wiki/Bellman%E2%80%93Ford_algorithm
func (ewdGraph *AdjacencyList) BellmanFord(start int) BFShortestPaths {
    n := ewdGraph.V() + 1    // Allowing vertices to start from 1
    dists := make([]float64, n)
    edges := make([]Edge, n)
    onQueue := make([]bool, n)
    var queue util.IntQueue
    var negativeCycle []Edge
    var pass int

    for k, _ := range ewdGraph.Nodes {
        dists[k] = math.Inf(1)
    }
    dists[start] = 0
    queue.Enqueue(start)
    onQueue[start] = true

    bf := bellmanFord{ewdGraph, &queue, onQueue, dists, edges, pass, &negativeCycle}

    // Relax all edges until no more edges have been relaxed
    for !queue.Empty() && len(*bf.negativeCycle) == 0 {
        v := queue.Dequeue()
        onQueue[v] = false
        bf.relax(v)
    }

    return BFShortestPaths{Dists: dists, Edges: edges, NegativeCycle: negativeCycle}
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
        bf.pass += 1
        if bf.pass % bf.graph.V() == 0 {
            cycle := bf.findNegativeCycle()
            if (cycle != nil) {
                *bf.negativeCycle = cycle
                break
            }
        }
    }
}

func (bf *bellmanFord) findNegativeCycle() []Edge {
    n := len(bf.edges)
    g := NewGraph(n, n)

    for v := 0; v < n; v++ {
        if (bf.edges[v] != Edge{}) {
            g.AddEdge(bf.edges[v].U, bf.edges[v].V, bf.edges[v].Weight)
        }
    }

    return g.FindNegativeCycle()
}