package humansolver

import (
	"fmt"
)

type Coord struct {
	Y int
	X int
}

var CList = []Coord{{0,0}, {0,1}, {0,2}, {1,0}, {1,1}, {1,2}, {2,0}, {2,1}, {2,2}}

func GetNonZeros(all [10]int8) []int8 {
	res := make([]int8, 0, 9)
	for _, num := range all {
		if num != 0 {
			res = append(res, num)
		}
	}
	return res
}

func AllDone(grid [][9]int8) bool {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if grid[y][x] == 0 {
				return false
			}
		}
	}
	return true
}

func GetUnfilledNumInBlock(grid [][9]int8, bY, bX int) []int8 {
	all := [10]int8{0,1,2,3,4,5,6,7,8,9}
	for y := bY*3; y < bY*3+3; y++ {
		for x := bX*3; x < bX*3+3; x++ {
			all[grid[y][x]] = 0
		}
	}
	return GetNonZeros(all)
}

func IsForbidden(num int8, bY, bX int, yPos, xPos int, colForbid, rowForbid [][9][3]int8) bool {
	for y := 0; y < 9; y++ {
		if bY*3 <= y || y < bY*3+3 {
			continue
		}
		if rowForbid[y][xPos][0] == num || rowForbid[y][xPos][1] == num || rowForbid[y][xPos][2] == num {
			return true
		}
	}
	for x := 0; x < 9; x++ {
		if bX*3 <= x || x < bX*3+3 {
			continue
		}
		if rowForbid[yPos][x][0] == num || rowForbid[yPos][x][1] == num || rowForbid[yPos][x][2] == num {
			return true
		}
	}
	return false
}

func GetPossibleCoordsForNumInBlock(grid [][9]int8, bY, bX int, num int8, colForbid, rowForbid [][9][3]int8) (coordList []Coord) {
	for y := bY*3; y < bY*3+3; y++ {
		for x := bX*3; x < bX*3+3; x++ {
			if grid[y][x] != 0 || IsForbidden(num, bY, bX, y, x, colForbid, rowForbid) ||
				HasNumInCol(grid, num, x) || HasNumInRow(grid, num, y) {
				continue
			}
			coordList = append(coordList, Coord{Y: y, X: x})
		}
	}
	return
}

func InSameCol(coordList []Coord) bool {
	c0 := coordList[0]
	for _, c := range coordList[1:] {
		if c0.X != c.X {
			return false
		}
	}
	return true
}

func InSameRow(coordList []Coord) bool {
	c0 := coordList[0]
	for _, c := range coordList[1:] {
		if c0.Y != c.Y {
			return false
		}
	}
	return true
}

func ExtractNewForbids(grid [][9]int8, colForbid, rowForbid [][9][3]int8) (newColForbid, newRowForbid [][9][3]int8) {
	newColForbid, newRowForbid = make([][9][3]int8, 9), make([][9][3]int8, 9)
	for _, blk := range CList {
		numbers := GetUnfilledNumInBlock(grid, blk.Y, blk.X)
		for _, num := range numbers {
			coordList := GetPossibleCoordsForNumInBlock(grid, blk.Y, blk.X, num, colForbid, rowForbid)
			if len(coordList) <= 1 {
				continue
			}
			if InSameCol(coordList) {
				for _, c := range coordList {
					if newColForbid[c.Y][c.X][0] == 0 {
						newColForbid[c.Y][c.X][0] = num
					} else if newColForbid[c.Y][c.X][1] == 0 {
						newColForbid[c.Y][c.X][1] = num
					} else if newColForbid[c.Y][c.X][2] == 0 {
						newColForbid[c.Y][c.X][2] = num
					}
				}
			}
			if InSameRow(coordList) {
				for _, c := range coordList {
					if newRowForbid[c.Y][c.X][0] == 0 {
						newRowForbid[c.Y][c.X][0] = num
					} else if newRowForbid[c.Y][c.X][1] == 0 {
						newRowForbid[c.Y][c.X][1] = num
					} else if newRowForbid[c.Y][c.X][2] == 0 {
						newRowForbid[c.Y][c.X][2] = num
					}
				}
			}
		}
	}
	return
}

func SetUnfilledNumInBlock(grid [][9]int8, num int8, bY, bX int) Coord {
	for y := bY*3; y < bY*3+3; y++ {
		for x := bX*3; x < bX*3+3; x++ {
			if grid[y][x] == 0 {
				grid[y][x] = num
				return Coord{Y: y, X: x}
			}
		}
	}
	return Coord{}
}

