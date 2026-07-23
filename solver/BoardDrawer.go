package solver

import (
	"fmt"
	"strings"
)

type BoardDrawer struct {
	Board
	sb strings.Builder
}

func new_boardDrawer(b *Board) *BoardDrawer {
	return &BoardDrawer{*b, strings.Builder{}}
}

func (bd *BoardDrawer) draw() string {
	bd.drawSeparator(0)
	for i := range 9 {
		nums := bd.getRowValues(i)
		bd.drawLine(nums, 0)
		bd.drawLine(nums, 1)
		bd.drawLine(nums, 2)
		bd.drawSeparator(i + 1)
	}
	bd.drawSummary()
	return bd.sb.String()
}

func (bd *BoardDrawer) drawLine(nums []Cell, line int) {
	for j, num := range nums {
		if j%3 == 0 {
			bd.sb.WriteString("|")
		}
		bd.drawCell(num, line)
		bd.sb.WriteString("|")
	}
	bd.drawNl()
}

func (bd *BoardDrawer) drawCell(c Cell, line int) {
	if c.isFixed() {
		switch line {
		case 0:
			bd.sb.WriteString(" /-\\ ")
		case 1:
			fmt.Fprintf(&bd.sb, "( %d )", c.getValue())
		case 2:
			bd.sb.WriteString(" \\-/ ")
		}
	} else {
		for i := range 3 {
			var value = (uint8)(line*3 + i + 1)
			if c.isFloating(value) {
				fmt.Fprintf(&bd.sb, "%d", value)
			} else {
				bd.sb.WriteString(" ")
			}
			if i < 2 {
				bd.sb.WriteString(" ")
			}
		}
	}

}

func (bd *BoardDrawer) drawSeparator(row int) {
	if row%3 == 0 {
		bd.sb.WriteString(strings.Repeat("+=====+=====+=====+", 3))
	} else {
		bd.sb.WriteString(strings.Repeat("+-----+-----+-----+", 3))
	}
	bd.drawNl()
}

func (bd *BoardDrawer) drawSummary() {
	counts := make([]int, 10) // TODO: VectorCache.drawSummaryVector
	var totalFloats uint32 = 0
	for _, c := range bd.Board.board {
		if c.isFixed() {
			counts[c.getValue()]++
		} else {
			totalFloats += c.getCount()
		}
	}
	var total int = 0
	for i := 1; i < len(counts); i++ {
		fmt.Fprintf(&bd.sb, " %d:%d ", i, 9-counts[i])
		total += 9 - counts[i]
	}
	fmt.Fprintf(&bd.sb, "   Left: %d (%d)", total, totalFloats)
	bd.drawNl()
}

func (bd *BoardDrawer) drawNl() {
	bd.sb.WriteString("\n")
}
