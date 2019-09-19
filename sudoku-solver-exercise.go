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

func solve(puzzle [][]int) [][]int {
	fmt.Println(validateSquares(puzzle))
	return puzzle
}

func compare(puzzle [][]int, completePuzzle [][]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if completePuzzle[row][col] != puzzle[row][col] {
				return false
			}
		}
	}
	return true
}

func validateRows(puzzle [][]int) bool {
	for row := 0; row < 9; row++ {
		counter := make([]int, 10)
		for _, num := range puzzle[row] {
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

func validateColumns(puzzle [][]int) bool {
	for col := 0; col < 9; col++ {
		counter := make([]int, 10)
		for row := 0; row <9; row++ {
			num := puzzle[row][col]
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
	for row := 0; row < 9; row += 3 {
		for col := 0; col < 9; col += 3 {
			counter := make([]int, 10)
			for rowSq := row; rowSq < row + 3; rowSq++ {
				for colSq := col; colSq < col + 3; colSq++ {
					num := puzzle[rowSq][colSq]
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
	var easyComplete = compare(solve(easySudoku), easySolution)

	var hardComplete = compare(solve(hardSudoku), hardSolution)

	fmt.Printf("Easy Sudoku Solved?: %v\n", easyComplete)
	fmt.Printf("Hard Sudoku Solved?: %v\n", hardComplete)
}
