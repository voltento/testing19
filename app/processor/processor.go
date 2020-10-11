package processor

import (
	"sort"
	"strconv"
	"strings"
)

type OrderMaker struct {
}

func NewOrderMaker() *OrderMaker {
	return &OrderMaker{}
}

/// The function orders data into queue base on merging two sorted arrays.
/// It returns "No" if the order can not be reached or amount of items isn't equal for both slices.
/// It returns ordered slice in string representation if ordered successfully.
/// Complexity: time complexity not longer than 2*N*log(N), space complexity: not grater than 2*N.
/// Where N is amount of age values
func (p *OrderMaker) Process(first []int, second []int) string {
	ordered := p.order(first, second)
	if ordered == nil {
		ordered = p.order(second, first)
	}

	if len(ordered) == 0 {
		return "No"
	}

	return p.intsToString(ordered)
}

func (p *OrderMaker) intsToString(data []int) string {
	answer := ""
	for _, v := range data {
		answer += strconv.Itoa(v) + " "
	}
	return strings.Trim(answer, " ")
}

func (p *OrderMaker) order(first []int, second []int) []int {
	if len(first) != len(second) {
		return nil
	}

	length := len(first)
	if length == 0 {
		return nil
	}

	// Sort for N*log(N)
	sort.Ints(first)
	sort.Ints(second)
	index := 0

	orderedValues := make([]int, 0, 2*length)
	isFirstTurn := true
	lastValue := first[index]

	for index < length {
		if isFirstTurn {
			if lastValue > first[index] {
				return nil
			}
			lastValue = first[index]
			isFirstTurn = false
		} else {
			if lastValue > second[index] {
				return nil
			}

			lastValue = second[index]
			isFirstTurn = true
			index += 1
		}

		orderedValues = append(orderedValues, lastValue)
	}

	return orderedValues
}
