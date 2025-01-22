package easy

type Question217 struct{}

func NewQuestion217() *Question217 {
	return &Question217{}
}
func (q Question217) ContainsDuplicate(nums []int) bool {
	list := map[int]bool{}
	for i := range nums {
		if _, isExist := list[nums[i]]; isExist {
			return true
		}
		list[nums[i]] = false
	}
	return false
}
