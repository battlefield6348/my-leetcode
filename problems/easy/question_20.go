package easy

import (
	"strings"
)

type Question20 struct{}

func NewQuestion20() *Question20 {
	return &Question20{}
}

func (q *Question20) IsValid(s string) bool {
	split := strings.Split(s, "")

	// 建立暫存區，沒比對過的都放進來
	temp := []string{}
	for _, v := range split {
		if len(temp) == 0 {
			temp = append(temp, v)
			continue
		}

		last := q.GetLast(temp)
		// 成功配對就扣掉最後一筆
		if q.IsMatch(last, v) {
			temp = temp[:len(temp)-1]
			continue
		}
		temp = append(temp, v)
	}
	return len(temp) == 0
}

func (q *Question20) GetLast(s []string) string {
	return s[len(s)-1]
}

func (q *Question20) IsMatch(left, right string) bool {
	matchList := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	return matchList[left] == right
}
