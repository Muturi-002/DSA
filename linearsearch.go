package main
import (
	"fmt"
	//"math/rand"....
)

func main() {
	s:= []int{66,33,22,92,1,29,73}
	var find int
	fmt.Print("Enter number of value to search for: ")
	fmt.Scan(&find)

	/* var N int= len(s)
	for i:=0;i<N;i++ {
		s=append(s, rand.Intn(9))
	}
	fmt.Println("Newly created slice",s,"\n Length: ",N)
	for j,s:=range s{
		fmt.Println(j,s)
	} */
	fmt.Println("Newly created slice",s)
	result:=linearSearch(s,find)
	fmt.Println(result)
}
func linearSearch(slice []int,value int)int {
	for i:=0;i<len(slice);i++{
		if slice[i]==value{
			fmt.Println("Found: ",slice[i])
		}
	}
	return value
	
}
