package search

// LRS Longest Repeating Substring
func LRS(s string) string {
	sa, n, lrs := NewSuffixArray(s), len(s), ""
	for i := 1; i < n; i++ {
		length := sa.LCP(i)
		if length > len(lrs) {
			lrs = sa[i][:length]
		}
	}
	return lrs
}
