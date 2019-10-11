package __error_handling

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

	fmt.Println("%s", file)
}

func myFunc(a string) (string, error) {

	if a == "Foo" {
		return "OK", nil
	}

	return "", errors.New("Error!")
}
