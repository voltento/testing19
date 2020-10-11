package processor

import (
	"errors"
	"fmt"
)

/// Stores amount of persons by age, where index of item is age and value is amount of persons of this age
type AgeStorage struct {
	agesCount []int
	nextAge   int
}

func NewAgeStorage(ages []int) (*AgeStorage, error) {
	storage := &AgeStorage{
		agesCount: make([]int, MaxAge),
		nextAge:   MaxAge,
	}
	for _, age := range ages {
		if err := storage.AddAge(age); err != nil {
			return nil, err
		}
	}
	return storage, nil
}

func (s *AgeStorage) AddAge(age int) error {
	if MaxAge <= age || age < 0 {
		return errors.New(fmt.Sprintf("unexpected age: %v", age))
	}
	s.agesCount[age] += 1
	if age < s.nextAge {
		s.nextAge = age
	}

	return nil
}

func (s *AgeStorage) Empty() bool {
	return s.nextAge >= MaxAge
}

func (s *AgeStorage) NextAge() (int, bool) {
	if s.nextAge >= MaxAge {
		return MaxAge, false
	}

	age := s.nextAge
	s.agesCount[age] -= 1
	s.nextAge = s.findNext()
	return age, true
}

func (s *AgeStorage) findNext() int {
	for index := s.nextAge; index < MaxAge; index += 1 {
		if s.agesCount[index] >= 1 {
			return index
		}
	}
	return MaxAge
}
