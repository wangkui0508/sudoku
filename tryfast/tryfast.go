package main

import (
	"fmt"

	"github.com/wangkui0508/sudoku/generator"
	"github.com/wangkui0508/sudoku/fastsolver"
	"github.com/wangkui0508/sudoku/ui"
)

func main() {
	test()
}

func try() {
	totalTry := 0
	totalSuccess := 0
	for {
		values, _, difficulty := generator.GenSudokuPuzzle()
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
		sudoku := fastsolver.InitSudoku(values)
		success := fastsolver.Rule012Loop(sudoku[:])
		fmt.Printf("%d %v\n", totalTry, success)
		if success {
			totalSuccess++
		}
		//if !success {
		//	ui.PrintSudoku(sudoku)
		//	fmt.Println("============Input===================")
		//	sudoku = fastsolver.InitSudoku(values)
		//	ui.PrintSudoku(sudoku)
		//	for i := 0; i < 9; i++ {
		//		row := solution[i]
		//		fmt.Printf("%d%d%d %d%d%d %d%d%d\n",
		//		row[0],row[1],row[2],row[3],row[4],row[5],row[6],row[7],row[8])
		//		if i==2 || i == 5 {
		//			fmt.Println()
		//		}
		//	}
		//	break
		//}
	}
	fmt.Printf("totalSuccess %d\n", totalSuccess)
}

func test() {
	//values := [9][9]int{
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
	//values := [9][9]int{
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
	//values := [9][9]int{
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
	//values := [9][9]int{
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
	//values := [9][9]int{
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
	values := [9][9]int{
		{5, 3, 7,    0, 2, 0,    1, 6, 9},
		{6, 0, 0,    0, 3, 0,    4, 7, 2},
		{0, 2, 0,    7, 0, 6,    3, 8, 5},

		{7, 0, 0,    2, 6, 0,    9, 1, 0},
		{0, 0, 0,    1, 0, 0,    8, 2, 6},
		{1, 6, 2,    0, 8, 9,    7, 5, 0},

		{0, 0, 0,    6, 0, 0,    2, 0, 7},
		{2, 9, 6,    0, 0, 0,    5, 0, 8},
		{0, 7, 0,    0, 5, 2,    6, 0, 1},
	}
	sudoku := fastsolver.InitSudoku(values)

	fastsolver.Rule012Loop(sudoku[:])
	ui.PrintSudoku(sudoku[:])

	//fmt.Println("=================Rule3Round==================")
	//fastsolver.Rule3Round(sudoku[:])
	//ui.PrintSudoku(sudoku)

	//fmt.Println("===================================")
	//fastsolver.Rule012Loop(sudoku[:])
	//ui.PrintSudoku(sudoku)
	//fmt.Println("=================Rule3Round==================")
	//fastsolver.Rule3Round(sudoku[:])
	//ui.PrintSudoku(sudoku)

	//fmt.Println("=================Rule0Loop==================")
	//fastsolver.Rule0Loop(sudoku[:])
	//ui.PrintSudoku(sudoku)
	//fmt.Println("=================Rule1Round==================")
	//fastsolver.Rule1Round(sudoku[:])
	//ui.PrintSudoku(sudoku)
	//fmt.Println("=================Rule2Round==================")
	//fastsolver.Rule2Round(sudoku[:])
	//ui.PrintSudoku(sudoku)
}

