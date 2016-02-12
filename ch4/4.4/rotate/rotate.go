package rotate

// Rotate returns a slice of ints based on an input slice 
// where it's items are rotated by a certain count either 
// right or left.
func Rotate(by int, s []int, right bool) []int {
	if by == 0 || len(s) == 0 {
		return s
	}
	// assume we're rotating left
	i := by % len(s)
	j := len(s) - i
	if right {
		i, j = j, i	// swap if rotating right
	}
	q := make([]int, len(s))
	copy(q, s[i:])
	copy(q[j:], s[:i])
	return q
}
