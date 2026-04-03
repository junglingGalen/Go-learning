package main

import (
	"fmt"
)

func mysort(arr []float64) []float64{
	if len(arr) == 1 {
		return arr
	}
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - i - 1; j++ {
			if arr[j] > arr [j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	fmt.Println(mysort([]float64{5,6,2,1,4,7,8,9,3}))
}