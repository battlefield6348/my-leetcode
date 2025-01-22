package easy

import (
	"strings"
)

type Question242 struct{}

func NewQuestion242() *Question242 {
	return &Question242{}
}

func (q Question242) IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sSplit := strings.Split(s, "")
	for _, v := range sSplit {
		if strings.Count(s, v) != strings.Count(t, v) {
			return false
		}
	}

	return true
}
