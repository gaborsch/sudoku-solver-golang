package solver

import "fmt"

type Exception interface {
	getMessage() string
}

func throw(e Exception) {
	err := fmt.Errorf("Exception: %s", e.getMessage())
	panic(err)
}
