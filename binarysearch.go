package main
import "fmt"
func main() {
	s:= []int{92,38,77,20,33,22,46,63}
	var search int 
	fmt.Print("Enter number for search: ")
	fmt.Scan(&search)
	fmt.Printf("Slice: %v\n", s)
	result:=binarySearch(s,search)
	fmt.Println(result)

}
func binarySearch(slice []int, value int) int {
	for i:=0;i<len(slice);i++ {
		for j:=0;j<len(slice)-1;j++ {
			if slice[j]>=slice[j+1]{
				slice[j],slice[j+1]=slice[j+1],slice[j]
			}
		}
	}
	fmt.Println("Newly sorted slice: ",slice)
	f:=0
	l:=len(slice)-1
	mid:= (f+l)/2
	fmt.Println("Middle element:",slice[mid])
	if len(slice)==2{
		if slice[f]==value{
			fmt.Println("Found!")
			return value
		}else if slice[l]==value{
			fmt.Println("Found!")
			return value
		}else{
			fmt.Println("Not found")
		}
	}else if f>l{
		fmt.Println("slice does not exist")
	}else{
		for i:=0;i<len(slice);i++{
			if slice[mid]==value{
				fmt.Println("Found!")
			}else if value<slice[mid]{
				fmt.Println("Discard greater values!")
				return binarySearch(slice[:mid],value)
			}else{
				fmt.Println("Discard less values!")
				return binarySearch(slice[mid:],value)
			}
		}
	}
	return value
}
