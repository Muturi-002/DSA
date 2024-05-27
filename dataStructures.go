package main
import "fmt"
func main(){
	fmt.Println(".....")
	array1:= [5]float64{9.44,02.88,0.9,12.9783,88.34}
	fmt.Println(array1)
	for _,m:=range array1{
		fmt.Println(m)
	}
	slice1:= array1[:]
	slice1= append(slice1,234.86,92.00)
	fmt.Println(slice1)
	fmt.Println(slice1[:4])
	fmt.Println(slice1[4:])
	mapping:= map[int]string{1:"ball",2:"kick",3:"catch",4:"fetch",5:"run",6:"sprint"}
	fmt.Println(mapping)
	
}