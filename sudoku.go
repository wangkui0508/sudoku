package main

import (
	"fmt"
	"math/bits"
)

func IsFixed(n uint16) bool {
	return bits.OnesCount16(n) == 1
}

func GetResult(n uint16) int {
	return bits.TrailingZeros16(n)+1
}

func YX(y, x int) int {
	return (x-1) + (y-1)*9
}

func InitSudoku(values [9][9]int) (sudoku [9][9]uint16) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			sudoku[i][j] = 0x1FF
		}
	}
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			n := values[y][x]
			if n > 0 {
				sudoku[y][x] = uint16(1) << (n-1)
			}
		}
	}
	return
}

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

func FillChart(sudoku [9][9]uint16, chart [][BlkSz*3]byte) {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			FillBlock(sudoku, chart, y, x)
		}
	}
}

func FillBlock(sudoku [9][9]uint16, chart [][BlkSz*3]byte, yStart, xStart int) {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			n := sudoku[yStart*3+y][xStart*3+x]
			FillNum(n, chart, yStart*BlkSz + y*4, xStart*BlkSz + x*4)
		}
	}
}

func FillNum(n uint16, chart [][BlkSz*3]byte, yStart, xStart int) {
	if IsFixed(n) {
		r := GetResult(n)
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

func CountFixed(sudoku [][9]uint16) int {
	res := 0
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if IsFixed(sudoku[y][x]) {
				res++
			}
		}
	}
	return res
}

// ===========================================

func Rule0Loop(sudoku [][9]uint16) {
	totalFixed := CountFixed(sudoku)
	for {
		if totalFixed == 9*9 {
			break //finished
		}
		Rule0Round(sudoku)
		tf := CountFixed(sudoku)
		if totalFixed == tf {
			break //no progress
		}
		totalFixed = tf
	}
}

type Coord struct {
	y int
	x int
}

func Rule0Round(sudoku [][9]uint16) {
	var coordList []Coord
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			n := sudoku[y][x]
			if IsFixed(n) {
				coordList = append(coordList, Coord{y: y, x: x})
			}
		}
	}
	for _, c := range coordList {
		Rule0ForPos(sudoku, c.y, c.x)
	}
}

func Rule0ForPos(sudoku [][9]uint16, yPos, xPos int) {
	n := sudoku[yPos][xPos]
	// the same column
	for y := 0; y < 9; y++ {
		if y == yPos {
			continue
		}
		sudoku[y][xPos] &= ^n
	}
	// the same row
	for x := 0; x < 9; x++ {
		if x == xPos {
			continue
		}
		sudoku[yPos][x] &= ^n
	}
	// the same block
	xStart := (xPos/3)*3
	yStart := (yPos/3)*3
	for x := xStart; x < xStart+3; x++ {
		for y := yStart; y < yStart+3; y++ {
			if x == xPos && y == yPos {
				continue
			}
			sudoku[y][x] &= ^n
		}
	}
}

// ===========================================

type Rule1Pos struct {
	IsVertical bool
	Mask       uint16
	Where      int
	Source     int
}

var CList = []Coord{{0,0}, {0,1}, {0,2}, {1,0}, {1,1}, {1,2}, {2,0}, {2,1}, {2,2}}

func FindRule1InBlock(sudoku [][9]uint16, isVertical bool, mask uint16, y, x int) (res []Rule1Pos) {
	for off := 0; off < 3; off++ {
		foundIt := true
		for _, c := range CList {
			if (isVertical && c.x == off) ||
			   (!isVertical && c.y == off) {
				continue
			}
			if (sudoku[y*3+c.y][x*3+c.x] & mask) != 0 {
				foundIt = false
				break
			}
		}
		if !foundIt {
			continue
		}
		if isVertical {
			res = append(res, Rule1Pos{IsVertical: isVertical, Mask: mask, Where: x*3+off, Source: y})
		} else {
			res = append(res, Rule1Pos{IsVertical: isVertical, Mask: mask, Where: y*3+off, Source: x})
		}
	}
	return
}

