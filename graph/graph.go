package graph

import (
    "fmt"
    "bufio"
    "strings"
    "github.com/sourcedelica/algorithms-go/util"
)

type AdjacencyList struct {
    Nodes map[int]Node
}

func NewGraph(numNodes int) *AdjacencyList {
    nodes := make(map[int]Node, 2 * numNodes)
    return &AdjacencyList{ Nodes: nodes }
}

func (a *AdjacencyList) First() Node {
    var first Node
    for _, v := range a.Nodes { first = v; break }
    return first
}

func (a *AdjacencyList) AddEdge(from int, to int, weight float64) {
    node, ok := a.Nodes[from]
    if (!ok) {
        node = Node{Id: from}
    }
    node.Edges = append(node.Edges, Edge{U: from, V: to, Weight: weight})
    a.Nodes[from] = node

    node, ok = a.Nodes[to]
    if (!ok) {
        node = Node{Id: to}
        a.Nodes[to] = node
    }
}

func (a *AdjacencyList) Size() int {
    return len(a.Nodes)
}

type Node struct {
    Id int
    Edges []Edge
}

type Edge struct {
    U int
    V int
    Weight float64
}

func (n Node) EdgeTo(i int) (*Edge, bool) {
    for _, edge := range(n.Edges) {
        if (edge.V == i) {
            return &edge, true
        }
    }
    return nil, false
}

func (n Node) HasEdge(i int) bool {
    _, ok := n.EdgeTo(i)
    return ok
}

func (e Edge) Other(id int) int {
    if id == e.U {
        return e.V
    } else if id == e.V {
        return e.U
    } else {
        panic(fmt.Sprintf("No other for %d in Edge %v", id, e))
    }
}

func (e Edge) From() int {
    return e.U
}

func (e Edge) To() int {
    return e.V
}

// Read edge-weighted undirected graph
func ReadEWUGraph(filename string) *AdjacencyList {
    return readEWGraph(filename, true)
}

// Read edge-weighted directed graph
func ReadEWDGraph(filename string) *AdjacencyList {
    return readEWGraph(filename, false)
}

// Reads graphs in the format
// #nodes #edges
// node1 node2 weight
// node1 node2 weight
// ...
func readEWGraph(filename string, undirected bool) *AdjacencyList {
    f := util.OpenFile(filename)
   	defer f.Close()
   	scanner := bufio.NewScanner(bufio.NewReader(f))

    sizes := strings.Split(util.ReadLine(scanner), " ")
    numNodes := util.Atoi(sizes[0])
    numEdges := util.Atoi(sizes[1])

    ewGraph := NewGraph(numNodes)

    for i := 0; i < numEdges; i++ {
        edgeParts := strings.Split(util.ReadLine(scanner), " ")
        from := util.Atoi(edgeParts[0])
        to := util.Atoi(edgeParts[1])
        weight := util.Atof(edgeParts[2])
        ewGraph.AddEdge(from, to, weight)
        if (undirected) {
            ewGraph.AddEdge(to, from, weight)
        }
    }
    return ewGraph
}