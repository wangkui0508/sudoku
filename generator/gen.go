package generator

/*
#cgo CXXFLAGS: -Wall -O3 -std=c++11
#cgo LDFLAGS: -lstdc++
#include "sudokuGen.h"
*/
import "C"

func GenSudokuPuzzle() (grid [9][9]int8, solution [9][9]int8, difficulty int) {
	puzzle := C.gen_sudoku_puzzle();
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			grid[i][j] = int8(puzzle.grid[i][j])
			solution[i][j] = int8(puzzle.solution[i][j])
		}
	}
	difficulty = int(puzzle.difficulty)
	return
}
