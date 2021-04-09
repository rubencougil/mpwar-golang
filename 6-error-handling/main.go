package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("foo.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", file)
}

func myFunc(a string) (string, error) {

	if a == "Foo" {
		return "OK", nil
	}

	return "", errors.New("error")
}
