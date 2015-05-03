package graph
import (
    "math"
    "sort"
)

func (graph *EdgeList) BellmanFordDP(start int) BFShortestPaths {
    V := graph.V()
    E := graph.E()
    infinity := math.Inf(1)
    dists := make([]float64, V + 1)
    edges := make([]Edge, E + 1)
    vs := make(map[int][]Edge, E*2)
    vertices := make(map[int]bool, V*2)
    var vmax int

    for _, edge := range graph.Edges {
        var vEdges []Edge
        if ve, ok := vs[edge.V]; ok {
            vEdges = append(ve, edge)
        } else {
            vEdges = []Edge{edge}
        }
        vs[edge.V] = vEdges
        vertices[edge.V] = true
        vertices[edge.U] = true

        if edge.V > vmax { vmax = edge.V }
    }

    for i := 0; i <= V; i++ {
        if i != start {
            dists[i] = infinity
        }
    }

    for i := 1; i <= V; i++ {
        even := i % 2 == 0
        vlist := sortedKeys(vertices, even)
        for _, v := range vlist {
            vEdges := vs[v]
            min := dists[v]
            for _, edge := range vEdges {
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

    var negCycle []Edge
    for _, edge := range graph.Edges {
        weight := edge.Weight
        if dists[edge.U] != infinity && dists[edge.U] + weight < dists[edge.V] {
            negCycle = append(negCycle, edge)
        }
    }

    return BFShortestPaths{Dists: dists, Edges: edges, NegativeCycle: negCycle}
}

// Go sucks so bad
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