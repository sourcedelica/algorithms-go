package main
import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "github.com/sourcedelica/algorithms-go/util"
    "github.com/sourcedelica/algorithms-go/unionfind"
)

// Using single-link clustering, find number of clusters
// where Hamming distance is at most 2
func main() {
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s filename\n", os.Args[0])
        os.Exit(1)
    }

    numIds, bits, nodes := readNodes(os.Args[1])
    uf := unionfind.Create(numIds)

    formClusters(bits, nodes, &uf)

    fmt.Printf("Number of clusters: %d\n", uf.Count)
}

// Compute all combinations of values 'bits' long
// with one or two bits set
func bitMasks(bits int) []int {
    masks := make([]int, 0)
    for i := 0; i < bits; i++ {
        masks = append(masks, bitAt(i))
    }
    for i := 0; i < bits; i++ {
        for j := i+1; j < bits; j++ {
            masks = append(masks, bitAt(i) | bitAt(j))
        }
    }
    return masks
}

// Return value with bit set at ith position
func bitAt(i int) int {
    return 1 << uint(i)
}

// Form clusters of nodes where the # bits differ by at most 2
func formClusters(bits int, nodes map[int]int, uf *unionfind.UnionFind) {
    masks := bitMasks(bits)

    // For each node
    for node, id := range nodes {
        // For each bit mask
        for _, mask := range masks {
            // Compute possible mate using XOR
            mate := node ^ mask
            // If mate exists, merge with it
            mateId, ok := nodes[mate]
            if (ok) {
                uf.Union(id, mateId)
            }
        }
    }
}

// Read file in format
// #nodes #bits
// b1 b2 ... bn [for node 1]
// b2 b2 ... bn [for node 2]
// ...
// Where bi is one bit of the node value
func readNodes(filename string) (int, int, map[int]int) {
    f := util.OpenFile(filename)
    defer f.Close()
    scanner := bufio.NewScanner(bufio.NewReader(f))

    parts := strings.Split(util.ReadLine(scanner), " ")
    bits := util.Atoi(parts[1])

    nodes := make(map[int]int)
    id := 0

    for scanner.Scan() {
        line := scanner.Text()
        line = strings.Replace(line, " ", "", -1)
        node := int(util.Btoi(line))
        if _, ok := nodes[node]; !ok {
            nodes[node] = id
            id++
        }
    }

    // id is now number of ids
    return id, bits, nodes
}
