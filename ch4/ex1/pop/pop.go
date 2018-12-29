package pop

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func SamePopCount(x, y uint64) int {
	ret := 0
	for i := uint(0); i < 64; i++ {
		if (x>>i)&1 == (y>>i)&1 {
			ret++
		}
	}
	return ret
}
