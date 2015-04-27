package main
import (
    "github.com/sourcedelica/algorithms-go/graph"
    "os"
    "fmt"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s filename\n", os.Args[0])
        os.Exit(1)
    }
    filename := os.Args[1]

    n, dists := graph.ReadTSPDistances(filename)

    tour := graph.TSP(n, dists)

    fmt.Printf("%v\n", tour)
}