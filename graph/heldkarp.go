package graph
import (
    "bufio"
    "github.com/sourcedelica/algorithms-go/util"
    "strings"
    "math"
)

type EuclidTSPNode struct {
    x float32
    y float32
}

type EuclidTSPNodes struct {
    Nodes []EuclidTSPNode
}

type MinimumTSP struct {
    Cost float32
    Tour []int    
}

// TSP using Held-Karp Algorithm
// http://en.wikipedia.org/wiki/Held%E2%80%93Karp_algorithm
func TSP(N int, dist [][]float32) MinimumTSP {
    n := uint(N)
    two := uint(2)
    infinity := float32(math.Inf(1))
    numSets := uint(1 << n)  // 2^n
    pred := make([][]byte, n + 1)
    cost := make([][]float32, numSets)
    cost[1] = make([]float32, n + 1)

    // Compute cost[{node}, 1] for each node
    for k := two; k <= n; k++ {
        cost[1][k] = dist[k][1]
        set := (1 << (k - 1)) | 1
        cost[set] = make([]float32, n + 1)
    }

    // Subproblem size goes from 2..n
    for s := two; s <= n; s++ {
        // Turn on all bits in set, ie, 0011 for s == 2
        set := (1 << s) - uint(1)

        // For all S subset of {1, 2, ..., n} of size s
        for ; (set & numSets) == 0; set = nextSubset(set) {
            if set & 1 != 0 {

                // For all k not in S
                for k := two; k <= n; k++ {
                    kmask := bitAt(k)

                    if set & kmask == 0 {   // k not in S
                        cks := infinity
                        var minm uint

                        // For each m in S, compute minimum cost of path through m to k
                        for m := two; m <= n; m++ {
                            mmask := bitAt(m)

                            if set & mmask != 0 {   // m in S
                                notm := set & ^mmask  // S - {m}
                                costNoj := cost[notm][m] + dist[k][m]
                                if (costNoj < cks) {
                                    cks = costNoj
                                    minm = m
                                }
                            }
                        }
                        // Record minimum and corresponding node
                        if len(cost[set]) == 0 {
                            cost[set] = make([]float32, n + 1)
                        }
                        cost[set][k] = cks
                        if len(pred[k]) == 0 {
                            pred[k] = make([]byte, numSets)
                        }
                        pred[k][set] = byte(minm)
                    }
                }
            }
        }
    }

    // Cost := min k != 1 { cost(k, {1, 2, ..., n} - k) + dist[1][k] }
    set := uint((1 << n) - 1)  // Set bit for each node in the set
    min := infinity
    var mink uint
    for k := two; k <= n; k++ {
        kmask := bitAt(k)
        notk := set & ^kmask   // set - {k}
        kcost := cost[notk][k] + dist[1][k]
        if kcost < min {
            min = kcost
            mink = k
        }
    }

    // Compute tour by finding pred[p][subset] starting with the full set
    // and removing one node at a time
    tour := []int{1}
    for p := uint(mink); len(pred[p]) != 0; p = uint(pred[p][set]) {
        tour = append(tour, int(p))
        pmask := bitAt(p)
        set = set & ^pmask
    }
    tour = append(tour, 1)

    return MinimumTSP{Cost: min, Tour: tour}
}

// Set the bit at the ith position
func bitAt(i uint) uint {
    return 1 << (i - 1)
}

// Create symmetric distance matrix based on Euclidean coordinates
// Returns two-dimensional slice indexed 1..n (same as nodes)
// Ignore the 0th row and column
func CalcEuclidDistances(nodes EuclidTSPNodes) [][]float32 {
    n := len(nodes.Nodes) - 1
    dist := make([][]float32, n + 1)

    for i := 1; i <= n; i++ {
        dist[i] = make([]float32, n + 1)
    }

    for i := 1; i <= n; i++ {
        for j := i + 1; j <= n; j++ {
            xsq := nodes.Nodes[i].x - nodes.Nodes[j].x
            xsq *= xsq
            ysq := nodes.Nodes[i].y - nodes.Nodes[j].y
            ysq *= ysq
            sq := float32(math.Sqrt(float64(xsq + ysq)))
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
        x := float32(util.Atof(coord[0]))
        y := float32(util.Atof(coord[1]))
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
func ReadTSPDistances(filename string) (int, [][]float32) {
    f := util.OpenFile(filename)
   	defer f.Close()
   	scanner := bufio.NewScanner(bufio.NewReader(f))

    n := util.Atoi(util.ReadLine(scanner))
    dists := make([][]float32, n + 1)
    for i := 1; i <= n; i++ {
        dists[i] = make([]float32, n + 1)
        sdist := strings.Split(util.ReadLine(scanner), " ")
        for j := 0; j < n; j++ {
            dists[i][j + 1] = float32(util.Atof(sdist[j]))
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