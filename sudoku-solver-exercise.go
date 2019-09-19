package main


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
					} else {
						puzzle[nRow][nCol] = 0
					}
				}
				return false
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
	return !(!validateRows(puzzle) || !validateColumns(puzzle) || !validateSquares(puzzle))
}

func validateRows(puzzle [][]int) bool {
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

func validateColumns(puzzle [][]int) bool {
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

func main() {

}