package solver

import (
	"fmt"
	"slices"
)

type MoveGen struct {
	state *State
	moves []Move
}

func new_moveGen(s *State) *MoveGen {
	return &MoveGen{state: s, moves: make([]Move, 0)}
}

type generatorStep func() []Move

func (g *MoveGen) generateMoves() {
	var generators []generatorStep = []generatorStep{
		g.generateMovesForSinglePossibilities,
		g.generateMovesForRowColExclusion,
		g.excludeByComplementers,
		g.generateMovesForPartitioning}
	g.generateMovesWithGenerators(generators)
}

func (g *MoveGen) generateMovesWithGenerators(generators []generatorStep) {
	var movesFromStep []Move
	for i, step := range generators {
		movesFromStep = step()
		if len(movesFromStep) > 0 {
			fmt.Printf("%d moves generated from step %d", len(movesFromStep), i+1)
			break
		}
	}
	if len(movesFromStep) == 0 {
		fmt.Println("No moves generated from any step")
	}

	g.state.addMoves(movesFromStep)
}

func (g *MoveGen) generateMovesForSinglePossibilities() []Move {
	var v uint8
	for v = 1; v <= 9; v++ {
		for rn := range 9 {
			// fmt.Printf("MoveGen 1, value %d, row %d\n", v, rn+1)
			g._checkSingleValue(v, board_getRowPositions(rn), fmt.Sprintf("single value in row %d", rn+1))
		}
		for cn := range 9 {
			// fmt.Printf("MoveGen 1, value %d, column %d\n", v, cn+1)
			g._checkSingleValue(v, board_getColPositions(cn), fmt.Sprintf("single value in column %d", cn+1))
		}
		for bn := range 9 {
			// fmt.Printf("MoveGen 1, value %d, box %d\n", v, bn+1)
			g._checkSingleValue(v, board_getBoxPositions(bn), fmt.Sprintf("single value in box %d", bn+1))
		}
	}
	return g._cloneMoves()
}

func (g *MoveGen) _checkSingleValue(value uint8, positions []int, note string) {
	if g.state.board.countFloatingValue(value, positions) == 1 {
		var pos = g.state.board.getFirstFloatingValuePosition(value, positions)
		g._setValue(pos, value, note)
	}
}

func (g *MoveGen) generateMovesForRowColExclusion() []Move {
	return []Move{}
}

func (g *MoveGen) excludeByComplementers() []Move {
	return []Move{}
}

func (g *MoveGen) generateMovesForPartitioning() []Move {
	return []Move{}
}

/*
 * Set value 'v' from given positions, mark setting if needed
 */
func (g *MoveGen) _setValue(position int, value uint8, note string) {
	cell := g.state.board.getCell(position)
	if *g.state.board.setFixedValue(position, value) != *cell {
		var move Move = *move_setValue(uint8(position), uint8(value), fmt.Sprintf("value set by "+note))
		g.moves = append(g.moves, move)
	}

}

func (g *MoveGen) _cloneMoves() []Move {
	movesFromStep := slices.Clone(g.moves)
	g.moves = make([]Move, 0)
	return movesFromStep
}
