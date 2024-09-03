package easy

import (
	"strconv"
	"strings"
)

type Question1945 struct{}

func NewQuestion1945() Question1945 {
	return Question1945{}
}

// https://leetcode.com/problems/sum-of-digits-of-string-after-convert/?envType=daily-question&envId=2024-09-03
func (q *Question1945) Question1945(s string, k int) int {
	var ans int
	sList := strings.Split(s, "")

	intList := q.TurnLetterToInt(sList)
	for i := 0; i < k; i++ {
		ans = q.SumList(intList)
		intList = q.IntToList(ans)
	}
	return ans
}

func (q *Question1945) TurnLetterToInt(list []string) []int {
	intList := []int{}
	for _, v := range list {
		intList = append(intList, q.GetLetterNumber(v))
	}
	return intList
}

func (q *Question1945) SumList(intList []int) int {
	sum := 0
	for _, v := range intList {
		if v >= 10 {
			vList := q.IntToList(v)
			v = q.SumList(vList)
		}
		sum = sum + v
	}
	return sum
}

func (q *Question1945) IntToList(num int) []int {
	stringNum := strconv.Itoa(num)
	stringList := strings.Split(stringNum, "")
	intList := []int{}
	for _, v := range stringList {
		newNum, _ := strconv.Atoi(v)
		intList = append(intList, newNum)
	}
	return intList
}

func (q *Question1945) GetLetterNumber(letter string) int {
	letters := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
	}
	return letters[letter]
}
