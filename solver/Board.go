package solver

import (
	"fmt"
)

const BOARD_SIZE = 81

const board_SAFE_MODE = true

type board struct {
	board [BOARD_SIZE]cell
}

var board_BOXES = [][]int{{0, 1, 2, 9, 10, 11, 18, 19, 20}, {3, 4, 5, 12, 13, 14, 21, 22, 23},
	{6, 7, 8, 15, 16, 17, 24, 25, 26}, {27, 28, 29, 36, 37, 38, 45, 46, 47},
	{30, 31, 32, 39, 40, 41, 48, 49, 50}, {33, 34, 35, 42, 43, 44, 51, 52, 53},
	{54, 55, 56, 63, 64, 65, 72, 73, 74}, {57, 58, 59, 66, 67, 68, 75, 76, 77},
	{60, 61, 62, 69, 70, 71, 78, 79, 80}}

var board_ROWS = [][]int{{0, 1, 2, 3, 4, 5, 6, 7, 8}, {9, 10, 11, 12, 13, 14, 15, 16, 17},
	{18, 19, 20, 21, 22, 23, 24, 25, 26}, {27, 28, 29, 30, 31, 32, 33, 34, 35},
	{36, 37, 38, 39, 40, 41, 42, 43, 44}, {45, 46, 47, 48, 49, 50, 51, 52, 53},
	{54, 55, 56, 57, 58, 59, 60, 61, 62}, {63, 64, 65, 66, 67, 68, 69, 70, 71},
	{72, 73, 74, 75, 76, 77, 78, 79, 80}}

var board_COLS = [][]int{{0, 9, 18, 27, 36, 45, 54, 63, 72}, {1, 10, 19, 28, 37, 46, 55, 64, 73},
	{2, 11, 20, 29, 38, 47, 56, 65, 74}, {3, 12, 21, 30, 39, 48, 57, 66, 75},
	{4, 13, 22, 31, 40, 49, 58, 67, 76}, {5, 14, 23, 32, 41, 50, 59, 68, 77},
	{6, 15, 24, 33, 42, 51, 60, 69, 78}, {7, 16, 25, 34, 43, 52, 61, 70, 79},
	{8, 17, 26, 35, 44, 53, 62, 71, 80}}

var board_BOXES_AND_ROWS = [][][]int{
	{{0, 1, 2}, {9, 10, 11}, {18, 19, 20}, {}, {}, {}, {}, {}, {}},
	{{3, 4, 5}, {12, 13, 14}, {21, 22, 23}, {}, {}, {}, {}, {}, {}},
	{{6, 7, 8}, {15, 16, 17}, {24, 25, 26}, {}, {}, {}, {}, {}, {}},
	{{}, {}, {}, {27, 28, 29}, {36, 37, 38}, {45, 46, 47}, {}, {}, {}},
	{{}, {}, {}, {30, 31, 32}, {39, 40, 41}, {48, 49, 50}, {}, {}, {}},
	{{}, {}, {}, {33, 34, 35}, {42, 43, 44}, {51, 52, 53}, {}, {}, {}},
	{{}, {}, {}, {}, {}, {}, {54, 55, 56}, {63, 64, 65}, {72, 73, 74}},
	{{}, {}, {}, {}, {}, {}, {57, 58, 59}, {66, 67, 68}, {75, 76, 77}},
	{{}, {}, {}, {}, {}, {}, {60, 61, 62}, {69, 70, 71}, {78, 79, 80}}}

var board_BOXES_AND_COLS = [][][]int{
	{{0, 9, 18}, {1, 10, 19}, {2, 11, 20}, {}, {}, {}, {}, {}, {}},
	{{}, {}, {}, {3, 12, 21}, {4, 13, 22}, {5, 14, 23}, {}, {}, {}},
	{{}, {}, {}, {}, {}, {}, {6, 15, 24}, {7, 16, 25}, {8, 17, 26}},
	{{27, 36, 45}, {28, 37, 46}, {29, 38, 47}, {}, {}, {}, {}, {}, {}},
	{{}, {}, {}, {30, 39, 48}, {31, 40, 49}, {32, 41, 50}, {}, {}, {}},
	{{}, {}, {}, {}, {}, {}, {33, 42, 51}, {34, 43, 52}, {35, 44, 53}},
	{{54, 63, 72}, {55, 64, 73}, {56, 65, 74}, {}, {}, {}, {}, {}, {}},
	{{}, {}, {}, {57, 66, 75}, {58, 67, 76}, {59, 68, 77}, {}, {}, {}},
	{{}, {}, {}, {}, {}, {}, {60, 69, 78}, {61, 70, 79}, {62, 71, 80}}}

func new_board() *board {
	var b board = board{}
	for i := range BOARD_SIZE {
		b.board[i] = cell_INITIAL_VALUE
	}
	return &b
}

