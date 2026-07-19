package solver

type SudokuSolver struct {
	state
	info           bool
	trace          bool
	processedmoves []move
}

func (s SudokuSolver) Draw() string {
	return s.board.draw()
}
