package main
import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "github.com/sourcedelica/algorithms-go/util"
    "github.com/sourcedelica/algorithms-go/unionfind"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s filename\n", os.Args[0])
        os.Exit(1)
    }

    numIds, bits, nodes := readNodes(os.Args[1])
    uf := unionfind.Create(numIds)

    formClusters(bits, nodes, &uf)
    fmt.Println(uf.Count)
}

func xorValues(bits int) []int {
    xors := make([]int, 0)
    for i := 0; i < bits; i++ {
        xors = append(xors, bitAt(i))
    }
    for i := 0; i < bits; i++ {
        for j := i+1; j < bits; j++ {
            xors = append(xors, bitAt(i) | bitAt(j))
        }
    }
    return xors
}

func bitAt(i int) int {
    mask := 1
    for ; i > 0; i-- {
        mask = mask << 1
    }
    return mask
}

func formClusters(bits int, nodes map[int]int, uf *unionfind.UnionFind) {
    masks := xorValues(bits)
    for node, id := range nodes {
        for _, mask := range masks {
            mate := node ^ mask
            mateId, ok := nodes[mate]
            if (ok) {
                uf.Union(id, mateId)
            }
        }
    }
}

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

    return id, bits, nodes
}
