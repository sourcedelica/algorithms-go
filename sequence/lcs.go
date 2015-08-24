package sequence

// Longest common subsequence
// https://en.wikipedia.org/wiki/Longest_common_subsequence_problem
//
// LCS(0, j) = 0
// LCS(i, 0) = 0
// LCS(i, j) =
// 	if s[i] == t[j] ⇒ 1 + LCS(i-1, j-1)
//  otherwise       ⇒ max[ LCS(i-2, j-1), LCS(i-1, j-2) ]

// Longest common subsequence, recursive, memoizing
func LCSrecursive(x, y string) string {
	memo := make(map[lcskey]lcs)
	return lcsRecursive(x, y, memo).subseq
}

func lcsRecursive(x, y string, memo map[lcskey]lcs) lcs {
	key := lcskey{x, y}
	var longest lcs
	if longest, ok := memo[key]; ok {
		return longest
	}
	if (len(x) == 0 || len(y) == 0) {
		longest = lcs{ 0, "" }
	} else if (last(x) == last(y)) {
		lcsr := lcsRecursive(initial(x), initial(y), memo)
		longest = lcs{ 1 + lcsr.longest, lcsr.subseq + string(last(x)) }
	} else {
		lx := lcsRecursive(initial(x), y, memo)
		ly := lcsRecursive(x, initial(y), memo)
		if (lx.longest > ly.longest) {
			longest = lx
		} else {
			longest = ly
		}
	}
	memo[key] = longest
	return longest
}

// Longest common subsequence, dynamic programming
func LCSdp(x, y string) string {
	if (len(x) == 0 || len(y) == 0) {
		return ""
	}

	memo := make([][]lcs, len(x) + 1)
	for i := 0; i <= len(x); i++ {
		memo[i] = make([]lcs, len(y) + 1)
	}

	for i := 1; i <= len(x); i++ {
		for j := 1; j <= len(y); j++ {
			var longest lcs
			if (x[i - 1] == y[j - 1]) {
				longest = lcs{ 1 + memo[i - 1][j - 1].longest,
								memo[i - 1][j - 1].subseq + string(x[i - 1]) }
			} else {
				if (memo[i][j - 1].longest > memo[i - 1][j].longest) {
					longest = memo[i][j - 1]
				} else {
					longest = memo[i - 1][j]
				}
			}
			memo[i][j] = longest
		}
	}
	return memo[len(x)][len(y)].subseq
}

type lcskey struct {
	x, y string
}

type lcs struct {
	longest int
	subseq string
}

func last(s string) uint8 {
	return s[len(s) - 1]
}

func initial(s string) string {
	return s[:len(s) - 1]
}