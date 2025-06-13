package easy

type Question1 struct{}

func NewQuestion1() *Question1 {
	return &Question1{}
}

// 將差值存入 map 中，當遇到相同的值時，直接返回索引
func (q *Question1) TwoSum(nums []int, target int) []int {
	ansList := make(map[int]int)
	for i := range nums {
		if ansIndex, ok := ansList[nums[i]]; ok {
			return []int{ansIndex, i}
		}

		ansList[target-nums[i]] = i
	}
	return []int{}
}