func GetUnfilledNumInRow(grid [][9]int8, y int) []int8 {
	all := [10]int8{0,1,2,3,4,5,6,7,8,9}
	for x := 0; x < 9; x++ {
		all[grid[y][x]] = 0
	}
	return GetNonZeros(all)
}

func SetUnfilledNumInRow(grid [][9]int8, num int8, y int) Coord {
	for x := 0; x < 9; x++ {
		if grid[y][x] == 0 {
			grid[y][x] = num
			return Coord{Y: y, X: x}
		}
	}
	return Coord{}
}

func GetUnfilledNumInCol(grid [][9]int8, x int) []int8 {
	all := [10]int8{0,1,2,3,4,5,6,7,8,9}
	for y := 0; y < 9; y++ {
		all[grid[y][x]] = 0
	}
	return GetNonZeros(all)
}

func SetUnfilledNumInCol(grid [][9]int8, num int8, x int) Coord {
	for y := 0; y < 9; y++ {
		if grid[y][x] == 0 {
			grid[y][x] = num
			return Coord{Y: y, X: x}
		}
	}
	return Coord{}
}

func HasNumInRow(grid [][9]int8, num int8, y int) bool {
	for x := 0; x < 9; x++ {
		if grid[y][x] == num {
			return true
		}
	}
	return false
}

func HasNumInCol(grid [][9]int8, num int8, x int) bool {
	for y := 0; y < 9; y++ {
		if grid[y][x] == num {
			return true
		}
	}
	return false
}

func FillOnlyRemainedNumber(grid [][9]int8) (coords4Blk, coords4Row, coords4Col []Coord) {
	for _, c := range CList {
		unfilledNumbers := GetUnfilledNumInBlock(grid, c.Y, c.X)
		if len(unfilledNumbers) == 1 {
			coord := SetUnfilledNumInBlock(grid, unfilledNumbers[0], c.Y, c.X)
			coords4Blk = append(coords4Blk, coord)
		}
	}
	for y := 0; y < 9; y++ {
		unfilledNumbers := GetUnfilledNumInRow(grid, y)
		if len(unfilledNumbers) == 1 {
			coord := SetUnfilledNumInRow(grid, unfilledNumbers[0], y)
			coords4Row = append(coords4Row, coord)
		}
	}
	for x := 0; x < 9; x++ {
		unfilledNumbers := GetUnfilledNumInCol(grid, x)
		if len(unfilledNumbers) == 1 {
			coord := SetUnfilledNumInCol(grid, unfilledNumbers[0], x)
			coords4Col = append(coords4Col, coord)
		}
	}
	return
}

func FillNumberWithOnlyPosition(grid [][9]int8, colForbid, rowForbid [][9][3]int8) (coordList []Coord) {
	for _, c := range CList {
		unfilledNumbers := GetUnfilledNumInBlock(grid, c.Y, c.X)
		for _, num := range unfilledNumbers {
			coords := GetPossibleCoordsForNumInBlock(grid, c.Y, c.X, num, colForbid, rowForbid)
			if len(coords) != 1 {
				continue
			}
			coord := coords[0]
			coordList = append(coordList, coord)
			grid[coord.Y][coord.X] = num
		}
	}
	return
}

func GetPossibleNumAt(grid [][9]int8, yPos, xPos int) []int8 {
	all := [10]int8{0,1,2,3,4,5,6,7,8,9}

	bY := yPos/3
	bX := xPos/3
	for y := bY*3; y < bY*3+3; y++ {
		for x := bX*3; x < bX*3+3; x++ {
			all[grid[y][x]] = 0
		}
	}
	for y := 0; y < 9; y++ {
		all[grid[y][xPos]] = 0
	}
	for x := 0; x < 9; x++ {
		all[grid[yPos][x]] = 0
	}
	return GetNonZeros(all)
}

func FillOnlyPossibleNumber(grid [][9]int8) (coordList []Coord) {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if grid[y][x] != 0 {
				continue
			}
			numbers := GetPossibleNumAt(grid, y, x)
			if len(numbers) != 1 {
				continue
			}
			grid[y][x] = numbers[0]
			coordList = append(coordList, Coord{Y: y, X: x})
		}
	}
	return coordList
}

func InitStrGrid(grid [][9]int8) [][9]string {
	var res [9][9]string
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if grid[y][x] == 0 {
				res[y][x] = " . "
			} else {
				res[y][x] = fmt.Sprintf(" %d ", grid[y][x])
			}
		}
	}
	return res[:]
}

