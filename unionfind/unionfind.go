package unionfind

// Weighted Union-Find with Path Compression
type UnionFind struct {
    id []int     // Parent
    size []int   // Size of component
    Count int    // # of components
}

// Ids are 0 to n-1
func Create(n int) UnionFind {
    var id = make([]int, n)
    for i, _ := range id { id[i] = i }
    var size = make([] int, n)
    for i, _ := range size { size[i] = 1 }
    return UnionFind { id: id, size: size, Count: n }
}

func (uf *UnionFind) Find(p int) int {
    for p != uf.id[p] {
        uf.id[p] = uf.id[uf.id[p]]  // Compress
        p = uf.id[p]
    }
    return p
}

func (uf *UnionFind) Connected(p int, q int) bool {
    return uf.Find(p) == uf.Find(q)
}

func (uf *UnionFind) Union(p int, q int) {
    pId := uf.Find(p)
    qId := uf.Find(q)
    if (pId == qId) {
        return
    }

    // Merge smaller component into larger one
    if (uf.size[pId] < uf.size[qId]) {
        uf.id[pId] = qId
        uf.size[qId] += uf.size[pId]
    } else {
        uf.id[qId] = pId
        uf.size[pId] += uf.size[qId]
    }

    uf.Count--
}