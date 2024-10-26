package easy

import (
	"strings"
)

type Question13 struct{}

func NewQuestion13() *Question13 {
	return &Question13{}
}

func (q *Question13) RomanToInt(s string) int {
	split := strings.Split(s, "")

	ans := 0
	for i := 0; i < len(split); i++ {
		target := split[len(split)-1-i]
		toInt := q.getInt(target)

		if ans >= 5 && target == "I" {
			toInt = 0 - toInt
		}

		if ans >= 50 && target == "X" {
			toInt = 0 - toInt
		}

		if ans >= 500 && target == "C" {
			toInt = 0 - toInt
		}

		ans = ans + toInt
	}
	return ans
}

func (q *Question13) getInt(s string) int {
	switch s {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	}

	return 0
}
