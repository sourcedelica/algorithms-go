package main
import (
    "os"
    "fmt"
    "github.com/sourcedelica/algorithms-go/util"
    "bufio"
    "strings"
    "sort"
    "github.com/sourcedelica/algorithms-go/unionfind"
)

type Edge struct {
    U int
    V int
    Distance int
}

type ByDistance []Edge

func (a ByDistance) Len() int           { return len(a) }
func (a ByDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDistance) Less(i, j int) bool { return a[i].Distance < a[j].Distance }

func main() {
    if len(os.Args) < 3 {
        fmt.Fprintf(os.Stderr, "Usage: %s filename k\n", os.Args[0])
        os.Exit(1)
    }
    filename := os.Args[1]
    k := util.Atoi(os.Args[2])

    edges := readEdges(filename)

    sort.Sort(ByDistance(edges))

    numEdges := len(edges)
    uf := unionfind.Create(numEdges)

    formClusters(k, edges, uf)

    fmt.Printf("Result for %d clusters: %d\n", k, findMinDistance(edges, uf))
}

func formClusters(k int, edges []Edge, uf unionfind.UnionFind) {
    numEdges := len(edges)
    for i := 0; i < numEdges; i++ {
        edge := edges[i]
        uf.Union(edge.U, edge.V)
        if uf.Count == k {
            break
        }
    }
}

func findMinDistance(edges []Edge, uf unionfind.UnionFind) int {
    numEdges := len(edges)
    for i := numEdges-1; i >= 0; i-- {
        edge := edges[i]
        if !uf.Connected(edge.U, edge.V) {
            return edge.Distance
        }
    }
    return 0
}

func readEdges(filename string) []Edge {
    f := util.OpenFile(filename)
   	defer f.Close()
   	scanner := bufio.NewScanner(bufio.NewReader(f))

    n := util.Atoi(util.ReadLine(scanner))
    edges := make([]Edge, 0, (n * n) / 2)

    for i := 0; i < n; i++ {
        parts := strings.Split(util.ReadLine(scanner), " ")
        p1 := util.Atoi(parts[0])
        p2 := util.Atoi(parts[1])
        distance := util.Atoi(parts[2])
        edges = append(edges, Edge{U: p1, V: p2, Distance: distance})
    }

    return edges
}