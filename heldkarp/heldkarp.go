package main
import (
    "os"
    "fmt"
    "github.com/sourcedelica/algorithms-go/graph"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s filename\n", os.Args[0])
        os.Exit(1)
    }
    filename := os.Args[1]

    n, nodes := graph.ReadTSPNodes(filename)
    dist := graph.CalcEuclidDistances(nodes)

    tour := graph.TSP(n, dist)
fmt.Printf("%v\n", tour)
}
