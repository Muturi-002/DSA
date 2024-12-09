//This file contains code for both quicksort and mergesort algorithms.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quick_sort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	pivot := slice[0]
	var left_slice, right_slice []int

	for i := 1; i < len(slice); i++ {
		if slice[i] <= pivot {
			left_slice = append(left_slice, slice[i])
		} else {
			right_slice = append(right_slice, slice[i])
		}
	}
	return append(append(quick_sort(left_slice), pivot), quick_sort(right_slice)...)
}

func contains(slice []int, numbers int) bool { //serves both quick_sort and merge_sort functions
	for _, v := range slice {
		if v == numbers {
			return true
		}
	}
	return false
}
func mergesort(slice []int) []int{
	//l,r:= 0,len(slice)-1
	if len(slice)<=1{
		return slice
	}
	mid:= len(slice)/2
	left_slice:=mergesort(slice[:mid])
	right_slice:=mergesort(slice[mid:])
	
	return merge(left_slice,right_slice)
}
func merge(left,right []int) []int{
	left_index, right_index:= 0,0 //mark the first index of the prototype slices respectively
	newSlice:= make([]int, 0, len(left)+len(right))//create a new  sorted slice
	for left_index<len(left) && right_index<len(right){
		if left[left_index]<=right[right_index]{
			newSlice= append(newSlice, left[left_index])
			left_index++
		}else{
			newSlice= append(newSlice, right[right_index])
			right_index++
		}
	}
	newSlice = append(newSlice, left[left_index:]...)
	newSlice = append(newSlice, right[right_index:]...)
	return newSlice
}
func main() {
	var sizeArray, slice []int
	for j := 5; j <= 30; j += 5 {
		sizeArray = append(sizeArray, j)
	}
	var choice int
	fmt.Println("Sizes for test: ", sizeArray)
	fmt.Println("Choose a sorting algorithm:")
	fmt.Println("\t1. Quicksort algorithm\n\t2. Mergesort algorithm")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		{
			for _, sizeArray := range sizeArray {
				fmt.Println("Size of slice: ", sizeArray)
				for i := 0; i < sizeArray; i++ {
					randomNo := rand.Intn(1000)
					if !contains(slice, randomNo) {
						slice = append(slice, randomNo)
					}
				}
				fmt.Println("Slice: ", slice)
				t := time.Now()
				slice = quick_sort(slice)
				fmt.Println("Time taken: ", time.Since(t))
				fmt.Println("Newly sorted slice: ", slice)
				fmt.Println("==============")
			}
		}
		break
	case 2:
		for _, sizeArray := range sizeArray {
			fmt.Println("Size of slice: ", sizeArray)
			for i := 0; i < sizeArray; i++ {
				randomNo := rand.Intn(1000)
				if !contains(slice, randomNo) {
					slice = append(slice, randomNo)
				}
			}
			fmt.Println("Slice: ", slice)
			t := time.Now()
			slice = mergesort(slice)
			fmt.Println("Time taken: ", time.Since(t))
			fmt.Println("Newly sorted slice: ", slice)
			fmt.Println("==============")
		}
		break
	default:
		break
	}
}
