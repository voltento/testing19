package processor

import (
	"github.com/gitchander/permutation"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	No = "No"
)

// Test all possible permutations
func doTest(t *testing.T, first, second []int, expect string) {
	firstPerm := permutation.New(permutation.IntSlice(first))
	nextFirstPerm := true
	for nextFirstPerm {

		secondPerm := permutation.New(permutation.IntSlice(second))
		nextSecondPerm := true
		for nextSecondPerm {
			processor := NewOrderMaker()
			result := processor.Process(first, second)
			assert.Equal(t, expect, result)

			result = processor.Process(second, first)
			assert.Equal(t, expect, result)

			nextSecondPerm = secondPerm.Next()
		}

		nextFirstPerm = firstPerm.Next()
	}
}

func TestOrderMaker_SandBox(t *testing.T) {
	doTest(t, []int{1, 2}, []int{1, 4}, "1 1 2 4")
}

func TestOrderMaker_OutOfRane(t *testing.T) {
	doTest(t, []int{1, 2}, []int{1, 250}, No)
	doTest(t, []int{1, 2}, []int{1, -1}, No)
}

func TestOrderMaker_DiffLength(t *testing.T) {
	doTest(t, []int{}, []int{1}, No)
	doTest(t, []int{2}, []int{1, 3}, No)
	doTest(t, []int{2}, []int{1, 3, 5}, No)
}

func TestOrderMaker_BothEmpty(t *testing.T) {
	doTest(t, []int{}, []int{}, No)
}

func TestOrderMaker_Correct(t *testing.T) {
	doTest(t, []int{1, 3}, []int{2, 5}, "1 2 3 5")
	doTest(t, []int{1, 3}, []int{2, 249}, "1 2 3 249")
	doTest(t, []int{0, 3}, []int{2, 5}, "0 2 3 5")
	doTest(t, []int{3, 1, 6}, []int{7, 2, 5}, "1 2 3 5 6 7")
	doTest(t, []int{4, 2, 8, 6, 10}, []int{1, 3, 9, 7, 5}, "1 2 3 4 5 6 7 8 9 10")
}

func TestOrderMaker_Equal(t *testing.T) {
	doTest(t, []int{1, 1, 1}, []int{1, 1, 1}, "1 1 1 1 1 1")
	doTest(t, []int{1, 1, 2}, []int{1, 1, 2}, "1 1 1 1 2 2")
	doTest(t, []int{1, 1, 2}, []int{1, 1, 3}, "1 1 1 1 2 3")

	doTest(t, []int{1, 2, 2, 5}, []int{1, 1, 2, 4}, "1 1 1 2 2 2 4 5")
}

func TestOrderMaker_Wrong(t *testing.T) {
	doTest(t, []int{1, 2, 3}, []int{3, 4, 5}, No)
	doTest(t, []int{4, 2, 3, 6}, []int{1, 3, 5, 7}, No)
	doTest(t, []int{1, 1, 3, 6}, []int{2, 2, 4, 7}, No)
}
