package graph

// DFS as a recursive anonymous function
type recdfs func(recdfs, int)

// Find a negative cycle in the graph
// Returns nil if none were found
func (graph *AdjacencyList) FindNegativeCycle() []Edge {
    var cycle []Edge
    n := graph.V() + 2
    marked  := make([]bool, n)
    onStack := make([]bool, n)
    edges   := make([]Edge, n)

    dfs := func(f recdfs, v int) {
        onStack[v] = true
        marked[v] = true

        for _, edge := range graph.Nodes[v].Edges {
            w := edge.To()
            // Short circuit if directed cycle found
            if cycle != nil {
                return
            // DFS recurrence for each neighbor
            } else if !marked[w] {
                edges[w] = edge
                f(f, w)     // dfs(w)
            // Trace back directed cycle and push onto `cycle`
            } else if onStack[w] {
                for edge.From() != w {
                    cycle = append(cycle, edge)
                    edge = edges[edge.From()]
                }
                cycle = append(cycle, edge)
            }
        }

        onStack[v] = false
    }

    // Directed graph DFS loop
    for v, _ := range graph.Nodes {
        if !marked[v] {
            dfs(dfs, v)
        }
    }

    return cycle
}
