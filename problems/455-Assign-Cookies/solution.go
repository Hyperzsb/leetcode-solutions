func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(a, b int) bool {
		return g[a] < g[b]
	})
	sort.Slice(s, func(a, b int) bool {
		return s[a] < s[b]
	})

	idxg, idxs := 0, 0

	for idxg < len(g) && idxs < len(s) {
		if g[idxg] <= s[idxs] {
			idxg++
			idxs++
		} else {
			idxs++
		}
	}

	return idxg
}

