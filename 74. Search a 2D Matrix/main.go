package main

import (
	"fmt"
	//"math"
)

func main() {
	matrix 		:= [][]int{{-8,-7,-5,-3,-3,-1,1}, {2,2,2,3,3,5,7}, {8,9,11,11,13,15,17}, {18,18,18,20,20,20,21}, {23,24,26,26,26,27,27}, {28,29,29,30,32,32,34}}
	//matrix		:= [][]int{{1}, {3}}
	//matrix          := [][]int{{1, 1}}
	boolValue 	:=searchMatrix(matrix, -5)
	fmt.Printf("%v", boolValue)
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}else if len(matrix[0]) == 0 {
		return false
	}

	left ,right := 0, len(matrix) - 1
	for mid := 0; left <= right; {
		mid = (left + right) / 2
		if matrix[mid][0] == target {
			return true
		}else if matrix[mid][0] > target {
			right = mid - 1
		}else {
			left = mid + 1
		}
	}
	if left > len(matrix) - 1 {
		left = len(matrix) - 1
	}

	lineNum := 0
	if matrix[left][0] > target {
		if left == 0 {
			return false
		}
		lineNum = left - 1
	}else {
		lineNum = left
	}

	row := matrix[lineNum]
	left, right = 0, len(row) - 1
	for mid := 0; left <= right; {
		mid = (left + right) / 2
		if row[mid] == target {
			return true
		}else if row[mid] > target {
			right = mid - 1
		}else {
			left = mid + 1
		}
	}
	return false
}