package main

import (
	"bufio"
	"fmt"
	"gaborsch/sudoku-solver/solver"
	"os"
	"regexp"
	"strings"
)

func main() {
	solver := new_sudokuSolverCli()
	solver.run()
}

type sudokuSolverCli struct {
	solver  *solver.SudokuSolver
	line    string
	linePos int
}

func new_sudokuSolverCli() sudokuSolverCli {
	return sudokuSolverCli{solver: solver.New_SudokuSolver(), line: "", linePos: 0}
}

func (s *sudokuSolverCli) reInit() {
	s.solver = solver.New_SudokuSolver()
}

func (s *sudokuSolverCli) run() {
	s.help()
	scanner := bufio.NewScanner(os.Stdin)
	isRunning := true

	for isRunning && scanner.Scan() {
		s.line = scanner.Text()
		s.linePos = 0
		Log("Input: '" + s.line + "'")

		switch {
		case s.check([]string{"help", "h"}):
			s.help()
		case s.check([]string{"board", "b"}):
			Log("board case")
			s.board(scanner)
		case s.check([]string{"example", "e"}):
			Log("example case")
			s.example()
		case s.check([]string{"set", "s"}):
			Log("set case")
		case s.check([]string{"clear", "c"}):
			Log("clear case")
		case s.check([]string{"exit", "x", "quit", "q"}):
			msg("Exiting...")
			isRunning = false
		default:
			msg("Unknown command")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading stdin:", err)
	}

}

func (s *sudokuSolverCli) check(words []string) bool {
	for _, word := range words {
		if strings.HasPrefix(s.line[s.linePos:], word) {
			s.linePos += len(word)
			return true
		}
	}
	return false
}

func (s *sudokuSolverCli) readAll() string {
	return s.line[s.linePos:]
}

func (s *sudokuSolverCli) readInt() int {
	var v int
	cnt, err := fmt.Sscanf(s.readAll(), "%d", &v)
	if cnt > 0 {
		s.linePos += len(fmt.Sprintf("%d", v))
		return v
	}
	panic(err)
}

func (s *sudokuSolverCli) example() {
	sample := SAMPLES[s.readInt()]
	s.reInit()

	moves := boardReader_getBoard(sample)
	s.solver.AddMoves(moves)
	if s == nil {
		fmt.Println("sudokuSolverCli s == nil")
	}
	if s.solver == nil {
		fmt.Println("sudokuSolverCli s.solver == nil")
	}

	msg(s.solver.Draw())
}

func (s *sudokuSolverCli) board(scanner *bufio.Scanner) {
	s.reInit()

	ROW_PATTERN := regexp.MustCompile(`^[1-9 ]{0,9}$`)
	var sb strings.Builder
	for i := range 9 {
		match := false
		for !match {
			match = true
			fmt.Printf("Row %d: ", i+1)
			if scanner.Scan() {
				row := scanner.Text()
				// Log("Board Input: '" + row + "'")
				if ROW_PATTERN.MatchString(row) {
					sb.WriteString(row)
				} else {
					match = false
					msg("Invalid values, try again!")
				}
			}
		}
	}
	moves := boardReader_getBoard(sb.String())
	s.solver.AddMoves(moves)
	msg(s.solver.Draw())
}

func Log(log string) {
	fmt.Println(log)
}

func msg(msg string) {
	fmt.Println(msg)
}

func (s sudokuSolverCli) help() {
	fmt.Println("Sudoku solver by gaborsch, (c) 2023-2026")
}

func boardReader_getBoard(s string) []solver.Move {
	var moves = make([]solver.Move, 0, solver.BOARD_SIZE)
	s = strings.ReplaceAll(s, "\r\n", "\n")

	for i, line := range strings.Split(s, "\n") {
		line += "         "
		for j, c := range line[:9] {
			if c >= '1' && c <= '9' {
				moves = append(moves, *solver.Move_initValue((uint8)(i*9+j), (uint8)(c-'0'), "initial"))
			}
		}
	}
	return moves
}
