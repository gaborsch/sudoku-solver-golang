package solver

type State struct {
	board           *Board
	setValueMoves   []Move
	clearFloatMoves []Move
}

func new_state() *State {
	var board = new_board()
	return &State{board, make([]Move, 0), make([]Move, 0)}
}

func (s *State) addMoves(moves []Move) {
	for _, m := range moves {
		s.addMove(&m)
	}
	for _, m := range moves {
		if m.moveType == moveType_INIT_VALUE {
			s.board.setFixedValue(int(m.pos), uint32(m.value))
		}
	}
}

func (s *State) addMove(m *Move) {
	switch m.moveType {
	case moveType_INIT_VALUE:
		s.setValueMoves = append(s.setValueMoves, *m)
	case moveType_SET_VALUE:
		s.setValueMoves = append(s.setValueMoves, *m)
	case moveType_CLEAR_FLOAT:
		s.clearFloatMoves = append(s.clearFloatMoves, *m)
	}
}

func (s *State) hasNextMove() bool {
	return len(s.setValueMoves) > 0 || len(s.clearFloatMoves) > 0
}

func (s *State) getNextMove() *Move {
	var m Move
	if len(s.setValueMoves) > 0 {
		m = s.setValueMoves[0]
		s.setValueMoves = s.setValueMoves[1:]
	} else {
		m = s.clearFloatMoves[0]
		s.clearFloatMoves = s.clearFloatMoves[1:]
	}
	return &m
}
