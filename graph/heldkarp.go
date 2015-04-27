package graph
import (
    "bufio"
    "github.com/sourcedelica/algorithms-go/util"
    "strings"
    "math"
)

type EuclidTSPNode struct {
    x float64
    y float64
}

type EuclidTSPNodes struct {
    Nodes []EuclidTSPNode
}

type EuclidTSP struct {
    Cost float64
    Tour []int    
}

// TSP using Held-Karp Algorithm
// http://en.wikipedia.org/wiki/Held%E2%80%93Karp_algorithm
func TSP(n int, dist [][]float64) EuclidTSP {
    numSets := (1 << uint(n))
    var highBit uint = 1 << uint(n)
    pred := make([][]uint, n + 1)
    cost := make([][]float64, numSets)
    cost[0] = make([]float64, n + 1)
    cost[1] = make([]float64, n + 1)

    var k uint
    for k = 2; k <= uint(n); k++ {
        cost[0][k] = dist[k][1]
        cost[1][k] = dist[k][1]
        var set uint = (1 << (k - 1)) | 1
        cost[set] = make([]float64, n + 1)
        cost[set][k] = dist[k][1]
    }

    var s uint
    // Subproblem size
    for s = 2; s <= uint(n); s++ {
        var set uint = (1 << s) - 1

        // For all S subset of {1, 2, ..., n} of size s
        for ; (set & highBit) == 0; set = nextSubset(set) {
            if set & 1 != 0 {

                // For all k not in S
                var k uint
                for k = 2; k <= uint(n); k++ {
                    var kmask uint = 1 << (k - 1)

                    if set & kmask == 0 {
                        cks := math.Inf(1)
                        var minm uint

                        // For each m in S, compute minimum cost of path through m to k
                        var m uint
                        for m = 2; m <= uint(n); m++ {
                            var mmask uint = 1 << (m - 1)

                            if set & mmask != 0 {
                                var notm uint = set & ^mmask  // S - {m}
                                costNoj := cost[notm][m] + dist[k][m]
                                if (costNoj < cks) {
                                    cks = costNoj
                                    minm = m
                                }
                            }
                        }
                        if len(cost[set]) == 0 {
                            cost[set] = make([]float64, n + 1)
                        }
                        cost[set][k] = cks
                        if len(pred[k]) == 0 {
                            pred[k] = make([]uint, numSets)
                        }
                        pred[k][set] = minm
                    }
                }
            }
        }
    }

    // Cost := min k != 1 { cost(k, {1, 2, ..., n}) + dist[1][k] }
    var set uint = (1 << uint(n)) - 1
    min := math.Inf(1)
    var mink uint
    for k = 2; k <= uint(n); k++ {
        var kmask uint = 1 << (k - 1)
        var notk uint = set & ^kmask
        kcost := dist[1][k] + cost[notk][k]
        if kcost < min {
            min = kcost
            mink = k
        }
    }

    // Compute tour by finding pred[p][subset] starting with the full set
    // and removing one node at a time
    tour := []int{1}
    var p uint
    for p = mink; len(pred[p]) != 0; {
        tour = append(tour, int(p))
        var pmask uint = 1 << (p - 1)
        set = set & ^pmask
        p = pred[p][set]
    }
    tour = append(tour, 1)

    return EuclidTSP{Cost: min, Tour: tour}
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

// Gosper's hack
// http://read.seas.harvard.edu/cs207/2012/?p=64
func nextSubset(x uint) uint {
    var y uint = x & -x
    var c uint = x + y
    x = (((x ^ c) >> 2) / y) | c
    return x
}