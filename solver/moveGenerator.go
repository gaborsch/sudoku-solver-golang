package solver

type MoveGenerator struct {
	state *State
}

func new_moveGenerator(s *State) *MoveGenerator {
	return &MoveGenerator{s}
}

func (g *MoveGenerator) generateMoves() {

}
