package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{{1, 2, 0}, {2, 4, 1}, {0, 2, 4}, {5, 6, 7}}
	setZeroes(matrix)
	fmt.Printf("%v", matrix)
}

func setZeroes(matrix [][]int)  {
	posList := [][]int{}
	xlength := 0
	ylength := 0
	for m, mValue := range matrix {
		for n, nValue := range mValue {
			if(nValue == 0){
				posList = append(posList, []int{m, n})
			}
			xlength = len(mValue)
		}
		ylength = len(matrix)
	}
	fmt.Printf("%v", posList)
	fmt.Print("\n")
	for _, pos := range posList {
		x := pos[0]
		y := pos[1]
		for i := 0; i < xlength; i++ {
			matrix[x][i] = 0
		}
		for i := 0; i < ylength; i++ {
			matrix[i][y] = 0
		}
	}
}