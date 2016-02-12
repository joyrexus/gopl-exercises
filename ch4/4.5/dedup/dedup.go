package dedup

// Dedup removes adjacent duplicates in a slice of strings.
func Dedup(items []string) []string {
	i := 0
	last := ""; 
	for _, s := range items {
		if s != last {
			items[i] = s
			last = s
			i++
		}
	}
	return items[:i]
}
