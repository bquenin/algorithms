package search

type LCPArray []int

func NewLCPArray(s string) LCPArray {
	return NewLCPArrayFromSA(s, NewSuffixArray(s))
}

func NewLCPArrayFromSA(s string, sa SuffixArray) LCPArray {
	n := len(s)
	lcp, rank := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		rank[sa[i].Index] = i
	}
	for i, k := 0, 0; i < n; i++ {
		if rank[i] == n-1 {
			k = 0
			continue
		}
		for j := sa[rank[i]+1].Index; i+k < n && j+k < n && s[i+k] == s[j+k]; k++ {
		}
		lcp[rank[i]] = k
		if k != 0 {
			k--
		}
	}
	return lcp
}
