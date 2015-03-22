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

type Node struct {
    Id int
    Edges []Edge
}

type Edge struct {
    U int
    V int
    Weight float64
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

// Reads graphs in the format
// #nodes #edges
// node1 node2 weight
// node1 node2 weight
// ...
func ReadEWGraph(filename string) AdjacencyList {
    f := util.OpenFile(filename)
   	defer f.Close()

   	scanner := bufio.NewScanner(bufio.NewReader(f))

    sizes := strings.Split(util.ReadLine(scanner), " ")
    numNodes := util.Atoi(sizes[0])
    numEdges := util.Atoi(sizes[1])

    nodes := make(map[int]Node, 2 * numNodes)

    for i := 0; i < numEdges; i++ {
        edgeParts := strings.Split(util.ReadLine(scanner), " ")
        node1 := util.Atoi(edgeParts[0])
        node2 := util.Atoi(edgeParts[1])
        weight := util.Atof(edgeParts[2])
        AddEdge(nodes, node1, node2, weight)
        AddEdge(nodes, node2, node1, weight)
    }

    return AdjacencyList{ Nodes: nodes }
}

func AddEdge(nodes map[int]Node, id1 int, id2 int, weight float64) {
    v, ok := nodes[id1]
    if (!ok) {
        v = Node{ Id: id1, Edges: nil }
    }
    v.Edges = append(v.Edges, Edge{ U: id1, V: id2, Weight: weight })
    nodes[id1] = v
}