package solver

type state struct {
	board           *board
	setValuemoves   []Move
	clearFloatmoves []Move
}

func new_state() state {
	var board = new_board()
	return state{&board, make([]Move, 0), make([]Move, 0)}
}

func (s *state) addMoves(moves []Move) {
	// TODO: moves.forEach(this::addMove);
	for _, m := range moves {
		if m.moveType == moveType_INIT_VALUE {
			s.board.setFixedValue(int(m.pos), uint32(m.value))
		}
	}
}