func MarkPositions(strGrid [][9]string, coords []Coord) {
	for _, c := range coords {
		s := strGrid[c.Y][c.X]
		strGrid[c.Y][c.X] = "["+string(s[1])+"]"
	}
}

func PrintStrGrid(strGrid [][9]string) {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			ending := ""
			if x == 2 || x == 5 {
				ending = "|"
			}
			fmt.Printf("%s%s", strGrid[y][x], ending)
		}
		fmt.Println()
		if y == 2 || y == 5 {
			fmt.Println("---------+---------+---------")
		}
	}
}

func toByte(i int8) byte {
	if i == 0 {
		return byte('.')
	}
	return byte(i+'0')
}

func PrintForbidGrid(grid [][9][3]int8) {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			ending := ""
			if x == 2 || x == 5 {
				ending = "|"
			}
			var s [3]byte
			s[0] = toByte(grid[y][x][0])
			s[1] = toByte(grid[y][x][1])
			s[2] = toByte(grid[y][x][2])
			fmt.Printf("%s %s", string(s[:]), ending)
		}
		fmt.Println()
		if y == 2 || y == 5 {
			fmt.Println("------------+------------+------------")
		}
	}
}


func SimpleLoop(grid [][9]int8) bool {
	colForbid, rowForbid := make([][9][3]int8, 9), make([][9][3]int8, 9)
	hasProgress := SimplePass(grid, colForbid, rowForbid)
	for hasProgress {
		if AllDone(grid) {
			return true
		}
		colForbid, rowForbid = ExtractNewForbids(grid, colForbid, rowForbid)
		hasProgress = SimplePass(grid, colForbid, rowForbid)
	}
	return false
}

func SimplePass(grid [][9]int8, colForbid, rowForbid [][9][3]int8) bool {
	fmt.Println("%%%%%%%%%%%% New Pass %%%%%%%%%%%%%%%%%%%")
	PrintStrGrid(InitStrGrid(grid[:]))
	fmt.Println("^^^^^^^^^^^ Input ^^^^^^^^^^^^^^^^")
	PrintForbidGrid(colForbid[:])
	fmt.Println("^^^^^^^^^^^ Column Region for helping ^^^^^^^^^^^^^^^^")
	PrintForbidGrid(rowForbid[:])
	fmt.Println("^^^^^^^^^^^ Row Region for helping ^^^^^^^^^^^^^^^^")
	hasProgress := false
	coords4Blk, coords4Row, coords4Col := FillOnlyRemainedNumber(grid)
	if len(coords4Blk) + len(coords4Row) + len(coords4Col) != 0 {
		hasProgress = true
		strGrid := InitStrGrid(grid)
		MarkPositions(strGrid, coords4Blk)
		MarkPositions(strGrid, coords4Row)
		MarkPositions(strGrid, coords4Col)
		PrintStrGrid(strGrid)
		for _, c := range coords4Blk {
			fmt.Printf("Fill block's only missing number at R%d,C%d with '%d'\n", c.Y, c.X, grid[c.Y][c.X])
		}
		for _, c := range coords4Row {
			fmt.Printf("Fill row's only missing number at R%d,C%d with '%d'\n", c.Y, c.X, grid[c.Y][c.X])
		}
		for _, c := range coords4Col {
			fmt.Printf("Fill col's only missing number at R%d,C%d with '%d'\n", c.Y, c.X, grid[c.Y][c.X])
		}
	}
	coordList := FillNumberWithOnlyPosition(grid, colForbid, rowForbid)
	if len(coordList) != 0 {
		hasProgress = true
		strGrid := InitStrGrid(grid)
		MarkPositions(strGrid, coordList)
		PrintStrGrid(strGrid)
		for _, c := range coordList {
			fmt.Printf("In the block, only position for '%d' is R%d,C%d\n", grid[c.Y][c.X], c.Y, c.X)
		}
	}
	coordList = FillOnlyPossibleNumber(grid)
	if len(coordList) != 0 {
		hasProgress = true
		strGrid := InitStrGrid(grid)
		MarkPositions(strGrid, coordList)
		PrintStrGrid(strGrid)
		for _, c := range coordList {
			fmt.Printf("The only left number for R%d,C%d is '%d'\n", c.Y, c.X, grid[c.Y][c.X])
		}
	}
	return hasProgress
}
