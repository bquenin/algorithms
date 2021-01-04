package search

// LCP Longest Common Prefix using vertical scanning
func LCP(strings ...string) string {
	if len(strings) == 0 {
		return ""
	}
	for i := 0; i < len(strings[0]); i++ {
		c := strings[0][i]
		for j := 1; j < len(strings); j++ {
			if i == len(strings[j]) || strings[j][i] != c {
				return strings[0][:i]
			}
		}
	}
	return strings[0]
}
