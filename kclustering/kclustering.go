package main
import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "sort"
    "github.com/sourcedelica/algorithms-go/unionfind"
    "github.com/sourcedelica/algorithms-go/util"
)

type edge struct {
    u int
    v int
    distance int
}

type byDistance []edge

func (a byDistance) Len() int           { return len(a) }
func (a byDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byDistance) Less(i, j int) bool { return a[i].distance < a[j].distance }

// Find maximum spacing for single-link k-clustering
func main() {
    if len(os.Args) < 3 {
        fmt.Fprintf(os.Stderr, "Usage: %s filename k\n", os.Args[0])
        os.Exit(1)
    }
    filename := os.Args[1]
    k := util.Atoi(os.Args[2])

    n, edges := readEdges(filename)

    sort.Sort(byDistance(edges))

    uf := unionfind.Create(n)

    formClusters(k, edges, &uf)

    maxSpacing := findMinDistance(edges, uf)

    fmt.Printf("Result for %d clusters: %d\n", k, maxSpacing)
}

// Form clusters of edges by closest distance
// until number of clusters reaches k
func formClusters(k int, edges []edge, uf *unionfind.UnionFind) {
    numEdges := len(edges)
    for i := 0; i < numEdges; i++ {
        edge := edges[i]
        uf.Union(edge.u, edge.v)
        if uf.Count == k {
            break
        }
    }
}

// Find edge with minimum distance where
// points of edge are not in the same cluster
func findMinDistance(edges []edge, uf unionfind.UnionFind) int {
    numEdges := len(edges)
    for i := 0; i < numEdges; i++ {
        edge := edges[i]
        if !uf.Connected(edge.u, edge.v) {
            return edge.distance
        }
    }
    return -1
}

// Read file in the format
// #edges
// p1 p2 distance
// ...
// Where p1 and p2 are point ids
func readEdges(filename string) (int, []edge) {
    f := util.OpenFile(filename)
   	defer f.Close()
   	scanner := bufio.NewScanner(bufio.NewReader(f))

    n := util.Atoi(util.ReadLine(scanner))
    edges := make([]edge, 0, (n * n) / 2)

    for scanner.Scan() {
        parts := strings.Split(scanner.Text(), " ")
        p1 := util.Atoi(parts[0])
        p2 := util.Atoi(parts[1])
        distance := util.Atoi(parts[2])
        edges = append(edges, edge{u: p1 - 1, v: p2 - 1, distance: distance})
    }

    return n, edges
}