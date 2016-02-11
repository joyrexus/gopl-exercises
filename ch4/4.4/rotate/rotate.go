package rotate

func Rotate(by int, s []int) {
	if by == 0 || len(s) == 0 {
		return
	}
	by = by % len(s)

	/*
	tmp := make([]int, by)
	copy(tmp, s[:by])
	copy(s, s[by:])
	copy(s[len(s)-by:], tmp)
	*/
}
