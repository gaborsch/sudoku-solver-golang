package solver

import "math/bits"

/*
 * 00000000 000000FF FFFFFFFX VVVVCCCC
 *
 * C: bits for possible counts, values: 1-9
 * V: cell value, values: 0: not fixed, 1-9: fixed
 * X: is the value fixed? 1: yes, 0: no
 * F: Floating (possible) values
 */

type cell struct {
	bits uint32
}

var cell_INITIAL_VALUE cell = cell{bits: 0x0003fe09} // not fixed, all 9 values possible

const cell_MASK_COUNT uint32 = 0x0000000f
const cell_MASK_VALUE uint32 = 0x000000f0
const cell_SHIFT_VALUE uint32 = 4
const cell_MASK_FIXED uint32 = 0x00000100
const cell_SHIFT_FIXED uint32 = 8

var cell_MASK_FLOATING_BY_VALUE = [...]uint32{0, 1 << 9, 1 << 10, 1 << 11, 1 << 12, 1 << 13, 1 << 14, 1 << 15, 1 << 16, 1 << 17}

const cell_MASK_FLOATING = 0x0003fe00
const cell_SHIFT_FLOATING = 8

func (c cell) getCount() uint32 {
	return c.bits & cell_MASK_COUNT
}

func (c cell) getValue() uint32 {
	return (c.bits & cell_MASK_VALUE) >> cell_SHIFT_VALUE
}

func (c cell) isFixed() bool {
	return (c.bits&cell_MASK_FIXED)>>cell_SHIFT_FIXED == 1
}

func (c cell) getFloatings() uint32 {
	return (c.bits & cell_MASK_FLOATING) >> cell_SHIFT_FLOATING
}

func (c cell) getFixedAsFloatings() uint32 {
	// simulate floating bit for fixed value
	return (c.bits&cell_MASK_FLOATING | cell_MASK_FLOATING_BY_VALUE[(c.bits&cell_MASK_VALUE)>>cell_SHIFT_VALUE]) >> cell_SHIFT_FLOATING
}

func (c cell) isFloating(value uint32) bool {
	return (c.bits & cell_MASK_FLOATING_BY_VALUE[value]) > 0
}

/*
 * finds the cell value if only 1 possible value is present
 */
func (c cell) findValue() uint32 {
	if c.isFixed() {
		return c.getValue()
	}
	if c.getCount() > 1 {
		return 0
	}
	var v uint32 = 1

	for v <= 9 {
		if (c.bits & cell_MASK_FLOATING_BY_VALUE[v]) > 0 {
			return v
		}
		v++
	}
	return 0
}

func cell_setValue(value uint32) cell {
	return cell{cell_MASK_FIXED | value<<cell_SHIFT_VALUE}
}

func (c cell) clearFloating(value uint32) cell {
	var mask = cell_MASK_FLOATING_BY_VALUE[value]
	if (c.bits & mask) == 0 {
		return c
	}
	return cell{c.bits &^ mask}
}

func (c cell) setFloating(value uint32) cell {
	var mask = cell_MASK_FLOATING_BY_VALUE[value]
	if (c.bits & mask) == mask {
		return c
	}
	return cell{c.bits | mask}
}

func (c cell) isValid() (bool, int) {
	if c.bits != (c.bits & 0x3ffff) {
		// surplus bits
		return false, 1
	} else if ((c.bits & cell_MASK_FIXED) == cell_MASK_FIXED) && ((c.bits & cell_MASK_FLOATING) > 0) {
		// if fixed cannot contain floating
		return false, 2
	} else if ((c.bits & cell_MASK_FIXED) == cell_MASK_FIXED) && ((c.bits & cell_MASK_VALUE) == 0) {
		// if fixed, must contain value
		return false, 3
	} else if ((c.bits & cell_MASK_FIXED) == 0) && ((c.bits & cell_MASK_FLOATING) == 0) {
		// if not fixed, must contain floating
		return false, 4
	} else if ((c.bits & cell_MASK_FIXED) == 0) && ((c.bits & cell_MASK_VALUE) > 0) {
		// if not fixed, must not contain value
		return false, 5
	} else if ((c.bits & cell_MASK_FIXED) == 0) && (bits.OnesCount32(c.bits&cell_MASK_FLOATING) != (int)(c.bits&cell_MASK_COUNT)) {
		// if not fixed, floating count must equal to the bits
		return false, 9

	}
	return true, 0
}

func (c cell) toString() {

}

/*

	public static String toString(int bits) {
		StringBuilder sb = new StringBuilder();
		if (isFixed(bits)) {
			sb.append("Fixed: ").append(getValue(bits));
		} else {
			sb.append("Floating (").append(getCount(bits)).append(") ");
			for (int i = 1; i <= 9; i++) {
				sb.append(isFloating(bits, i) ? (char)('0'+ i) : '_');
			}
		}
		sb.append("  ").append(Integer.toBinaryString(bits));

		return sb.toString();
	}

	public static String explain(int bits) {
		StringBuilder sb = new StringBuilder();
		if (isFixed(bits)) {
			sb.append("Fixed ");
		}
		sb.append("value: ").append(getValue(bits));
		sb.append(", Floatings: (").append(getCount(bits)).append(") ");
		for (int i = 1; i <= 9; i++) {
			sb.append(isFloating(bits, i) ? (char)('0'+ i) : '_');
		}
		sb.append("  Binary: ").append(Integer.toBinaryString(bits));

		return sb.toString();
	}
*/
