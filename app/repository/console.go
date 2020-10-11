package repository

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
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
	maleAges, err := ParseInts(data)
	if err != nil {
		return nil, nil, errors.Wrap(err, fmt.Sprintf("can not parse data '%v'", data))
	}

	data = c.readData("Female ages [] - ")
	femaleAges, err := ParseInts(data)
	if err != nil {
		return nil, nil, errors.Wrap(err, fmt.Sprintf("can not parse data '%v'", data))
	}

	return maleAges, femaleAges, nil
}

func (c *Console) readData(invotation string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(invotation)
	text, _ := reader.ReadString('\n')
	return strings.Trim(text, TrimSymbols)
}
