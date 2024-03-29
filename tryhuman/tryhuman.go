package main

import (
	"fmt"

	"github.com/wangkui0508/sudoku/generator"
	"github.com/wangkui0508/sudoku/humansolver"
	"github.com/wangkui0508/sudoku/fastsolver"
	"github.com/wangkui0508/sudoku/ui"
)


func main() {
	//try()
	test()
}

func try() {
	totalTry := 0
	totalSuccess := 0
	for {
		values, solution, difficulty := generator.GenSudokuPuzzle()

		//for y := 0; y < 9; y++ {
		//	for x := 0; x < 9; x++ {
		//		grid[y][x] = int8(values[y][x])
		//	}
		//}
		humansolver.PrintStrGrid(humansolver.InitStrGrid(solution[:]))

		if difficulty < 100 {
			continue
		}
		totalTry++
		if totalTry > 100 {
			break
		}
		if totalTry % 10 == 0 {
			fmt.Printf("Now totalTry %d\n", totalTry)
		}
		var grid [9][9]int8
		for y := 0; y < 9; y++ {
			for x := 0; x < 9; x++ {
				grid[y][x] = int8(values[y][x])
			}
		}
		success := humansolver.SimpleLoop(grid[:])
		if success {
			totalSuccess++
		} else {
			break
		}
	}
	fmt.Printf("totalSuccess %d\n", totalSuccess)
}

