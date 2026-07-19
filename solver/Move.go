package solver

type moveType uint8

const (
	moveType_INIT_VALUE = iota
	moveType_SET_VALUE
	moveType_CLEAR_FLOAT
)

type move struct {
	moveType
	pos   uint8
	value uint8
	note  string
}

func move_initValue(pos uint8, value uint8, note string) move {
	return move{moveType_INIT_VALUE, pos, value, note}
}

func move_setValue(pos uint8, value uint8, note string) move {
	return move{moveType_SET_VALUE, pos, value, note}
}
func move_clearFloat(pos uint8, value uint8, note string) move {
	return move{moveType_CLEAR_FLOAT, pos, value, note}
}
