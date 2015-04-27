package graph
import (
    "bufio"
    "github.com/sourcedelica/algorithms-go/util"
    "strings"
    "math"
    "fmt"
)

type EuclidTSPNode struct {
    x float64
    y float64
}

type EuclidTSPNodes struct {
    Nodes []EuclidTSPNode
}

//type euclidTSP struct {
//    dist [][]float32
//    path [][]float32
//}

type EuclidTSP struct {
    Cost float64
    Tour []int    
}

// TSP using Held-Karp Algorithm
// http://en.wikipedia.org/wiki/Held%E2%80%93Karp_algorithm
func TSP(n int, dist [][]float64) EuclidTSP {
    printDistances(dist)

    cost := make([][]float64, n + 1)
    numSets := (1 << uint(n))
    cost[1] = make([]float64, numSets)
    var highBit uint = 1 << uint(n)

    var k uint
    for k = 2; k <= uint(n); k++ {
        cost[k] = make([]float64, numSets)
        cost[k][0] = dist[k][1]
        cost[k][1] = dist[k][1]
        var set uint = (1 << (k - 1)) | 1
        cost[k][set] = dist[k][1]
//fmt.Printf("cost[%d][%2d]=dist[%d][1]=%4.1f\n", k, set, k, dist[k][1])
    }

    var s uint
    // Subproblem size
    for s = 2; s <= uint(n); s++ {
        var set uint = (1 << s) - 1

        // For all S subset of {1, 2, ..., n} of size s
        for ; (set & highBit) == 0; set = nextSubset(set, highBit) {
            if set & 1 != 0 {
//fmt.Printf("s=%d set=%2d %08b\n", s, set, set)

                // For all k in S
                var k uint
                for k = 2; k <= uint(n); k++ {
                    var kmask uint = 1 << (k - 1)

                    if set & kmask == 0 {
                        cks := math.Inf(1)
//                        var notk uint = set & ^kmask
//fmt.Printf("  k=%d kmask=%08b\n", k, kmask)

                        // For each node m != k and m in S
                        var m uint
                        for m = 2; m <= uint(n); m++ {
                            var mmask uint = 1 << (m - 1)
                            if (set & mmask != 0) {
//fmt.Printf("    m=%d mmask=%08b\n", m, mmask)
//
                                var notm uint = set & ^mmask
//fmt.Printf("      cost[%d][%2d]=%4.1f dist[%d][%2d]=%4.1f\n", m, notm, cost[m][notm], k, m, dist[k][m])
                                costNoj := cost[m][notm] + dist[k][m]
                                if (costNoj < cks) {
                                    cks = costNoj
                                }
                            }
                        }
                        cost[k][set] = cks
//fmt.Printf("    cost[%d][%2d]=%4.1f\n", k, set, cks)
                    }
                }
            }
        }
    }

    // Cost := min k != 1 { cost(k, {1, 2, ..., n}) + dist[1][k] }
    var set uint = (1 << uint(n)) - 1
    var min = math.Inf(1)
    for k = 2; k <= uint(n); k++ {
        var kmask uint = 1 << (k - 1)
        var notk uint = set & ^kmask
        kcost := dist[1][k] + cost[k][notk]
        if kcost < min {
            min = kcost
        }
    }

    return EuclidTSP{Cost: min}  // TODO - Tour
}

// Create symmetric distance matrix based on Euclidean coordinates
// Returns two-dimensional slice indexed 1..n (same as nodes)
// Ignore the 0th row and column
func CalcEuclidDistances(nodes EuclidTSPNodes) [][]float64 {
    n := len(nodes.Nodes) - 1
    dist := make([][]float64, n + 1)

    for i := 1; i <= n; i++ {
        dist[i] = make([]float64, n + 1)
    }

    for i := 1; i <= n; i++ {
        for j := i + 1; j <= n; j++ {
            xsq := nodes.Nodes[i].x - nodes.Nodes[j].x
            xsq *= xsq
            ysq := nodes.Nodes[i].y - nodes.Nodes[j].y
            ysq *= ysq
            sq := math.Sqrt(xsq + ysq)
            dist[i][j] = sq
            dist[j][i] = sq
        }
    }
    return dist
}

// Reads a file containing nodes in the format
// #nodes
// node1x node1y
// ...
// nodenx nodeny
// Returns the number of nodes and an a slice of TSPNodes indexed 1 to n
// The 0th entry should be ignored
func ReadTSPNodes(filename string) (int, EuclidTSPNodes) {
    f := util.OpenFile(filename)
   	defer f.Close()
   	scanner := bufio.NewScanner(bufio.NewReader(f))

    size := util.Atoi(util.ReadLine(scanner))
    nodes := make([]EuclidTSPNode, size + 1)
    
    for i := 1; i <= size; i++ {
        coord := strings.Split(util.ReadLine(scanner), " ")
        x := util.Atof(coord[0])
        y := util.Atof(coord[1])
        nodes[i] = EuclidTSPNode{x, y}
    }
    
    return size, EuclidTSPNodes{nodes}
}

// Reads a file containing distances in the format
// #nodes
// d11 d12 ... d1n
// d21 d22 ... d2n
// ...
// dn1 dn2 ... dnn
// Returns the number of nodes and a 2-dimensional slice indexed 1 to n
// The 0th row and column should be ignored
func ReadTSPDistances(filename string) (int, [][]float64) {
    f := util.OpenFile(filename)
   	defer f.Close()
   	scanner := bufio.NewScanner(bufio.NewReader(f))

    n := util.Atoi(util.ReadLine(scanner))
    dists := make([][]float64, n + 1)
    for i := 1; i <= n; i++ {
        dists[i] = make([]float64, n + 1)
        sdist := strings.Split(util.ReadLine(scanner), " ")
        for j := 0; j < n; j++ {
            dists[i][j + 1] = util.Atof(sdist[j])
        }
    }
    return n, dists
}

func printDistances(dist [][]float64) {
    n := len(dist) - 1

    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            fmt.Printf("%2.1f ", dist[i][j])
        }
        fmt.Println()
    }
}

func nextSubset(x uint, highBit uint) uint {
    var y uint = x & -x
    var c uint = x + y
    x = (((x ^ c) >> 2) / y) | c
//    if ((x & highBit) != 0) {
//        x = ((x & (highBit - 1)) << 2) | 3
//    }
    return x
}