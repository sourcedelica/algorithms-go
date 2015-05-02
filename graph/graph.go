package graph

import (
    "fmt"
    "bufio"
    "strings"
    "github.com/sourcedelica/algorithms-go/util"
)

type AdjacencyList struct {
    numNodes int
    numEdges int
    Nodes map[int]Node
}

type EdgeList struct {
    numVertices int
    numEdges int
    Edges []Edge
}

func NewGraph(numNodes int, numEdges int) *AdjacencyList {
    nodes := make(map[int]Node, 2 * numNodes)
    return &AdjacencyList{numNodes: numNodes, numEdges: numEdges, Nodes: nodes}
}

func NewEdgeList(numVertices int, numEdges int) *EdgeList {
    return &EdgeList{numVertices, numEdges, make([]Edge, numEdges)}
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

func (a *AdjacencyList) V() int {
    return a.numNodes
}

func (a *AdjacencyList) E() int {
    return a.numEdges
}

func (a *EdgeList) V() int {
    return a.numVertices
}

func (a *EdgeList) E() int {
    return a.numEdges
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

// Read edge-weighted undirected graph into an AdjacencyList
func ReadEWUGraph(filename string) *AdjacencyList {
    return readEWGraph(filename, true)
}

// Read edge-weighted directed graph into an AdjacencyList
func ReadEWDGraph(filename string) *AdjacencyList {
    return readEWGraph(filename, false)
}

// Reads graphs in the format into an AdjacencyList
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

    ewGraph := NewGraph(numNodes, numEdges)

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

// Reads graphs in the format into an EdgeList
// #nodes #edges
// node1 node2 weight
// node1 node2 weight
// ...
func ReadEdgeList(filename string) *EdgeList {
    f := util.OpenFile(filename)
    defer f.Close()
    scanner := bufio.NewScanner(bufio.NewReader(f))

    sizes := strings.Split(util.ReadLine(scanner), " ")
    numNodes := util.Atoi(sizes[0])
    numEdges := util.Atoi(sizes[1])

    edgeList := NewEdgeList(numNodes, numEdges)

    for i := 0; i < numEdges; i++ {
        edgeParts := strings.Split(util.ReadLine(scanner), " ")
        from := util.Atoi(edgeParts[0])
        to := util.Atoi(edgeParts[1])
        weight := util.Atof(edgeParts[2])
        edgeList.Edges[i] = Edge{from, to, weight}
    }
    return edgeList
}