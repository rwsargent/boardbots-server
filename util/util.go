package util

type Position struct {
	Row, Col int;
}

func AbsInt(val int) int {
	y := val >> 31
	return (val ^ y) - y;
}
