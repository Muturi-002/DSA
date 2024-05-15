package main
import "fmt"
func main() {
	s:=[]int{33,42,12,99,372,92,65,8}
	merge_sort(s)
	fmt.Println("Newly sorted slice: ", s)
}
func merge_sort(slice1 []int){
	length:=len(slice1)
	merge_sort_recursion(slice1,0,length-1)
}
func merge_sort_recursion(slice []int,l,r int){
	//var 'l' defines the first index of the slice; var 'r' defines the last index of the slice
	if l<r{
		midpoint:= (l +(r-l))/2
		merge_sort_recursion(slice,l,midpoint)
		merge_sort_recursion(slice,midpoint+1,r)
		merge_sorted_arrays(slice,l,midpoint,r)
	}
}
func merge_sorted_arrays(slice []int,l,m,r int){
	//var 'm' is a placeholder for the midpoint
	length_left:=len(slice[:m])
	length_right:=len(slice[m:])
	left_slice:=slice[length_left]
	right_slice:=slice[length_right]
	var i,j,k int
	for i=0;i<length_left;i++{
		left_slice= slice[l+i]
	}
	for i=0;i<length_right;i++{
		right_slice= slice[m+1+i]
	}
	for k=l;k<=r;k++{
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
	}



}

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
func quick_sort(slice []int){

}