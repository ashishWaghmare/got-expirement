package main

import "fmt"

func main() {
	preferences := [][]int{{7, 5, 6, 4},
		{5, 4, 6, 7},
		{4, 5, 6, 7},
		{4, 5, 6, 7},
		{0, 1, 2, 3},
		{0, 1, 2, 3},
		{0, 1, 2, 3},
		{0, 1, 2, 3},
	}
	stable(preferences)
}
func stable([][]int) {
	womensPartner := make([]int, 4)
	menFree := make([]bool, 4)
	for i := 0; i < len(womensPartner); i++ {
		womensPartner[i] = -1
	}
	var freeCount = 4
	for freeCount > 0 {
		for freeMenIndex := 0; freeMenIndex < count; freeMenIndex++ {
			if menFree[freeMenIndex] == false {
				break
			}
		}
	}
	fmt.Println("Mens Free", menFree)
	fmt.Println("Womens found", womensPartner)
}
