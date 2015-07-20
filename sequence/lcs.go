package sequence
import (
	"github.com/sourcedelica/algorithms-go/util"
)

// Longest common subsequence
// https://en.wikipedia.org/wiki/Longest_common_subsequence_problem

// Longest common subsequence, recursive, memoizing
func LCSrecursive(x, y string) int {
	memo := make(map[lcskey]int)
	return lcsRecursive(x, y, memo)
}

// TODO - return the subsequence too
func lcsRecursive(x, y string, memo map[lcskey]int) int {
	key := lcskey{x, y}
	var longest int
	if longest, ok := memo[key]; ok {
		return longest
	}
	if (len(x) == 0 || len(y) == 0) {
		longest = 0
	} else if (last(x) == last(y)) {
		longest = 1 + lcsRecursive(initial(x), initial(y), memo)
	} else {
		lx := lcsRecursive(initial(x), y, memo)
		ly := lcsRecursive(x, initial(y), memo)
		longest = util.Max(lx, ly)
	}
	memo[key] = longest
	return longest
}

// Longest common subsequence, dynamic programming
func LCSdp(x, y string) int {
	// TODO
	return 0
}

type lcskey struct {
	x, y string
}

func last(s string) uint8 {
	return s[len(s) - 1]
}

func initial(s string) string {
	return s[:len(s) - 1]
}