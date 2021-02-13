package ui

import (
	"fmt"

	"github.com/wangkui0508/sudoku/types"
)

type MaskType = types.MaskType

const (
	BlkSz = 12
)

func EmptyChart() (chart [BlkSz*3][BlkSz*3]byte) {
	for x := range chart {
		for y := range chart[x] {
			chart[y][x] = ' '
		}
	}
	for _, x := range []int{BlkSz-1, BlkSz*2-1} {
		for y := 0; y < BlkSz*3; y++ {
			chart[y][x] = '|'
		}
	}
	for _, y := range []int{BlkSz-1, BlkSz*2-1} {
		for x := 0; x < BlkSz*3; x++ {
			chart[y][x] = '-'
		}
	}
	return
}

func PrintChart(chart [][BlkSz*3]byte) {
	for _, line := range chart {
		fmt.Println(string(line[:]))
	}
}

func FillChart(sudoku [9][9]MaskType, chart [][BlkSz*3]byte) {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			FillBlock(sudoku, chart, y, x)
		}
	}
}

func FillBlock(sudoku [9][9]MaskType, chart [][BlkSz*3]byte, yStart, xStart int) {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			n := sudoku[yStart*3+y][xStart*3+x]
			FillNum(n, chart, yStart*BlkSz + y*4, xStart*BlkSz + x*4)
		}
	}
}

func FillNum(n MaskType, chart [][BlkSz*3]byte, yStart, xStart int) {
	if types.IsFixed(n) {
		r := types.GetResult(n)
		for i := 0; i < 9; i++ {
			y, x := i/3, i%3
			chart[yStart+y][xStart+x] = byte('#')
		}
		chart[yStart+1][xStart+1] = byte('0')+byte(r)
		return
	}
	for i := 0; i < 9; i++ {
		y, x := i/3, i%3
		if (n & (1<<i)) != 0 {
			chart[yStart+y][xStart+x] = byte('0')+byte(i+1)
		} else {
			chart[yStart+y][xStart+x] = byte('-')
		}
	}
}

func PrintSudoku(sudoku [9][9]MaskType) {
	chart := EmptyChart()
	FillChart(sudoku, chart[:])
	PrintChart(chart[:])
}

