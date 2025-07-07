package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x byte) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(x & 1)
		x >>= 1
	}
	return sum
}
