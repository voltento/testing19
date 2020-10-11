package repository

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"strings"
)

type Console struct {
}

func NewConsole() *Console {
	return &Console{}
}

func (c *Console) Provide() ([]int, []int, error) {
	fmt.Print("Please, print the ages\n")
	var data string

	data = c.readData("Male ages [] - ")
	maleAges, err := c.parseInts(data)
	if err != nil {
		return nil, nil, errors.Wrap(err, fmt.Sprintf("can not parse data '%v'", data))
	}

	data = c.readData("Female ages [] - ")
	femaleAges, err := c.parseInts(data)
	if err != nil {
		return nil, nil, errors.Wrap(err, fmt.Sprintf("can not parse data '%v'", data))
	}

	return maleAges, femaleAges, nil
}

func (c *Console) readData(invotation string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(invotation)
	text, _ := reader.ReadString('\n')
	return strings.Trim(text, " \n\t\r")
}

func (c *Console) parseInts(data string) ([]int, error) {
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