func test() {
	//grid := [9][9]int8{
	//	{3, 0, 4,    0, 8, 7,    2, 0, 5},
	//	{0, 7, 0,    3, 6, 0,    8, 1, 9},
	//	{1, 0, 0,    9, 0, 0,    0, 4, 0},

	//	{0, 1, 5,    8, 7, 6,    4, 0, 3},
	//	{7, 6, 0,    4, 0, 2,    0, 0, 8},
	//	{2, 0, 0,    5, 3, 1,    0, 7, 6},

	//	{8, 2, 0,    6, 0, 0,    0, 3, 1},
	//	{0, 3, 1,    0, 0, 8,    7, 9, 0},
	//	{4, 5, 0,    7, 1, 3,    6, 0, 2},
	//}
	//grid := [9][9]int8{
	//	{0, 0, 0,    0, 8, 0,    0, 0, 4},
	//	{6, 4, 0,    0, 1, 0,    0, 0, 2},
	//	{0, 8, 0,    0, 6, 5,    0, 0, 0},

	//	{0, 0, 6,    0, 0, 0,    0, 1, 3},
	//	{5, 0, 0,    0, 0, 4,    2, 0, 0},
	//	{0, 3, 0,    9, 0, 0,    0, 0, 0},

	//	{0, 0, 0,    5, 7, 3,    9, 0, 0},
	//	{0, 0, 0,    0, 0, 6,    0, 3, 1},
	//	{0, 0, 9,    0, 4, 2,    0, 0, 0},
	//}
	//grid := [9][9]int8{
	//	{0, 0, 0,    3, 7, 0,    0, 0, 6},
	//	{0, 0, 0,    0, 0, 1,    7, 0, 0},
	//	{0, 0, 0,    0, 2, 5,    0, 0, 9},

	//	{9, 1, 0,    0, 6, 0,    0, 0, 0},
	//	{8, 5, 0,    7, 0, 0,    0, 4, 0},
	//	{0, 0, 0,    0, 0, 8,    0, 9, 2},

	//	{0, 4, 0,    0, 3, 0,    2, 0, 0},
	//	{0, 6, 0,    0, 4, 0,    0, 0, 1},
	//	{0, 8, 0,    0, 9, 0,    3, 5, 0},
	//}
	//grid := [9][9]int8{
	//	{8, 2, 0,    0, 4, 0,    0, 0, 0},
	//	{0, 0, 3,    5, 0, 0,    0, 0, 7},
	//	{0, 1, 6,    0, 0, 0,    0, 4, 0},

	//	{2, 0, 0,    0, 0, 4,    1, 6, 0},
	//	{0, 3, 0,    0, 0, 0,    0, 8, 0},
	//	{0, 0, 5,    2, 0, 0,    0, 0, 0},

	//	{0, 7, 0,    0, 9, 0,    6, 0, 0},
	//	{0, 0, 0,    0, 0, 1,    0, 0, 0},
	//	{0, 5, 0,    4, 0, 0,    0, 0, 2},
	//}
	//grid := [9][9]int8{
	//	{5, 3, 0,    0, 0, 0,    1, 0, 0},
	//	{0, 0, 0,    0, 3, 0,    0, 0, 2},
	//	{0, 0, 0,    7, 0, 6,    0, 0, 0},
          
	//	{7, 0, 0,    2, 0, 0,    9, 1, 0},
	//	{0, 0, 0,    1, 0, 0,    0, 0, 6},
	//	{1, 0, 0,    0, 8, 9,    0, 5, 0},
          
	//	{0, 0, 0,    6, 0, 0,    2, 0, 7},
	//	{2, 9, 0,    0, 0, 0,    0, 0, 8},
	//	{0, 0, 0,    0, 5, 0,    6, 0, 1},
	//}
	//grid := [9][9]int8{
	//	{0, 0, 0,    0, 0, 0,    0, 0, 0},
	//	{0, 0, 0,    9, 1, 5,    8, 3, 0},
	//	{8, 0, 0,    0, 0, 0,    0, 6, 0},
          
	//	{0, 0, 6,    2, 8, 4,    0, 9, 0},
	//	{0, 0, 0,    0, 0, 3,    0, 4, 0},
	//	{0, 8, 7,    6, 0, 9,    0, 2, 0},
          
	//	{0, 0, 0,    5, 0, 6,    0, 0, 0},
	//	{3, 2, 0,    4, 0, 0,    0, 0, 0},
	//	{0, 1, 0,    0, 0, 0,    3, 0, 0},
	//}

	//grid := [9][9]int8{
	//	{8, 0, 0,    0, 0, 7,    0, 9, 0},
	//	{0, 0, 7,    2, 0, 0,    0, 0, 0},
	//	{0, 0, 1,    0, 6, 0,    0, 0, 0},
          
	//	{0, 0, 3,    4, 0, 5,    9, 0, 0},
	//	{0, 0, 0,    9, 0, 0,    0, 8, 0},
	//	{0, 1, 2,    0, 0, 0,    0, 0, 5},
          
	//	{0, 4, 0,    0, 0, 6,    0, 2, 0},
	//	{0, 0, 6,    0, 0, 9,    5, 0, 0},
	//	{1, 2, 0,    0, 0, 0,    7, 0, 0},
	//}

	//grid := [9][9]int8{
	//	{0, 0, 0,    0, 0, 0,    0, 0, 0},
	//	{0, 0, 0,    0, 0, 0,    0, 0, 0},
	//	{0, 0, 0,    0, 0, 0,    0, 0, 0},

	//	{0, 0, 0,    0, 0, 0,    0, 0, 0},
	//	{0, 0, 0,    0, 0, 0,    0, 0, 0},
	//	{0, 0, 0,    0, 0, 0,    0, 0, 0},

	//	{0, 0, 0,    0, 0, 0,    0, 0, 0},
	//	{0, 0, 0,    0, 0, 0,    0, 0, 0},
	//	{0, 0, 0,    0, 0, 0,    0, 0, 0},
	//}

	grid := [9][9]int8{
		{5, 0, 0,    0, 6, 0,    0, 0, 8},
		{1, 6, 0,    0, 9, 0,    0, 4, 0},
		{0, 0, 8,    0, 0, 0,    1, 0, 7},

		{0, 5, 0,    2, 0, 0,    0, 0, 0},
		{0, 7, 0,    0, 0, 0,    0, 0, 1},
		{9, 0, 3,    4, 5, 0,    0, 0, 0},

		{0, 0, 0,    0, 0, 0,    0, 3, 9},
		{0, 0, 0,    8, 3, 0,    0, 0, 5},
		{0, 0, 9,    0, 2, 0,    7, 0, 0},
	}
	success := humansolver.SimpleLoop(grid[:])
	if !success {
		sudoku := fastsolver.InitSudoku(grid[:])
		fastsolver.Rule012Loop(sudoku[:])
		ui.PrintSudoku(sudoku[:])
	}
}
