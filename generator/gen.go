package generator

/*
#cgo CXXFLAGS: -Wall -O3 -std=c++11
#cgo LDFLAGS: -lstdc++
#include "sudokuGen.h"
*/
import "C"

func GenSudokuPuzzle() (grid [9][9]int, solution [9][9]int, difficulty int) {
	puzzle := C.gen_sudoku_puzzle();
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			grid[i][j] = int(puzzle.grid[i][j])
			solution[i][j] = int(puzzle.solution[i][j])
		}
	}
	difficulty = int(puzzle.difficulty)
	return
}
