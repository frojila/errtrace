package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/frojila/errtrace"
)

func main() {
	// create new error
	err1 := errtrace.New("test error one")

	// wrap error without message
	err2 := errtrace.Wrap(err1)

	// wrap error with message
	err3 := errtrace.Message("test error three").Wrap(err2)

	log.Print(err3)

	// check if err3 contain err1
	if errors.Is(err3, err1) {
		fmt.Println("err3 is contain err1")
	}

	// check if err is a valid errtrace
	ok := errtrace.Valid(err3)
	fmt.Println(ok) // should print true

	ok = errtrace.Valid(errors.New("any-errror"))
	fmt.Println(ok) // should print false
}
