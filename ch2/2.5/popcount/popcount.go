package popcount

// pc[i] is the population count of i.
var pc [256]byte // 256 = 8^2 possible combinations of 0 and 1 for 8 slots

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountLoop returns the population count (number of set bits) of x.
func PopCountLoop(x uint64) int {
	var s byte
	for i := uint64(0); i < 8; i++ {
		s += pc[byte(x>>(i*8))]
	}
	return int(s)
}

// PopCountShift returns the population count (number of set bits) of x.
func PopCountShift(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}

// PopCountClear returns the population count (number of set bits) of x.
func PopCountClear(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1)
		n++
	}
	return n
}
