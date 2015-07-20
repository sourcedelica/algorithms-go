package main
import (
	"os"
	"fmt"
	"github.com/sourcedelica/algorithms-go/sequence"
)

func main() {
	argc := len(os.Args)

	if (argc < 2 || (os.Args[1] != "recursive" && os.Args[1] != "dp")) {
		fmt.Printf("Usage: %s recursive|dp price...\n", os.Args[0])
		os.Exit(1)
	}

	x := os.Args[2]
	y := os.Args[3]

	var longest int
	if os.Args[1] == "recursive" {
		longest = sequence.LCSrecursive(x, y)
	} else {
		longest = sequence.LCSdp(x, y)
	}
	fmt.Printf("Longest common subsequence: %d\n", longest)
}
