package solver

import "fmt"

type SudokuSolver struct {
	state          *state
	info           bool
	trace          bool
	processedmoves []Move
}

func New_SudokuSolver() *SudokuSolver {
	return &SudokuSolver{state: new_state(), info: true, trace: true, processedmoves: make([]Move, BOARD_SIZE)}
}

func (s *SudokuSolver) Draw() string {
	fmt.Println("Draw()")
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
