package reverse

import "unicode/utf8"

// Reverse reverses (in place) a slice of bytes containing UTF-8
// encoded characters.
func Reverse(s []byte) {
	var p, q rune  // first and last runes
	var pz, qz int // sizes of first and last runes

	for i, j := 0, len(s); i < j-1; i, j = i+qz, j-pz {
		p, pz = utf8.DecodeRune(s[i:])
		q, qz = utf8.DecodeLastRune(s[:j])

		// shift bytes between first and last runes 
		// if size of last > first
		if qz > pz {
			copy(s[i+qz:], s[i+pz:j-qz])
		}

		// copy last rune to first
		copy(s[i:], []byte(string(q)))

		// copy first rune to last
		copy(s[j-pz:], []byte(string(p)))
	}
}
