package main

import "fmt"

func bubbleSort(unsorted []int) (sorted []int) {
	sorted = unsorted
	for i := range sorted {
		for j := range sorted {
			iv := sorted[i]
			jv := sorted[j]
			if iv < jv {
				t := iv
				sorted[i] = jv
				sorted[j] = t
			}
		}
	}
	return
}

func main() {
	somelist := []int{7, 10, 1, 5, 2, 4, 3, 6}
	bubbleSort(somelist)
	fmt.Println(somelist)
}
