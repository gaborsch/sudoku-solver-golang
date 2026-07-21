package solver

import "fmt"

type moveType uint8

const (
	moveType_INIT_VALUE = iota
	moveType_SET_VALUE
	moveType_CLEAR_FLOAT
)

var moveType_toString []string = []string{"INIT_VALUE", "SET_VALUE", "CLEAR_FLOAT"}

type Move struct {
	moveType
	pos   uint8
	value uint8
	note  string
}

func Move_initValue(pos uint8, value uint8, note string) *Move {
	return &Move{moveType_INIT_VALUE, pos, value, note}
}

func move_setValue(pos uint8, value uint8, note string) *Move {
	return &Move{moveType_SET_VALUE, pos, value, note}
}
func move_clearFloat(pos uint8, value uint8, note string) *Move {
	return &Move{moveType_CLEAR_FLOAT, pos, value, note}
}

func (m *Move) getCoords() string {
	return fmt.Sprintf("row %d, col %d", board_getRowNum(int(m.pos)), board_getColNum(int(m.pos)))
}

func (m *Move) getRowCoord() string {
	return fmt.Sprintf("row %d", board_getRowNum(int(m.pos)))
}

func (m *Move) getColCoord() string {
	return fmt.Sprintf("col %d", board_getColNum(int(m.pos)))
}

func (m *Move) getBoxCoord() string {
	return fmt.Sprintf("box %d", board_getBoxNum(int(m.pos)))
}

func (m *Move) getRowNum() int {
	return board_getRowNum(int(m.pos))
}

func (m *Move) getColNum() int {
	return board_getColNum(int(m.pos))
}

func (m *Move) getBoxNum() int {
	return board_getBoxNum(int(m.pos))
}

func (m *Move) toString() string {
	if m.note == "" {
		return fmt.Sprintf("%s %d, %s", moveType_toString[m.moveType], m.value, board_posToString(int(m.pos)))
	} else {
		return fmt.Sprintf("%s %d, %s (%s)", moveType_toString[m.moveType], m.value, board_posToString(int(m.pos)), m.note)
	}
}

func (m Move) Equals(m2 Move) bool {
	return m.moveType == m2.moveType &&
		m.pos == m2.pos &&
		m.value == m2.value
}
