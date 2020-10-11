package main

import (
	"github.com/pkg/errors"
	"github.com/voltento/testing19/app/processor"
	"github.com/voltento/testing19/app/repository"
	"log"
	"os"
)

func main() {
	var r repository.Repository

	men, women, err := r.Ages()
	if err != nil {
		err = errors.Wrap(err, "can not load ages")
		log.Print(err.Error())
		os.Exit(1)
	}

	p := processor.NewOrderMaker()
	answer := p.Process(men, women)
	log.Println(answer)
}
