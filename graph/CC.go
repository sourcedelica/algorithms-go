package graph

// Compute connected components of an undirected graph
// Returns two-dimensional integer slice, indexed by component# first, then by members of that component
func (graph *AdjacencyList) CC() [][]int {
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

    // Convert component map to slices
    components := make([][]int, count)
    for k, v := range compMap {
        components[v] = append(components[v], k)
    }
    return components
}
