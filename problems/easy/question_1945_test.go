package easy_test

import (
	"my-leetcode/problems/easy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestion1945(t *testing.T) {
	q := easy.NewQuestion1945()

	t.Run("test 1", func(t *testing.T) {
		ans := q.Question1945("iiii", 1)
		assert.Equal(t, 36, ans)
	})

	t.Run("test 2", func(t *testing.T) {
		ans := q.Question1945("leetcode", 2)
		assert.Equal(t, 6, ans)
	})

	t.Run("test 3", func(t *testing.T) {
		ans := q.Question1945("zbax", 2)
		assert.Equal(t, 8, ans)
	})

	t.Run("test 4", func(t *testing.T) {
		ans := q.Question1945("vbyytoijnbgtyrjlsc", 2)
		assert.Equal(t, 15, ans)
	})
}

func TestGetLetterNumber(t *testing.T) {
	q := easy.NewQuestion1945()
	t.Run("test get a series", func(t *testing.T) {
		number := q.GetLetterNumber("a")
		assert.Equal(t, 1, number)
	})

	t.Run("test get l series", func(t *testing.T) {
		number := q.GetLetterNumber("l")
		assert.Equal(t, 12, number)
	})

	t.Run("test get z series", func(t *testing.T) {
		number := q.GetLetterNumber("z")
		assert.Equal(t, 26, number)
	})
}

func TestTurnToInt(t *testing.T) {
	q := easy.NewQuestion1945()
	t.Run("test turn letter to int", func(t *testing.T) {
		intList := q.TurnLetterToInt([]string{"a", "b", "c"})
		assert.Equal(t, []int{1, 2, 3}, intList)
	})
}

func TestSumList(t *testing.T) {
	q := easy.NewQuestion1945()
	t.Run("test sum list", func(t *testing.T) {
		sum := q.SumList([]int{1, 2, 3})
		assert.Equal(t, 6, sum)
	})

	t.Run("test sum list", func(t *testing.T) {
		sum := q.SumList([]int{12, 5, 5, 20, 3, 15, 4, 5})
		assert.Equal(t, 33, sum)
	})
}

func TestIntToList(t *testing.T) {
	q := easy.NewQuestion1945()
	t.Run("test turn letter to int", func(t *testing.T) {
		sum := q.IntToList(36)
		assert.Equal(t, []int{3, 6}, sum)
	})
}
