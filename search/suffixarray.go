package search

import (
	"github.com/bquenin/algorithms/math"
	"sort"
)

type SuffixArray []string

func NewSuffixArray(s string) SuffixArray {
	sa := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		sa[i] = s[i:]
	}
	sort.Strings(sa)
	return sa
}
func (sa SuffixArray) LCPArray() []int {
	lcp := make([]int, len(sa))
	for i := 1; i < len(sa); i++ {
		lcp[i] = sa.LCP(i)
	}
	return lcp
}

func (sa SuffixArray) LCP(i int) int {
	if i <= 0 || i >= len(sa) {
		return -1
	}
	return sa.lcp(sa[i-1], sa[i])
}

func (sa SuffixArray) lcp(s, t string) int {
	n := math.Min(len(s), len(t))
	for i := 0; i < n; i++ {
		if s[i] != t[i] {
			return i
		}
	}
	return n
}
