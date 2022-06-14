package main

import "fmt"

// the code below is used to create a method for doing quick sort

// the method below takes the slice and the start index and the end index as input
// and returns the sorted array as output
func quickSort(inputArr []int, startIndex int, endIndex int) []int {

	// the code below is used to check that if the startIndex is less then
	// endIndex then doing quickSort of the slice
	if startIndex < endIndex {
		// the code below is used to get the pivot element
		pivotElementIndex := partionAndPivotElementIndex(inputArr, startIndex, endIndex)

		// the code below is used to call the quickSort() function for first using the quicksort
		// for the elements on the left side of the pivot i.e. the elements whose values
		// are less then the value of the element at the pivot index

		// in thecode below we are passing the endIndex as the index of the element
		// that occurs just before the pivot element
		inputArr = quickSort(inputArr, startIndex, pivotElementIndex-1)

		// the code below is used to call the quickSort() function for  using the quicksort
		// for the elements on the right side of the pivot i.e. the elements whose values
		// are greater then the value of the element at the pivot index

		// in thecode below we are passing the startIndex as the index of the element
		// that occurs just after the pivot element
		inputArr = quickSort(inputArr, pivotElementIndex+1, endIndex)

	}
	return inputArr
}

// the code below is used to create a method to do the partition of the slice and
// return the pivot element

// the method below takes the slice and the start index and the end index as input
// and returns the index of pivot element as output
func partionAndPivotElementIndex(inputArr []int, startIndex int, endIndex int) int {
	// the code below is used to create a pivot element
	var pivotElement int = inputArr[endIndex] // here we are considering the last
	// element as the pivot element

	// creating a variable for getting the start index
	var i int = startIndex

	// the code below is to use the for loop for traversing the slice inputArr
	// from index i to the endIndex position
	for j := i; j < endIndex; j++ {
		// the code below is used to check that if the element at the jth position
		// is smaller than the pivot element then replacing the element at index
		// i with the element at index j
		if inputArr[j] < pivotElement {
			// the code below is used to swap the value of element at index i
			// with the value at index j
			inputArr[i], inputArr[j] = inputArr[j], inputArr[i]
			// the code below is used to increase the value of index i by 1
			i++
		}
	}
	// the code below is used to change the value at index i with the value at endIndex
	inputArr[i], inputArr[endIndex] = inputArr[endIndex], inputArr[i]

	// the code below is used to return the value of i
	return i
}

// the method below is used to do quick sort in go
func main() {
	// the code below is used to  create a dummy slice
	var arr = []int{20, 35, -15, 7, 55, 1, -22}

	// the code below is used to call the quickSort() method and then sort
	// the arr slice
	fmt.Println(quickSort(arr, 0, len(arr)-1))
}
