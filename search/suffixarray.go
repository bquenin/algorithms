package search

import "sort"

type Suffix struct {
	Text  string
	Index int
}

type SuffixArray []*Suffix

// sort.Interface
func (sa SuffixArray) Len() int           { return len(sa) }
func (sa SuffixArray) Swap(i, j int)      { sa[i], sa[j] = sa[j], sa[i] }
func (sa SuffixArray) Less(i, j int) bool { return sa[i].Text < sa[j].Text }

func NewSuffixArray(s string) SuffixArray {
	sa := make(SuffixArray, len(s))
	for i := 0; i < len(s); i++ {
		sa[i] = &Suffix{string(s[i:]), i}
	}
	sort.Sort(sa)
	return sa
}

func (sa *SuffixArray) LCP(i int) int { return len(LCP((*sa)[i-1].Text, (*sa)[i].Text)) }
