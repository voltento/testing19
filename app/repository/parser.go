package repository

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

const (
	TrimSymbols = " \n\t\n"
)

func ParseInts(data string) ([]int, error) {
	data = strings.Trim(data, TrimSymbols)
	parts := strings.Split(data, " ")
	values := make([]int, 0, len(parts))
	for _, part := range parts {
		val, err := strconv.Atoi(part)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("can not parse '%v'", part))
		}
		values = append(values, val)
	}
	return values, nil
}
