package processor

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	MaxAge = 250
)

type OrderMaker struct {
}

func NewOrderMaker() *OrderMaker {
	return &OrderMaker{}
}

/// The function stores data into array by ages and then go through thees arrays and order data.
/// It returns "No" if the order can not be reached or amount of items isn't equal for both slices.
/// It returns ordered slice in string representation if ordered successfully.
/// Complexity: time complexity O(n)=2N+M, space complexity: not grater than 2N+M
/// Where N is amount of age values and M is max age == 250
/// Time complexity explanation:
/// Step 1: We use an array for storing persons' age. Saving all ages takes N time.
/// Step 2: We iterate from minYear to MaxYear in worst case it takes MaxYears iterations.
/// And we can get extra iteration for each age value, in worst case it's N - amount of persons' age
/// Step 3: Parse data into string, is linear
func (p *OrderMaker) Process(first []int, second []int) string {
	const No = "No"

	ordered, err := p.order(first, second)
	if err == nil {
		return p.intsToString(ordered)
	}

	ordered, err = p.order(second, first)
	if err != nil {
		return No
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

/// Ordering data by counting sort
func (p *OrderMaker) order(first []int, second []int) ([]int, error) {
	if len(first) != len(second) {
		return nil, errors.New("data sets have unequal length")
	}

	length := len(first)
	if length == 0 {
		return nil, errors.New("data set is empty")
	}

	firstStorage, err := NewAgeStorage(first)
	if err != nil {
		return nil, err
	}

	secondStorage, err := NewAgeStorage(second)
	if err != nil {
		return nil, err
	}

	orderedValues := make([]int, 0, 2*length)
	isFirstTurn := true
	lastValue := 0

	for !firstStorage.Empty() || !secondStorage.Empty() {
		if isFirstTurn {
			age, success := firstStorage.NextAge()
			if !success {
				return nil, errors.New("first storage is empty")
			}
			if age < lastValue {
				return nil, buildAgeOrderError(age)
			}
			lastValue = age
			isFirstTurn = false
		} else {
			age, success := secondStorage.NextAge()
			if !success {
				return nil, errors.New("second storage is empty")
			}
			if age < lastValue {
				return nil, buildAgeOrderError(age)
			}
			lastValue = age
			isFirstTurn = true
		}

		orderedValues = append(orderedValues, lastValue)
	}

	return orderedValues, nil
}

func buildAgeOrderError(age int) error {
	return errors.New(fmt.Sprintf("the value %v cannot be ordered", age))
}
