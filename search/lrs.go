package search

// LRS Longest Repeating Substring
func LRS(s string) string {
	sa, n, lrs := NewSuffixArray(s), len(s), ""
	lcp := NewLCPArrayFromSA(s, sa)
	for i := 0; i < n-1; i++ {
		length := lcp[i]
		if length > len(lrs) {
			lrs = sa[i].Text[:length]
		}
	}
	return lrs
}
