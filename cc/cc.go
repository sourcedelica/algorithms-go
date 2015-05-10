package main
import (
    "os"
    "fmt"
    "github.com/sourcedelica/algorithms-go/graph"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s filename start\n", os.Args[0])
        os.Exit(1)
    }
    filename := os.Args[1]
    ugraph := graph.ReadUGraph(filename)

    components := ugraph.CC()

    fmt.Printf("%v\n", components)
}