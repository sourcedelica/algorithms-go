package graph
import (
    "math"
)

func (graph *AdjacencyList) BellmanFordDP(start int) BFShortestPaths {
    n := graph.V()

    dists := make([][]float64, n + 1)
    edges := make([]Edge, n + 1)

    for i := 0; i < n; i++ {
        dists[i] = make([]float64, n + 1)
        if i == start {
            dists[0][i] = 0
        } else {
            dists[0][i] = math.Inf(1)
        }
    }

    for i := 1; i < n; i++ {
        same := true
        for v := 0; v < n; v++ {   // TODO: add v <= n and check for v exists (for graphs starting at 1)
            min := math.Inf(1)
            for w := 0; w < n; w++ {   // TODO: more efficient way to check for (w, v) than this loop
                if node, ok := graph.Nodes[w]; ok {
                    for _, edge := range node.Edges {
                        if edge.To() == v && v != start {
                            dist := edge.Weight + dists[i - 1][w]
                            if dist < min {
                                min = dist
                                edges[v] = edge
                            }
                        }
                    }
                }
            }
            dists[i][v] = math.Min(dists[i - 1][v], min)
            same = same && dists[i][v] == dists[i - 1][v]
        }
        if same { break }
    }

    var d []float64 = []float64{0}
    for v := 1; v <= n; v++ {
        d = append(d, dists[n - 1][v])
    }

    // TODO: is there a way to check for negative cycle in the DP algo?

    return BFShortestPaths{Dists: d, Edges: edges}
}