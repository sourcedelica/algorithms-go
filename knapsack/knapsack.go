package knapsack
import (
    "strings"
    "bufio"
    "os"
    "fmt"
    "github.com/sourcedelica/algorithms-go/util"
)

type Item struct {
    value int
    weight int
}

type Knapsack struct {
    maxWeight int
    items []Item
}

type key struct {
    i, x int
}

type Result struct {
    Items []int
    Weight int
    Value int
}

func (sack *Knapsack) W() int {
    return sack.maxWeight
}

func (sack *Knapsack) SetW(w int) {
    sack.maxWeight = w
}

func (sack *Knapsack) Size() int {
    return len(sack.items)
}

func (sack *Knapsack) Value(i int) int {
    return sack.items[i - 1].value
}

func (sack *Knapsack) Weight(i int) int {
    return sack.items[i - 1].weight
}

func (sack *Knapsack) Item(i int) Item {
    return sack.items[i - 1]
}

// Solve using recursion and memoization
func (sack *Knapsack) Find() Result {
    memo := make(map[key]Result)

    result := sack.find(sack.Size(), sack.W(), memo)
    return result
}

func (sack *Knapsack) find(i int, x int, memo map[key]Result) Result {
    if i == 0 {
        return Result{}
    }

    key := key{i, x}
    if result, ok := memo[key]; ok {
        return result
    }

    w := sack.Weight(i)
    v := sack.Value(i)
    var result Result

    if w > x {
        result = sack.find(i - 1, x, memo)
    } else {
        leftRes := sack.find(i - 1, x, memo)
        rightRes := sack.find(i - 1, x - w, memo)
        if (rightRes.Value + v > leftRes.Value) {
            result = Result{Items: append(rightRes.Items, i), Weight: rightRes.Weight + w, Value: rightRes.Value + v}
        } else {
            result = leftRes
        }
    }

    memo[key] = result
    return result
}

// Load knapsack from a file in the format
// maxWeight #items
// value1 weight1
// value2 weight2
// ...
func LoadFrom(filename string) Knapsack {
    f := util.OpenFile(filename)
    defer f.Close()
    scanner := bufio.NewScanner(bufio.NewReader(f))

    parts := strings.Split(util.ReadLine(scanner), " ")
    weight := util.Atoi(parts[0])
    numItems := util.Atoi(parts[1])

    items := make([]Item, 0, numItems)

    for scanner.Scan() {
        parts := strings.Split(scanner.Text(), " ")
        item := Item{value: util.Atoi(parts[0]), weight: util.Atoi(parts[1])}
        items = append(items, item)
    }

    return Knapsack{maxWeight: weight, items: items}
}

// Load a knapsack using argv[1]
func Load() Knapsack {
    if (len(os.Args) < 2) {
        fmt.Fprintf(os.Stderr, "Usage: %s filename\n", os.Args[0])
        os.Exit(1)
    }
    filename := os.Args[1]

    return LoadFrom(filename)
}

func New(maxWeight int, items []Item) Knapsack {
    return Knapsack{maxWeight: maxWeight, items: items}
}