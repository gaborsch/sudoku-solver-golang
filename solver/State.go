package solver

type state struct {
	board
	setValuemoves   []move
	clearFloatmoves []move
}
