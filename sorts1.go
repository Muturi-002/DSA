//This file contains code for bubble sort, insertion sort, and selection sort.
package main
import "fmt"
func main(){
	slice1:=[]int{}
	fmt.Println("Empty slice: ",slice1)
	slice1= append(slice1, 0,88,92,10,22,3,7,45)
	fmt.Println("slice1: ",slice1)
	fmt.Println("=======")
	bubble_sort(slice1)
	insertion_sort(slice1)
	selection_sort(slice1)
	fmt.Println("=======")
	array1:=[]int{1,5,83,99,2,3,1,95}
	fmt.Println("Array: ",array1)
	bubble_sort(array1)
	insertion_sort(array1)
	selection_sort(array1)
}
func bubble_sort(slice []int){
	fmt.Println("Bubble Sort ******")
	fmt.Println("Slice to be sorted: ",slice)
	n:=len(slice)
	for i:=0;i<n;i++{
		for j:=0;j<n-1;j++{
			if slice[j]>slice[j+1]{
				bigel:=slice[j]
				slice[j]=slice[j+1]
				slice[j+1]=bigel
			}
		}
	}
	fmt.Println("newly sorted array: ",slice)
}
func insertion_sort(slice []int){
	fmt.Println("Insertion sort ****")
	l:=len(slice)
	for k:=1;k<l;k++ {
		insert:=slice[k]
		j:=k-1
		if(j>=0){
			if(slice[j]>insert){
				slice[j+1]=slice[j]
				j-=1
				slice[j+1]=insert
			}
		}
	}
	fmt.Println("newly sorted array: ",slice)
}


func selection_sort(slice []int){
	fmt.Println("Selection sort ****")
	sorted_slice:=[]int{}
	for i :=0;i<len(slice);i++{
		min_element:=i
		for j:=1;j<len(slice);j++{
			if slice[i]>slice[j]{
				min_element=j
				sorted_slice= append(sorted_slice,min_element)
			}
		}
	}
	fmt.Println("newly sorted array: ",slice)
}