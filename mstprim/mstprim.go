package main
import (
    "os"
    "fmt"
    "container/heap"
    "github.com/sourcedelica/algorithms-go/graph"
)

// A min-heap of Edges by weight
type EdgeHeap []graph.Edge

func (h EdgeHeap) Len() int           { return len(h) }
func (h EdgeHeap) Less(i, j int) bool { return h[i].Weight < h[j].Weight }
func (h EdgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EdgeHeap) Push(x interface{}) {
	*h = append(*h, x.(graph.Edge))
}

func (h *EdgeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s filename\n", os.Args[0])
        os.Exit(1)
    }
    filename := os.Args[1]
    ewGraph := graph.ReadEWUGraph(filename)

    mst := MST(ewGraph)

    fmt.Printf("%f\n", Sum(mst))
}

// Compute minimum spanning tree using Prim's algorithm
// http://en.wikipedia.org/wiki/Prim%27s_algorithm
func MST(ewGraph graph.AdjacencyList) []graph.Edge {
    edgeHeap := &EdgeHeap{}
    visited := make(map[int]bool, 2 * len(ewGraph.Nodes))
    mst := make([]graph.Edge, 0)

    Visit(ewGraph, edgeHeap, visited, ewGraph.First().Id)

    for edgeHeap.Len() > 0 {
        edge := heap.Pop(edgeHeap).(graph.Edge)
        u, v := edge.U, edge.V

        if visited[u] && visited[v] {
            continue
        }
        mst = append(mst, edge)

        if !visited[u] { Visit(ewGraph, edgeHeap, visited, u) }
        if !visited[v] { Visit(ewGraph, edgeHeap, visited, v) }
    }

    return mst
}

func Visit(graph graph.AdjacencyList, h *EdgeHeap, visited map[int]bool, id int) {
    visited[id] = true
    for _, edge := range graph.Nodes[id].Edges {
        if !visited[edge.Other(id)] {
            heap.Push(h, edge)
        }
    }
}

func Sum(mst []graph.Edge) float64 {
    sum := float64(0)
    for _, edge := range mst {
        sum += edge.Weight
    }
    return sum
}