func clone_board(orig *board) *board {
	var b board = board{}
	b.board = orig.board
	return &b
}

func board_getRowPositions(rowIndex int) []int {
	return board_ROWS[rowIndex]
}

func board_getColPositions(colIndex int) []int {
	return board_COLS[colIndex]
}

func board_getBoxPositions(boxIndex int) []int {
	return board_BOXES[boxIndex]
}

func (b *board) fetchByMapping(mapping []int) []cell {
	v := make([]cell, len(mapping)) // TODO: VectorCache.fetchByMappingVector
	for i, m := range mapping {
		v[i] = b.board[m]
	}
	return v
}

func (b *board) getRowValues(rowIndex int) []cell {
	return b.fetchByMapping(board_ROWS[rowIndex])
}

func (b *board) getColValues(colIndex int) []cell {
	return b.fetchByMapping(board_COLS[colIndex])
}

func (b *board) getBoxValues(boxIndex int) []cell {
	return b.fetchByMapping(board_BOXES[boxIndex])
}

func board_intersectBoxRow(boxNum int, rowNum int) []int {
	return board_BOXES_AND_ROWS[boxNum][rowNum]
}

func board_intersectBoxCol(boxNum int, colNum int) []int {
	return board_BOXES_AND_COLS[boxNum][colNum]
}

/*
 * counts how many times the given floating value exists at the given positions
 */
func (b *board) countFloatingValue(value uint32, positions []int) int {
	var count int = 0
	for _, pos := range positions {
		if b.board[pos].isFloating(value) {
			count++
		}
	}
	return count
}

/*
 * returns the first position of a floating value using position array
 */
func (b *board) getFirstFloatingValuePosition(value uint32, positions []int) int {
	for _, pos := range positions {
		if b.board[pos].isFloating(value) {
			return pos
		}
	}
	return -1
}

func (b *board) setFixedValue(pos int, value uint32) *cell {
	return b._writeBoard(pos, cell_setValue(value))
}

func (b *board) getCell(pos int) *cell {
	return &b.board[pos]
}

func (b *board) setCell(pos int, c *cell) {
	b._writeBoard(pos, c)
}

func (b *board) clearFloating(pos int, value uint32) *cell {
	return b._writeBoard(pos, b.board[pos].clearFloating(value))
}

func (b *board) _writeBoard(pos int, c *cell) *cell {
	if board_SAFE_MODE {
		cellValid, errCode := c.isValid()
		if !cellValid {
			throw(invalidMoveException{pos, c.bits, []string{fmt.Sprintf("Cell error: %d", errCode)}})
		}
		b.board[pos] = *c
		board_errors := b._isValid()
		if len(board_errors) > 0 {
			throw(invalidMoveException{pos, c.bits, board_errors})
		}
	} else {
		b.board[pos] = *c
	}

	return c
}

func board_getRowNum(pos int) int {
	return pos / 9
}

func board_getColNum(pos int) int {
	return pos % 9
}

func board_getBoxNum(pos int) int {
	return (pos/27)*3 + (pos%9)/3
}

func (b *board) draw() string {
	return new_boardDrawer(b).draw()
}

func (b *board) _isSolved() bool {
	for _, c := range b.board {
		if !c.isFixed() {
			return false
		}
	}
	return true
}

func board_posToString(pos int) string {
	return fmt.Sprintf("row %d, column %d", board_getRowNum(pos)+1, board_getColNum(pos)+1)
}

func (b *board) _isValid() []string {
	var errors []string
	for i, posArr := range board_BOXES {
		valid, error := b._checkValidValues(posArr)
		if !valid {
			errors = append(errors, fmt.Sprintf("%s in box %d", error, i+1))
		}
	}
	for i, posArr := range board_ROWS {
		valid, error := b._checkValidValues(posArr)
		if !valid {
			errors = append(errors, fmt.Sprintf("%s in row %d", error, i+1))
		}
	}
	for i, posArr := range board_COLS {
		valid, error := b._checkValidValues(posArr)
		if !valid {
			errors = append(errors, fmt.Sprintf("%s in col %d", error, i+1))
		}
	}
	return errors
}

func (b *board) _checkValidValues(posArr []int) (bool, string) {
	fixedValues := make([]int, 10) // TODO: VectorCache.checkValidValues
	var floatings uint32 = 0
	for _, pos := range posArr {
		floatings |= b.board[pos].getFixedAsFloatings()
		if b.board[pos].isFixed() {
			fixedValues[b.board[pos].getValue()]++
		}
	}
	for i, fixedValue := range fixedValues {
		if fixedValue > 1 {
			return false, fmt.Sprintf("Duplicate value: %d", i)
		}
	}
	if (floatings & 0x3fe) != 0x3fe {
		return false, fmt.Sprintf("Cleared all possible values (%x)", floatings)
	}
	return true, ""
}
