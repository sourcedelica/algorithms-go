package graph

// Compute connected components of an undirected graph
// Returns two-dimensional integer slice, indexed by component# first,
// then by members of that component
// TODO: adapt this for SCC (Kosaraju)
func (graph *AdjacencyList) CC() (map[int]int, int) {
    n := graph.V() + 1
    count := 0
    marked  := make(map[int]bool, n)
    compMap := make(map[int]int)

    // DFS function
    dfs := func(f recdfs, v int) {
        marked[v] = true
        compMap[v] = count

        for _, edge := range graph.Nodes[v].Edges {
            w := edge.To()
            if !marked[w] {
                f(f, w)     // dfs(w)
            }
        }
    }

    // DFS on each vertex that hasn't been visited
    for v, _ := range graph.Nodes {
        if !marked[v] {
            dfs(dfs, v)
            count++
        }
    }


    return compMap, count
}
