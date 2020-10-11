package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/voltento/testing19/app/processor"
	"github.com/voltento/testing19/app/repository"
	"log"
	"os"
)

func main() {
	provider := repository.NewConsole()

	maleAges, femaleAges, err := provider.Provide()
	if err != nil {
		err = errors.Wrap(err, "can not load ages")
		log.Print(err.Error())
		os.Exit(1)
	}

	p := processor.NewOrderMaker()
	answer := p.Process(maleAges, femaleAges)
	fmt.Print(answer)
}
