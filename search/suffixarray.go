package search

import "sort"

type Suffix struct {
	Text  string
	Index int
}

type SuffixArray []*Suffix

func (t SuffixArray) Len() int           { return len(t) }
func (t SuffixArray) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t SuffixArray) Less(i, j int) bool { return t[i].Text < t[j].Text }

func NewSuffixArray(s string) SuffixArray {
	n := len(s)
	sa := make(SuffixArray, n)
	for i := 0; i < n; i++ {
		sa[i] = &Suffix{s[i:], i}
	}
	sort.Sort(sa)
	return sa
}
