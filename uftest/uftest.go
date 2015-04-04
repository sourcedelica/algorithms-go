package main
import (
    "github.com/sourcedelica/algorithms-go/util"
    "bufio"
    "strings"
    "github.com/sourcedelica/algorithms-go/unionfind"
    "fmt"
    "os"
)

// Reads data in the format
// # of initial components
// c1 c2
// c1 c2
// where c1 and c2 are components to union
func main() {
    filename := os.Args[1]
    f := util.OpenFile(filename)
   	defer f.Close()
   	scanner := bufio.NewScanner(bufio.NewReader(f))

    n := util.Atoi(util.ReadLine(scanner))
    uf := unionfind.Create(n)

    for i := 0; i < n; i++ {
        parts := strings.Split(util.ReadLine(scanner), " ")
        c1 := util.Atoi(parts[0])
        c2 := util.Atoi(parts[1])
        uf.Union(c1, c2)
    }

    fmt.Printf("%d components\n", uf.Count)
}
