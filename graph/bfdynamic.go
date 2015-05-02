package graph
import (
    "math"
)

func (graph *EdgeList) BellmanFordDP(start int) BFShortestPaths {
    V := graph.V()
    E := graph.E()
    infinity := math.Inf(1)
    dists := make([]float64, V + 1)
    edges := make([]Edge, E + 1)

    for i := 0; i < V; i++ {
        if i != start {
            dists[i] = infinity
        }
    }

    for i := 1; i < V; i++ {
        same := true
        for _, edge := range graph.Edges {
            if edge.V != start {
                dist := dists[edge.U] + edge.Weight
                if dist < dists[edge.V] {
                    dists[edge.V] = dist
                    edges[edge.V] = edge
                    same = false
                }
            }
        }
        if same { break }
    }

    var negCycle []Edge
    for _, edge := range graph.Edges {
        weight := edge.Weight
        if dists[edge.U] != infinity && dists[edge.U] + weight < dists[edge.V] {
            negCycle = append(negCycle, edge)
        }
    }

    return BFShortestPaths{Dists: dists, Edges: edges, NegativeCycle: negCycle}
}