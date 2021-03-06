package graph

import (
    "math"
)

type FWShortestPaths struct {
    NegativeCycle bool  // true if a negative cycle was detected
    Paths []Edge        // Shortest s->t paths with their lengths
    edges [][]int       // edges[s][t] contains the "from" vertex of the edge to t
    // along the shortest s->t path
}

// All-pairs shortest paths using Floyd-Warshall algorithm
// http://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm
func (graph *AdjacencyList) FloydWarshall() FWShortestPaths {
    n := graph.V()

    dists := make([][]float64, n + 1)
    edges := make([][]int, n + 1)

    // Initialization
    for i := 1; i <= n; i++ {
        dists[i] = make([]float64, n + 1)
        edges[i] = make([]int, n + 1)
        for j := 1; j <= n; j++ {
            dists[i][j] = math.Inf(1)
        }
        for _, edge := range graph.Nodes[i].Edges {
            dists[i][edge.To()] = edge.Weight
            edges[i][edge.To()] = edge.To()
        }
    }

    // Compute shortest paths using 1..k as intermediate vertices
    for k := 1; k <= n; k++ {
        for i := 1; i <= n; i++ {
            for j := 1; j <= n; j++ {
                left := dists[i][j]
                right := dists[i][k] + dists[k][j]
                if left > right {
                    dists[i][j] = right
                    edges[i][j] = edges[i][k]
                }
            }
            if dists[i][i] < 0 {
                return FWShortestPaths{NegativeCycle: true}
            }
        }
    }

    var paths []Edge

    // Collect the shortest i->j paths and their length
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            if i != j && !math.IsInf(dists[i][j], 1) {
                paths = append(paths, Edge{i, j, dists[i][j]})
            }
        }
    }

    return FWShortestPaths{Paths: paths, edges: edges}
}

// Reconstructs the from->to shortest path from the edges matrix
func (sp *FWShortestPaths) Path(from int, to int) []int {
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