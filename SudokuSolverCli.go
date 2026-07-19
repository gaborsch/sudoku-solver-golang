package main

import (
	"bufio"
	"fmt"
	"gaborsch/sudoku-solver/solver"
	"os"
	"strings"
)

type sudokuSolverCli struct {
	solver solver.SudokuSolver
}

func main() {
	solver := sudokuSolverCli{solver.SudokuSolver{}}
	solver.run()
}

func (s sudokuSolverCli) run() {
	s.help()
	scanner := bufio.NewScanner(os.Stdin)
	isRunning := true

	for isRunning && scanner.Scan() {
		line := scanner.Text()
		Log("Input: '" + line + "'")
		switch {
		case line == "h" || strings.HasPrefix(line, "help"):
			s.help()
		case line == "b" || strings.HasPrefix(line, "board"):
			Log("board case")
			s.board()
		case line == "s" || strings.HasPrefix(line, "set"):
			Log("set case")
		case line == "c" || strings.HasPrefix(line, "clear"):
			Log("clear case")
		case line == "x" || strings.HasPrefix(line, "exit"):
			Log("exit case")
			isRunning = false
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading stdin:", err)
	}

}

func (s sudokuSolverCli) board() {
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
