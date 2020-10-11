package main

import (
	"github.com/pkg/errors"
	"log"
	"os"
	"testing19/app/processor"
	"testing19/app/repository"
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
