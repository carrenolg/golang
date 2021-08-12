package appendint

func AppendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + 1
	// check
	if zlen <= cap(x) {
		z = x[:zlen] // extend the slice
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	//z[len(x)] = y
	copy(z[len(x):], y)
	return z
}
