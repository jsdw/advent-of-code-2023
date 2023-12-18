package mathutils

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

// find Least Common Multiple (LCM) via GCD
func LCMSlice(ints []int) int {
	if len(ints) == 0 {
		panic("> 0 ins needed for LCM")
	}
	if len(ints) == 1 {
		return ints[0]
	}

	res := LCM(ints[0], ints[1])
	for _, n := range ints[2:] {
		res = LCM(res, n)
	}

	return res
}
