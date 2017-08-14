package main

import "fmt"

func main() {
	nums := []int{1,2,0,1,0,0,1,2,2,1}
	fmt.Printf("%v", sortColors(nums))
}

func sortColors(nums []int) []int {
	left, right := 0, len(nums) - 1
	for i, j := 0, 0; i <= right; {
		switch nums[i] {
			case 0:{
				temp := nums[left]
				nums[left] = nums[i]
				nums[i] = temp
				left++
				i++
			}
			case 1:{
				i++
			}
			case 2:{
				temp := nums[right]
				nums[right] = nums[i]
				nums[i] = temp
				right--
			}
		}
		j++
	}
	return nums
}