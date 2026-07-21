package solver

type state struct {
	board           *board
	setValueMoves   []Move
	clearFloatMoves []Move
}

func new_state() state {
	var board = new_board()
	return state{&board, make([]Move, 0), make([]Move, 0)}
}

func (s *state) addMoves(moves []Move) {
	for _, m := range moves {
		s.addMove(m)
	}
	for _, m := range moves {
		if m.moveType == moveType_INIT_VALUE {
			s.board.setFixedValue(int(m.pos), uint32(m.value))
		}
	}
}

func (s *state) addMove(m Move) {
	switch m.moveType {
	case moveType_INIT_VALUE:
		s.setValueMoves = append(s.setValueMoves, m)
	case moveType_SET_VALUE:
		s.setValueMoves = append(s.setValueMoves, m)
	case moveType_CLEAR_FLOAT:
		s.clearFloatMoves = append(s.clearFloatMoves, m)
	}
}

func (s *state) hasNextMove() bool {
	return len(s.setValueMoves) > 0 || len(s.clearFloatMoves) > 0
}

func (s *state) getNextMove() Move {
	var m Move
	if len(s.setValueMoves) > 0 {
		m = s.setValueMoves[0]
		s.setValueMoves = s.setValueMoves[1:]
	} else {
		m = s.clearFloatMoves[0]
		s.clearFloatMoves = s.clearFloatMoves[1:]
	}
	return m
}
