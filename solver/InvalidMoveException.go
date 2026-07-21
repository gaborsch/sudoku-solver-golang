package solver

import (
	"fmt"
	"strings"
)

type InvalidMoveException struct {
	pos    int
	bits   uint32
	errors []string
}

func (e InvalidMoveException) getMessage() string {
	return fmt.Sprintf("Invalid move at %s, bits: %8x: %s", board_posToString(e.pos), e.bits, strings.Join(e.errors, ", "))
}
