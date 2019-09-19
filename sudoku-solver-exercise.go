package main

import "fmt"

// Sudoku Solver Exercise
// ======================
//
// The aim is to complete the solve() function so that it returns a completed
// sudoku puzzle. We've setup the execise to use an array to represent the
// state of a puzzle but if you think there is a more suitable data structure
// then feel free to change it.
//
// Please do not spend more than two hours on this, this can be a difficult
// exercise so it's not expected that you complete it within the given time. We
// want to see how you approach coding problems, so please show your working as
// best you can - bonus points if you use git to record your working history!

func solve(puzzle [][]int, completePuzzle [][]int) bool {
	if compare(puzzle, completePuzzle) {
		return true
	}
	for nRow := 0; nRow < 9; nRow++ {
		for nCol := 0; nCol < 9; nCol++ {
			num := puzzle[nRow][nCol]
			if num == 0 {
				for newNum := 1; newNum <= 9; newNum++ {
					puzzle[nRow][nCol] = newNum
					if validate(puzzle) {
						if solve(puzzle, completePuzzle) {
							return true
						}
						puzzle[nRow][nCol] = 0
					}
				}
			}
		}
	}
	return false
}

func compare(puzzle [][]int, completePuzzle [][]int) bool {
	for nRow := 0; nRow < 9; nRow++ {
		for nCol := 0; nCol < 9; nCol++ {
			if completePuzzle[nRow][nCol] != puzzle[nRow][nCol] {
				return false
			}
		}
	}
	return true
}

func validate(puzzle [][]int) bool {
	return !(!validatenRows(puzzle) || !validatenColumns(puzzle) || !validateSquares(puzzle))
}

func validatenRows(puzzle [][]int) bool {
	for nRow := 0; nRow < 9; nRow++ {
		counter := make([]int, 10)
		for _, num := range puzzle[nRow] {
			if num != 0 {
				counter[num]++
			}
		}
		if hasDuplicates(counter) {
			return false
		}
	}
	return true
}

func validatenColumns(puzzle [][]int) bool {
	for nCol := 0; nCol < 9; nCol++ {
		counter := make([]int, 10)
		for nRow := 0; nRow <9; nRow++ {
			num := puzzle[nRow][nCol]
			if num != 0 {
				counter[num]++
			}
		}
		if hasDuplicates(counter) {
			return false
		}
	}
	return true
}

func validateSquares(puzzle [][]int) bool {
	for nRow := 0; nRow < 9; nRow += 3 {
		for nCol := 0; nCol < 9; nCol += 3 {
			counter := make([]int, 10)
			for nRowSq := nRow; nRowSq < nRow + 3; nRowSq++ {
				for nColSq := nCol; nColSq < nCol + 3; nColSq++ {
					num := puzzle[nRowSq][nColSq]
					if num != 0 {
						counter[num]++
					}
				}
			}
			if hasDuplicates(counter) {
				return false
			}
		}
	}
	return true
}

func hasDuplicates(numbers []int) bool {
	for _, num := range numbers {
		if num > 1 {
			return true
		}
	}
	return false
}

var easySudoku = [][]int{
	{0, 0, 3, 0, 2, 0, 6, 0, 0},
	{9, 0, 0, 3, 0, 5, 0, 0, 1},
	{0, 0, 1, 8, 0, 6, 4, 0, 0},
	{0, 0, 8, 1, 0, 2, 9, 0, 0},
	{7, 0, 0, 0, 0, 0, 0, 0, 8},
	{0, 0, 6, 7, 0, 8, 2, 0, 0},
	{0, 0, 2, 6, 0, 9, 5, 0, 0},
	{8, 0, 0, 2, 0, 3, 0, 0, 9},
	{0, 0, 5, 0, 1, 0, 3, 0, 0},
}

var easySolution = [][]int{
	{4, 8, 3, 9, 2, 1, 6, 5, 7},
	{9, 6, 7, 3, 4, 5, 8, 2, 1},
	{2, 5, 1, 8, 7, 6, 4, 9, 3},
	{5, 4, 8, 1, 3, 2, 9, 7, 6},
	{7, 2, 9, 5, 6, 4, 1, 3, 8},
	{1, 3, 6, 7, 9, 8, 2, 4, 5},
	{3, 7, 2, 6, 8, 9, 5, 1, 4},
	{8, 1, 4, 2, 5, 3, 7, 6, 9},
	{6, 9, 5, 4, 1, 7, 3, 8, 2},
}

var hardSudoku = [][]int{
	{6, 0, 0, 0, 0, 0, 1, 5, 0},
	{9, 5, 4, 7, 1, 0, 0, 8, 0},
	{0, 0, 0, 5, 0, 2, 6, 0, 0},
	{8, 0, 0, 0, 9, 4, 0, 0, 6},
	{0, 0, 3, 8, 0, 5, 4, 0, 0},
	{4, 0, 0, 3, 7, 0, 0, 0, 8},
	{0, 0, 6, 9, 0, 3, 0, 0, 0},
	{0, 2, 0, 0, 4, 7, 8, 9, 3},
	{0, 4, 9, 0, 0, 0, 0, 0, 5},
}

var hardSolution = [][]int{
	{6, 3, 2, 4, 8, 9, 1, 5, 7},
	{9, 5, 4, 7, 1, 6, 3, 8, 2},
	{1, 7, 8, 5, 3, 2, 6, 4, 9},
	{8, 1, 7, 2, 9, 4, 5, 3, 6},
	{2, 9, 3, 8, 6, 5, 4, 7, 1},
	{4, 6, 5, 3, 7, 1, 9, 2, 8},
	{7, 8, 6, 9, 5, 3, 2, 1, 4},
	{5, 2, 1, 6, 4, 7, 8, 9, 3},
	{3, 4, 9, 1, 2, 8, 7, 6, 5},
}

func main() {
	var easyComplete = solve(easySudoku, easySolution)

	var hardComplete = solve(hardSudoku, hardSolution)

	fmt.Printf("Easy Sudoku Solved?: %v\n", easyComplete)
	fmt.Println(easySudoku)
	fmt.Printf("Hard Sudoku Solved?: %v\n", hardComplete)
	fmt.Println(hardSudoku)
}
