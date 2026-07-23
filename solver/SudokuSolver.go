package solver

import (
	"fmt"
	"slices"
)

type SudokuSolver struct {
	state          *State
	logInfo        bool
	logTrace       bool
	processedMoves []Move
}

func New_SudokuSolver() *SudokuSolver {
	return &SudokuSolver{state: new_state(), logInfo: true, logTrace: true, processedMoves: make([]Move, BOARD_SIZE)}
}

func (s *SudokuSolver) Draw() string {
	if s == nil {
		fmt.Println("s == nil")
	}
	if s.state == nil {
		fmt.Println("s.state == nil")
	}
	if s.state.board == nil {
		fmt.Println("s.state.board == nil")
	}
	return s.state.board.draw()
}

func (s *SudokuSolver) AddMoves(moves []Move) {
	s.state.addMoves(moves)
}

func (s *SudokuSolver) Solve() *Board {
	var boardHash int = 0
	loop := true
	for loop {

		for !s.state.board.isSolved() && s.state.hasNextMove() {
			m := s.state.getNextMove()
			b := s.state.board

			if (s.logTrace && boardHash != b.hashCode()) || (s.logInfo && boardHash == 0) {
				if s.logTrace {
					s.trace(b.draw())
				} else {
					s.info(b.draw())
				}
				boardHash = b.hashCode()
			}

			if !slices.ContainsFunc(s.processedMoves, func(m2 Move) bool { return m.Equals(m2) }) {
				s.info(m.toString())
				s.processedMoves = append(s.processedMoves, *m)
			}

			switch m.moveType {
			case moveType_INIT_VALUE:
				s._doInitValue(b, m)
			case moveType_SET_VALUE:
				s._doSetValue(b, m)
			case moveType_CLEAR_FLOAT:
				s._doClearFloating(b, m)
			}

		}

		if !s.state.board.isSolved() {
			// s.info("Generating moves...")
			new_moveGen(s.state).generateMoves()
			boardHash = 0
		}

		loop = (!s.state.board.isSolved() && s.state.hasNextMove())
	}
	return s.state.board
}

func (s *SudokuSolver) _doInitValue(b *Board, m *Move) {
	// fmt.Printf("_doInitValue: %s\n", m.toString())
	b.setFixedValue(int(m.pos), m.value)
	s.state.addMove(move_setValue(m.pos, m.value, "Initial check"))
}

func (s *SudokuSolver) _doSetValue(b *Board, m *Move) {
	// fmt.Printf("_doSetValue: %s\n", m.toString())
	rn := m.getRowNum()
	s._setClearFloatsTo(b, board_getRowPositions(rn), b.getRowValues(rn), m.value, "clearing "+m.getRowCoord()+" for "+m.getCoords())
	cn := m.getColNum()
	s._setClearFloatsTo(b, board_getColPositions(cn), b.getColValues(cn), m.value, "clearing "+m.getColCoord()+" for "+m.getCoords())
	bn := m.getBoxNum()
	s._setClearFloatsTo(b, board_getBoxPositions(bn), b.getBoxValues(bn), m.value, "clearing "+m.getBoxCoord()+" for "+m.getCoords())
}

func (s *SudokuSolver) _setClearFloatsTo(b *Board, positions []int, cellValues []Cell, value uint8, note string) {
	// fmt.Printf("SudokuSolver._setClearFloatsTo: pos=%v, value=%d, because %s\n", positions, value, note)

	for i, cellValue := range cellValues {
		if !cellValue.isFixed() {
			s._setClearFloatTo(b, positions[i], cellValue, value, note)
		}
	}
}

func (s *SudokuSolver) _setClearFloatTo(b *Board, pos int, c Cell, value uint8, note string) {
	// fmt.Printf("SudokuSolver._setClearFloatTo: pos=%s, value=%d, because %s\n", board_posToString(pos), value, note)
	if c.isFloating(value) {
		// fmt.Printf("_setClearFloatTo: pos=%s, value=%d, because %s\n", board_posToString(pos), value, note)
		// clear float value
		newCell := c.clearFloating(value)
		// check if there is only 1 candidate left
		fixedValue := newCell.findValue()

		if fixedValue == 0 {
			// if more than 1, then save the cleared value
			b.setCell(pos, newCell)
			// and mark the clear float for processing
			s.state.addMove(move_clearFloat(uint8(pos), value, note))
		} else {
			// if all others are cleared, save the new fixed value on the board
			b.setCell(pos, cell_setValue(fixedValue))
			// and mark the set value for processing
			s.state.addMove(move_setValue(uint8(pos), uint8(fixedValue), "single cell value found while "+note))
		}

	}
}

func (s *SudokuSolver) _doClearFloating(b *Board, m *Move) {
	// is it fixed? then nothing to do
	if b.getCell(int(m.pos)).isFixed() {
		return
	}

	// clear the floating value if it has not happened yet
	var cell = b.clearFloating(int(m.pos), m.value)

	// check if there's only one floating left, if yes, set that value
	var fixedValue = cell.findValue()
	if fixedValue != 0 {
		b.setCell(int(m.pos), cell_setValue(fixedValue))
		s.state.addMove(move_setValue(m.pos, uint8(fixedValue), "single cell value"))
	} else {
		// check if only one possibility remained in the box / row / column
		s._checkIfHasOnlyOneFloatingAtPositions(b, m.value, board_getRowPositions(m.getRowNum()), m.getRowCoord())
		s._checkIfHasOnlyOneFloatingAtPositions(b, m.value, board_getColPositions(m.getColNum()), m.getColCoord())
		s._checkIfHasOnlyOneFloatingAtPositions(b, m.value, board_getBoxPositions(m.getBoxNum()), m.getBoxCoord())
	}
}

func (s *SudokuSolver) _checkIfHasOnlyOneFloatingAtPositions(b *Board, value uint8, positions []int,
	coord string) {
	var count = s.state.board.countFloatingValue(value, positions)
	if count == 1 {
		var pos = s.state.board.getFirstFloatingValuePosition(value, positions)
		b.setFixedValue(pos, value)
		s.state.addMove(move_setValue(uint8(pos), uint8(value), "only one left in "+coord))
	}
}

func (s *SudokuSolver) info(msg string) {
	if s.logInfo {
		fmt.Println(msg)
	}
}

func (s *SudokuSolver) trace(msg string) {
	if s.logTrace {
		fmt.Println(msg)
	}
}
