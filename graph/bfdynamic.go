package graph
import (
    "math"
)

// Bellman-Ford algorithm using dynamic programming
func (graph *EdgeList) BellmanFordDP(start int) BFShortestPaths {
    V := graph.V()
    E := graph.E()
    infinity := math.Inf(1)
    dists := make([]float64, V+1)
    edges := make([]Edge, E+1)
    vs := make(map[int][]Edge, E*2)

    // Create vs, map vertex -> vertex's edges
    for _, edge := range graph.Edges {
        var vEdges []Edge
        if ve, ok := vs[edge.V]; ok {
            vEdges = append(ve, edge)
        } else {
            vEdges = []Edge{edge}
        }
        vs[edge.V] = vEdges
    }

    // Set dist to +infinity except start, which is 0
    for i := 0; i <= V; i++ {
        if i != start {
            dists[i] = infinity
        }
    }

    // For |V|-1 iterations
    for i := 1; i < V; i++ {
        same := true

        // For each vertex
        for v, vEdges := range vs {

            // Find current shortest path through U to V
            min := dists[v]
            for _, edge := range vEdges {
                dist := dists[edge.U] + edge.Weight
                if dist < min {
                    min = dist
                    edges[v] = edge
                    same = false
                }
            }
            dists[v] = min
        }
        if same { break }  // If algorithm has converged, exit loop
    }

    // If any edges are returned they are part of a negative cycle
    var negCycle []Edge
    for _, edge := range graph.Edges {
        if dists[edge.U] + edge.Weight < dists[edge.V] {
            negCycle = append(negCycle, edge)
        }
    }

    return BFShortestPaths{Dists: dists, Edges: edges, NegativeCycle: negCycle}
}