package solver

type moveGenerator struct {
	state *state
}

func new_moveGenerator(s *state) *moveGenerator {
	return &moveGenerator{s}
}

func (g *moveGenerator) generateMoves() {

}
