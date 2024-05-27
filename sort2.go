package main
import "fmt"
func main() {
	s:=[]int{33,42,12,99,372,92,65,8}
	fmt.Println("Slice: ", s)
	//s= quickSort(s)...
	s=quick_sort(s)
	fmt.Println("Newly sorted slice: ",s)
}

/* func merge_sort(slice1 []int)  []int{
	length:=len(slice1)
	return merge_sort_recursion(slice1,0,length-1)
}
func merge_sort_recursion(slice []int,l,r int) []int{
	//var 'l' defines the first index of the slice; var 'r' defines the last index of the slice
	if l<r{
		midpoint:= (l +(r-l))/2
		merge_sort_recursion(slice,l,midpoint)
		merge_sort_recursion(slice,midpoint+1,r)
		return merge_sorted_arrays(slice,l,midpoint,r)
	}
}
func merge_sorted_arrays(slice []int,l,m,r int) []int{
	//var 'm' is a placeholder for the midpoint
	length_left:=len(slice[:m])
	length_right:=len(slice[m:])
	var left_slice []int
	var right_slice []int
	for p, _:=range slice[:length_left] {
		left_slice = append(left_slice,p)
	}
	for q, _:=range slice[length_right:]{
		right_slice = append(right_slice,q)
	}
	
	var i,j,k int
	for k=l;k<=r;k++{
		i=0
		j=0
		for i<length_left && j<length_right{
			if left_slice[i]<=right_slice[j]{
				slice[k]=left_slice[i]
				i++
			}else{
				slice[k]=right_slice[j]
				j++
			}
		}
		for i<length_left{
			slice[k]=left_slice[i]
			i++
			k++
		}
		for j<length_right{
			slice[k]=right_slice[j]
			j++
			k++
		}
	} */
	// fmt.Println(left_slice)
	// fmt.Println(right_slice)
	/*for k=l;k<=r;k++{
		for i=0{
			for j=0{
				if (i<length_left && j>=length_right) || left_slice[i]<=right_slice[j]{
					slice[k]=left_slice[i]
					i++
				}else{
					slice[k]=right_slice[j]
					j++
				}
			}
		}
	}*/
	// return slice
// }

/*func merge_sort(slice []int){
	fmt.Println("Merge sort====")
	fmt.Print("Slice to be sorted: ",slice,"\n")
	if len(slice)==1{
		return
	}
	k:=0
	midLength:=len(slice)/2
	left:=slice[:midLength]
	right:=slice[midLength:]
	merge_sort(left)
	merge_sort(right)
	left_index:= 0
	right_index:= 0
	for left_index< len(left) && right_index<len(right){
		if left[left_index]<right[right_index]{
			slice[k]=left[left_index]
			left_index++
		}else{
			slice[k]=right[right_index]
			right_index++
		}
		k++
	}
	fmt.Println("Newly sorted slice: ",slice)

}*/
func quick_sort(slice []int) []int{
	if len(slice)< 2{
		return slice
	}

	pivot:=slice[0]
	var left_slice, right_slice []int

	for i:=1;i<len(slice);i++{
		if slice[i]<=pivot{
			left_slice=append(left_slice,slice[i])
		}else{
			right_slice=append(right_slice,slice[i])
		}
	}
	return append(append(quick_sort(left_slice),pivot),quick_sort(right_slice)...)
}
func quickSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	pivot := slice[0]
	var leftSlice, rightSlice []int

	for i := 1; i < len(slice); i++ {
		if slice[i] <= pivot {
			leftSlice = append(leftSlice, slice[i])
		} else {
			rightSlice = append(rightSlice, slice[i])
		}
	}

	return append(append(quickSort(leftSlice), pivot), quickSort(rightSlice)...)
}