package main

import "fmt"

// Function to convert a slice to a set (map)
func sliceToSet(slice []int) map[int]bool {
	set := make(map[int]bool)
	for _, item := range slice {
		set[item] = true
	}
	return set
}

func main() {
	slice := []int{1, 2, 3, 2, 4, 5, 4}
	set := sliceToSet(slice)

	// Printing the elements of the set
	fmt.Println("Set:")
	sliceSet := make([]int, 0)
	for item := range set {
		fmt.Println(item)
		sliceSet = append(sliceSet, item)
	}
	fmt.Printf("%T\n", set)
	
	fmt.Println(sliceSet)
}