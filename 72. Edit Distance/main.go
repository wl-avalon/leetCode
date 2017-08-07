package main

import (
	"fmt"
	//"strings"
	"math"
)

func main() {
	testString 	:= "ultramicroscopically"
	testString1	:= "pneumonoultramicroscopicsilicovolcanoconiosis"
	fmt.Println(minDistance(testString, testString1))
}

func minDistance(word1 string, word2 string) int {
	matrix := make([][]int, len(word1) + 1)
	for i := 0; i <= len(word1); i++ {
		matrix[i] = make([]int, len(word2) + 1)
		for j := 0; j <= len(word2); j++ {
			matrix[0][j] = j
		}
		matrix[i][0] = i
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i - 1] == word2[j - 1] {
				matrix[i][j] = matrix[i - 1][j - 1]
			}else {
				min := math.Min(float64(matrix[i - 1][j]), float64(matrix[i][j - 1]))
				min = math.Min(float64(min), float64(matrix[i - 1][j - 1]))
				matrix[i][j] = int(min) + 1
			}
		}
 	}
	return matrix[len(word1)][len(word2)]
}