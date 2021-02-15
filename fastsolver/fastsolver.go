package fastsolver

import (
	"fmt"

	"github.com/wangkui0508/sudoku/types"
	"github.com/wangkui0508/sudoku/ui"
)

type MaskType = types.MaskType

func InitSudoku(values [9][9]int) (sudoku [9][9]MaskType) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			sudoku[i][j] = 0x1FF
		}
	}
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			n := values[y][x]
			if n > 0 {
				sudoku[y][x] = MaskType(1) << (n-1)
			}
		}
	}
	return
}

func CountFixed(sudoku [][9]MaskType) int {
	res := 0
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if types.IsFixed(sudoku[y][x]) {
				res++
			}
		}
	}
	return res
}

func CountPossibleValueCount(sudoku [][9]MaskType) int {
	res := 0
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			res += types.PossibleValueCount(sudoku[y][x])
		}
	}
	return res
}

// ===========================================

func Rule0Loop(sudoku [][9]MaskType) {
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
	Y int
	X int
}

func Rule0Round(sudoku [][9]MaskType) {
	var coordList []Coord
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			n := sudoku[y][x]
			if types.IsFixed(n) {
				coordList = append(coordList, Coord{Y: y, X: x})
			}
		}
	}
	for _, c := range coordList {
		Rule0ForPos(sudoku, c.Y, c.X)
	}
}

func Rule0ForPos(sudoku [][9]MaskType, yPos, xPos int) {
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
	Mask       MaskType
	Where      int
	Source     int
}

var CList = []Coord{{0,0}, {0,1}, {0,2}, {1,0}, {1,1}, {1,2}, {2,0}, {2,1}, {2,2}}

func FindRule1InBlock(sudoku [][9]MaskType, isVertical bool, mask MaskType, y, x int) (res []Rule1Pos) {
	for off := 0; off < 3; off++ {
		foundIt := true
		for _, c := range CList {
			if (isVertical && c.X == off) ||
			   (!isVertical && c.Y == off) {
				continue
			}
			if (sudoku[y*3+c.Y][x*3+c.X] & mask) != 0 {
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

func FindRule1(sudoku [][9]MaskType) (res []Rule1Pos) {
	for i := 0; i < 9; i++ {
		mask := MaskType(1) << i
		for _, c := range CList {
			r := FindRule1InBlock(sudoku, true, mask, c.Y, c.X)
			res = append(res, r...)
			r = FindRule1InBlock(sudoku, false, mask, c.Y, c.X)
			res = append(res, r...)
		}
	}
	return
}

func Rule1Round(sudoku [][9]MaskType) {
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

// ================================================

func Rule2Round(sudoku [][9]MaskType) {
	for n := 0; n < 9; n++ {
		Rule2ForMask(sudoku, MaskType(1) << n, 1)
	}
	for n := 0; n < 9; n++ {
		for m := 0; m < n; m++ {
			Rule2ForMask(sudoku, (MaskType(1)<<n)|(MaskType(1)<<m), 2)
		}
	}
	for n := 0; n < 9; n++ {
		for m := 0; m < n; m++ {
			for l := 0; l < m; l++ {
				Rule2ForMask(sudoku, (MaskType(1)<<n)|(MaskType(1)<<m)|(MaskType(1)<<l), 3)
			}
		}
	}
}

func Rule2ForMask(sudoku [][9]MaskType, mask MaskType, totalCount int) {
	// the same column
	coordList := make([]Coord, 0, 9)
	patialMatch := 0
	analyzeList := func() {
		if len(coordList) == totalCount && patialMatch == 0 {
			for _, c := range coordList {
				sudoku[c.Y][c.X] = mask
			}
		}
		coordList = coordList[:0]
		patialMatch = 0
	}
	for xPos := 0; xPos < 9; xPos++ {
		for y := 0; y < 9; y++ {
			if (sudoku[y][xPos] & mask) == mask {
				coordList = append(coordList, Coord{Y: y, X: xPos})
			} else if (sudoku[y][xPos] & mask) != 0 {
				patialMatch++
			}
		}
		analyzeList()
	}
	// the same row
	for yPos := 0; yPos < 9; yPos++ {
		coordList = coordList[:0]
		for x := 0; x < 9; x++ {
			if (sudoku[yPos][x] & mask) == mask {
				coordList = append(coordList, Coord{Y: yPos, X: x})
			} else if (sudoku[yPos][x] & mask) != 0 {
				patialMatch++
			}
		}
		analyzeList()
	}
	// the same block
	for _, c := range CList {
		xStart := c.X*3
		yStart := c.Y*3
		for x := xStart; x < xStart+3; x++ {
			for y := yStart; y < yStart+3; y++ {
				if (sudoku[y][x] & mask) == mask {
					coordList = append(coordList, Coord{Y: y, X: x})
				} else if (sudoku[y][x] & mask) != 0 {
					patialMatch++
				}
			}
		}
		analyzeList()
	}
}

// ===========================================

func Rule012Loop(sudoku [][9]MaskType) (success bool) {
	totalPossible := CountPossibleValueCount(sudoku)
	for {
		if CountFixed(sudoku) == 9*9 {
			return true
		}
		Rule0Round(sudoku)
		fmt.Println("========== Rule0Round: ============")
		ui.PrintSudoku(sudoku)
		Rule1Round(sudoku)
		fmt.Println("========== Rule1Round: ============")
		ui.PrintSudoku(sudoku)
		Rule2Round(sudoku)
		fmt.Println("========== Rule2Round: ============")
		ui.PrintSudoku(sudoku)
		tp := CountPossibleValueCount(sudoku)
		if totalPossible == tp {
			return false
		}
		totalPossible = tp
	}
}

