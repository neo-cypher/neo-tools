package BaseConvert

func BaseConvert(x int, base int) []int {
	r := []int{}
	for x > 0 {
		r = append(r, x%base)
		x /= base
	}
	return r
}
