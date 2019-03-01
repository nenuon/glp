package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount is ...
// memo:8byteだから1byteづつずらしてカウントする
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

func PopCountLoop(x uint64) int {
	ret := 0
	for i := uint(0); i < 8; i++ {
		ret += int(pc[byte(x>>(i*8))])
	}
	return ret
}

// 2^64
func PopCountBit(x uint64) int {
	ret := 0
	for x > 0 {
		if x&1 == 1 {
			ret++
		}
		x >>= 1
	}
	return ret
}

func PopCountLeastBit(x uint64) int {
	ret := 0
	for x > 0 {
		ret++
		x = x & (x - 1)
	}
	return ret
}
