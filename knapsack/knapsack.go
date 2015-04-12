package knapsack
import (
    "strings"
    "github.com/sourcedelica/algorithms-go/util"
    "bufio"
    "os"
    "fmt"
)

type Item struct {
    value int
    weight int
}

type Knapsack struct {
    maxWeight int
    items []Item
}

func (k *Knapsack) W() int {
    return k.maxWeight
}

func (k *Knapsack) Size() int {
    return len(k.items)
}

func (k *Knapsack) Value(i int) int {
    return k.items[i - 1].value
}

func (k *Knapsack) Weight(i int) int {
    return k.items[i - 1].weight
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
        item := Item{ value: util.Atoi(parts[0]), weight: util.Atoi(parts[1]) }
        items = append(items, item)
    }

    return Knapsack{ maxWeight: weight, items: items }
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