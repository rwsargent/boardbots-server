package util

type Position struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

func AbsInt(val int) int {
	y := val >> 31
	return (val ^ y) - y
}
