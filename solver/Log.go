package solver

import (
	"fmt"
	"slices"
	"strings"
)

const LOG_LEVEL = const_LOG_LEVEL_INFO

const COMPONENT_BOARD = "Board"
const COMPONENT_BOARD_DRAWER = "BoardDrawer"
const COMPONENT_CELL = "Cell"
const COMPONENT_MOVE = "Move"
const COMPONENT_MOVE_GEN = "MoveGen"
const COMPONENT_STATE = "State"
const COMPONENT_SUDOKU_SOLVER = "SudokuSolver"
const COMPONENT_VECTOR_CACHE = "VectorCache"

var LOG_COMPONENTS = []string{}

type LogLevel uint8

const (
	const_LOG_LEVEL_TRACE LogLevel = iota
	const_LOG_LEVEL_DEBUG
	const_LOG_LEVEL_INFO
	const_LOG_LEVEL_WARN
	const_LOG_LEVEL_ERROR
	const_LOG_LEVEL_FATAL
	const_LOG_LEVEL_NONE
)

var LOG_LEVEL_LABELS = []string{"[ TRACE ] ", "[ DEBUG ] ", "[ INFO  ] ", "[ WARN  ] ", "[ ERROR ] ", "[ FATAL ] ", "[ NONE  ] "}

func HasTrace(component string) bool {
	return const_LOG_LEVEL_TRACE >= LOG_LEVEL && (len(LOG_COMPONENTS) == 0 || slices.Contains(LOG_COMPONENTS, component))
}

func HasDebug(component string) bool {
	return const_LOG_LEVEL_DEBUG >= LOG_LEVEL && (len(LOG_COMPONENTS) == 0 || slices.Contains(LOG_COMPONENTS, component))
}

func HasInfo() bool {
	return const_LOG_LEVEL_INFO >= LOG_LEVEL
}

func Log(level LogLevel, msg string) {
	if level >= LOG_LEVEL {
		print(LOG_LEVEL_LABELS[level])
		println(strings.TrimSuffix(msg, "\n"))
	}
}

func LogC(component string, level LogLevel, msg string) {
	if level >= LOG_LEVEL && (len(LOG_COMPONENTS) == 0 || slices.Contains(LOG_COMPONENTS, component)) {
		print(LOG_LEVEL_LABELS[level])
		if len(LOG_COMPONENTS) == 0 {
			print(fmt.Sprintf("[ %12s ]: ", _center(component, 12)))
		}
		println(strings.TrimSuffix(msg, "\n"))
	}
}

func LogTrace(component string, msg string) {
	LogC(component, const_LOG_LEVEL_TRACE, msg)
}

func LogDebug(component string, msg string) {
	LogC(component, const_LOG_LEVEL_DEBUG, msg)
}

func LogInfo(msg string) {
	Log(const_LOG_LEVEL_INFO, msg)
}
func LogWarn(msg string) {
	Log(const_LOG_LEVEL_WARN, msg)
}
func LogError(msg string) {
	Log(const_LOG_LEVEL_ERROR, msg)
}
func LogFatal(msg string) {
	Log(const_LOG_LEVEL_FATAL, msg)
	panic(msg)
}

func _center(s string, width int) string {
	if len(s) >= width {
		return s
	}
	return strings.Repeat(" ", (width-len(s))/2) + s + strings.Repeat(" ", (width-len(s)+1)/2)
}
