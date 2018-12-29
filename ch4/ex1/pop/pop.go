package pop

var pc [256]byte

func DiffPopCount(x, y uint64) int {
	ret := 0
	for i := uint(0); i < 64; i++ {
		if (x>>i)&1 != (y>>i)&1 {
			ret++
		}
	}
	return ret
}
