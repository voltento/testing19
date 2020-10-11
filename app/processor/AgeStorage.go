package processor

import (
	"errors"
	"fmt"
)

const (
	maxAge = 250
)

/// Stores amount of persons by age, where index of item is age and value is amount of persons of this age
type AgeStorage struct {
	agesCount []int
}

func NewAgeStorage(ages []int) (*AgeStorage, error) {
	storage := &AgeStorage{agesCount: make([]int, maxAge)}
	for _, age := range ages {
		if err := storage.AddAge(age); err != nil {
			return nil, err
		}
	}
	return storage, nil
}

func (s *AgeStorage) AddAge(age int) error {
	if len(s.agesCount) >= age {
		return errors.New(fmt.Sprintf("unexpected age: %v", age))
	}
	s.agesCount[age] += 1
	return nil
}

func (s *AgeStorage) TakeAge(age int) bool {
	if len(s.agesCount) >= age {
		return false
	}

	if s.agesCount[age] > 0 {
		s.agesCount[age] -= 1
		return true
	}

	return false
}
