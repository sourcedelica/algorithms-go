package graph
import (
    "math"
    "sort"
)

// Funky variant of Bellman-Ford
func (graph *EdgeList) BellmanFordAlternating(start int) BFShortestPaths {
    V := graph.V()
    E := graph.E()
    infinity := math.Inf(1)
    dists := make([]float64, V+1)
    edges := make([]Edge, E+1)
    vs := make(map[int][]Edge, E*2)
    vertices := make(map[int]bool, V*2)

    for _, edge := range graph.Edges {
        var vEdges []Edge
        if ve, ok := vs[edge.V]; ok {
            vEdges = append(ve, edge)
        } else {
            vEdges = []Edge{edge}
        }
        vs[edge.V] = vEdges
        vertices[edge.U] = true
        vertices[edge.V] = true
    }

    // Initialize dists to +infinity
    for i := 0; i <= V; i++ {
        if i != start {
            dists[i] = infinity
        }
    }

    // For |V|-1 iterations
    for i := 1; i < V; i++ {
        even := i % 2 == 0

        // On odd iterations of the outer loop, iterate over vertices from 1..n
        // On even iterations of the outher loop, iterate over the vertices from n..1
        vlist := sortedKeys(vertices, even)
        for _, v := range vlist {

            // Find current shortest path through U to V
            // considering only vertices where U > V on even iterations
            // or only vertices where V > U on odd iterations
            min := dists[v]
            for _, edge := range vs[v] {
                if even && edge.U > edge.V || !even && edge.U < edge.V {
                    dist := dists[edge.U] + edge.Weight
                    if dist < min {
                        min = dist
                        edges[v] = edge
                    }
                }
            }
            dists[v] = min
        }
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

// Return a sorted slice of keys from vmap
// (Go needs generics and functional programming)
func sortedKeys(vmap map[int]bool, reverse bool) []int {
    keys := make([]int, len(vmap))
    i := 0
    for k := range vmap {
        keys[i] = k
        i += 1
    }

    if reverse {
        sort.Sort(sort.Reverse(sort.IntSlice(keys)))
    } else {
        sort.Ints(keys)
    }
    return keys
}