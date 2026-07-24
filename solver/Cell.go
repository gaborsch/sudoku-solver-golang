package solver

import (
	"fmt"
	"math/bits"
)

/*
 * 00000000 000000FF FFFFFFFX VVVVCCCC
 *
 * C: bits for possible counts, values: 1-9
 * V: cell value, values: 0: not fixed, 1-9: fixed
 * X: is the value fixed? 1: yes, 0: no
 * F: Floating (possible) values
 */

type Cell struct {
	bits uint32
}

var cell_INITIAL_VALUE Cell = Cell{bits: 0x0003fe09} // not fixed, all 9 values possible

const cell_MASK_COUNT uint32 = 0x0000000f
const cell_MASK_VALUE uint32 = 0x000000f0
const cell_SHIFT_VALUE uint32 = 4
const cell_MASK_FIXED uint32 = 0x00000100
const cell_SHIFT_FIXED uint32 = 8

var cell_MASK_FLOATING_BY_VALUE = [...]uint32{0, 1 << 9, 1 << 10, 1 << 11, 1 << 12, 1 << 13, 1 << 14, 1 << 15, 1 << 16, 1 << 17}

const cell_MASK_FLOATING = 0x0003fe00
const cell_SHIFT_FLOATING = 8
const cell_MASK_ALL = 0x0003ffff

func (c *Cell) getCount() uint32 {
	return c.bits & cell_MASK_COUNT
}

func (c *Cell) getValue() uint8 {
	return uint8((c.bits & cell_MASK_VALUE) >> cell_SHIFT_VALUE)
}

func (c *Cell) isFixed() bool {
	return (c.bits&cell_MASK_FIXED)>>cell_SHIFT_FIXED == 1
}

func (c *Cell) getFloatings() uint32 {
	return (c.bits & cell_MASK_FLOATING) >> cell_SHIFT_FLOATING
}

func (c *Cell) getFixedAsFloatings() uint32 {
	// simulate floating bit for fixed value
	return (c.bits&cell_MASK_FLOATING | cell_MASK_FLOATING_BY_VALUE[(c.bits&cell_MASK_VALUE)>>cell_SHIFT_VALUE]) >> cell_SHIFT_FLOATING
}

func (c *Cell) isFloating(value uint8) bool {
	LogTrace(COMPONENT_CELL, fmt.Sprintf("isFloating: %d (%018b & %018b = %018b), %t\n", value, c.bits, cell_MASK_FLOATING_BY_VALUE[value], c.bits&cell_MASK_FLOATING_BY_VALUE[value], (c.bits&cell_MASK_FLOATING_BY_VALUE[value]) > 0))
	return (c.bits & cell_MASK_FLOATING_BY_VALUE[value]) > 0
}

/*
 * finds the cell value if only 1 possible value is present
 */
func (c *Cell) findValue() uint8 {
	if c.isFixed() {
		return c.getValue()
	}
	if c.getCount() > 1 {
		return 0
	}
	var v uint8 = 1

	for v <= 9 {
		if (c.bits & cell_MASK_FLOATING_BY_VALUE[v]) > 0 {
			return v
		}
		v++
	}
	return 0
}

func cell_setValue(value uint8) *Cell {
	return &Cell{cell_MASK_FIXED | uint32(value)<<cell_SHIFT_VALUE}
}

func (c *Cell) clearFloating(value uint8) *Cell {
	var mask = cell_MASK_FLOATING_BY_VALUE[value]
	if (c.bits & mask) == 0 {
		LogTrace(COMPONENT_CELL, fmt.Sprintf("clearFloating1 %08x %08x %t\n", c.bits, mask, (c.bits&mask) == 0))
		return c
	}
	LogTrace(COMPONENT_CELL, fmt.Sprintf("clearFloating2 %08x %08x %08x %t\n", c.bits, (c.bits&^mask)-1, (c.bits&(cell_MASK_ALL^mask))-1, (c.bits&(cell_MASK_ALL^mask))-1 == (c.bits&^mask)))
	return &Cell{(c.bits & (cell_MASK_ALL ^ mask)) - 1}
}

func (c *Cell) setFloating(value uint8) *Cell {
	var mask = cell_MASK_FLOATING_BY_VALUE[value]
	if (c.bits & mask) == mask {
		return c
	}
	return &Cell{c.bits | mask}
}

func (c *Cell) isValid() (bool, int) {
	if c.bits != (c.bits & 0x3ffff) {
		// surplus bits
		fmt.Printf("Not valid cell 1: surplus bits %8x \n", c.bits)
		return false, 1
	} else if ((c.bits & cell_MASK_FIXED) == cell_MASK_FIXED) && ((c.bits & cell_MASK_FLOATING) > 0) {
		// if fixed cannot contain floating
		fmt.Printf("Not valid cell 2: fixed cannot contain floating %8x \n", c.bits)
		return false, 2
	} else if ((c.bits & cell_MASK_FIXED) == cell_MASK_FIXED) && ((c.bits & cell_MASK_VALUE) == 0) {
		// if fixed, must contain value
		fmt.Printf("Not valid cell 3: fixed must contain value %8x \n", c.bits)
		return false, 3
	} else if ((c.bits & cell_MASK_FIXED) == 0) && ((c.bits & cell_MASK_FLOATING) == 0) {
		// if not fixed, must contain floating
		fmt.Printf("Not valid cell 4: not fixed must contain floating %8x \n", c.bits)
		return false, 4
	} else if ((c.bits & cell_MASK_FIXED) == 0) && ((c.bits & cell_MASK_VALUE) > 0) {
		// if not fixed, must not contain value
		fmt.Printf("Not valid cell 5: not fixed must not contain value %8x \n", c.bits)
		return false, 5
	} else if ((c.bits & cell_MASK_FIXED) == 0) && (bits.OnesCount32(c.bits&cell_MASK_FLOATING) != (int)(c.bits&cell_MASK_COUNT)) {
		// if not fixed, floating count must equal to the bits
		fmt.Printf("Not valid cell 9:  if not fixed, floating count must equal to the bits %8x \n", c.bits)
		return false, 9

	}
	return true, 0
}

func (c *Cell) toString() string {
	return fmt.Sprintf("Cell{bits=%08x, count=%d, value=%d, fixed=%t, floatings=%09b}", c.bits, c.getCount(), c.getValue(), c.isFixed(), c.getFloatings())
}