func FindRule1(sudoku [][9]uint16) (res []Rule1Pos) {
	for i := 0; i < 9; i++ {
		mask := uint16(1) << i
		for _, c := range CList {
			r := FindRule1InBlock(sudoku, true, mask, c.y, c.x)
			res = append(res, r...)
			r = FindRule1InBlock(sudoku, false, mask, c.y, c.x)
			res = append(res, r...)
		}
	}
	return
}

func Rule1Round(sudoku [][9]uint16) {
	rule1posList := FindRule1(sudoku)
	for _, pos := range rule1posList {
		if pos.IsVertical {
			for y := 0; y < 9; y++ {
				if pos.Source*3 <= y && y < pos.Source*3 + 3 {
					continue
				}
				sudoku[y][pos.Where] &= ^pos.Mask
			}
		} else {
			for x := 0; x < 9; x++ {
				if pos.Source*3 <= x && x < pos.Source*3 + 3 {
					continue
				}
				sudoku[pos.Where][x] &= ^pos.Mask
			}
		}
	}
}

// ===========================================

func Rule2Round(sudoku [][9]uint16) {
	for n := 0; n < 9; n++ {
		Rule2ForMask(sudoku, uint16(1) << n)
	}
}

func Rule2ForMask(sudoku [][9]uint16, mask uint16) {
	// the same column
	for xPos := 0; xPos < 9; xPos++ {
		count := 0
		where := -1
		for y := 0; y < 9; y++ {
			if (sudoku[y][xPos] & mask) != 0 {
				count++
				where = y
			}
		}
		if count == 1 {
			sudoku[where][xPos] = mask
		}
	}
	// the same row
	for yPos := 0; yPos < 9; yPos++ {
		count := 0
		where := -1
		for x := 0; x < 9; x++ {
			if (sudoku[yPos][x] & mask) != 0 {
				count++
				where = x
			}
		}
		if count == 1 {
			sudoku[yPos][where] = mask
		}
	}
	// the same block
	for _, c := range CList {
		xStart := c.x*3
		yStart := c.x*3
		count := 0
		whereX, whereY := -1, -1
		for x := xStart; x < xStart+3; x++ {
			for y := yStart; y < yStart+3; y++ {
				if (sudoku[y][x] & mask) != 0 {
					count++
					whereX = x
					whereY = y
				}
			}
		}
		if count == 1 {
			sudoku[whereY][whereX] = mask
		}
	}
}

// ===========================================

func Rule012Loop(sudoku [][9]uint16) {
	totalFixed := CountFixed(sudoku)
	for {
		if totalFixed == 9*9 {
			break //finished
		}
		Rule0Round(sudoku)
		Rule1Round(sudoku)
		Rule2Round(sudoku)
		tf := CountFixed(sudoku)
		if totalFixed == tf {
			break //no progress
		}
		totalFixed = tf
	}
}

// ===========================================

func PrintSudoku(sudoku [9][9]uint16) {
	chart := EmptyChart()
	FillChart(sudoku, chart[:])
	PrintChart(chart[:])
}

func main() {
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
	values := [9][9]int{
		{0, 0, 0,    0, 8, 0,    0, 0, 4},
		{6, 4, 0,    0, 1, 0,    0, 0, 2},
		{0, 8, 0,    0, 6, 5,    0, 0, 0},

		{0, 0, 6,    0, 0, 0,    0, 1, 3},
		{5, 0, 0,    0, 0, 4,    2, 0, 0},
		{0, 3, 0,    9, 0, 0,    0, 0, 0},

		{0, 0, 0,    5, 7, 3,    9, 0, 0},
		{0, 0, 0,    0, 0, 6,    0, 3, 1},
		{0, 0, 9,    0, 4, 2,    0, 0, 0},
	}
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
	sudoku := InitSudoku(values)
	Rule012Loop(sudoku[:])
	PrintSudoku(sudoku)

	//fmt.Println("=================Rule0Loop==================")
	//Rule0Loop(sudoku[:])
	//PrintSudoku(sudoku)
	//fmt.Println("=================Rule1Round==================")
	//Rule1Round(sudoku[:])
	//PrintSudoku(sudoku)
	//fmt.Println("=================Rule2Round==================")
	//Rule2Round(sudoku[:])
	//PrintSudoku(sudoku)
}

