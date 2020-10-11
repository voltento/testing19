package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInts_IncorrectSymbols(t *testing.T) {
	var err error

	_, err = ParseInts("1 2 q")
	assert.Error(t, err)

	_, err = ParseInts("1 2 !")
	assert.Error(t, err)

	_, err = ParseInts("]")
	assert.Error(t, err)
}

func TestParseInts_Correct(t *testing.T) {
	var err error
	var values []int

	values, err = ParseInts("1 2 3")
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, values)

	values, err = ParseInts("1")
	assert.NoError(t, err)
	assert.Equal(t, []int{1}, values)

	values, err = ParseInts("1 ")
	assert.NoError(t, err)
	assert.Equal(t, []int{1}, values)
}
