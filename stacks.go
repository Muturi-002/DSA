//from https://www.youtube.com/watch?v=fsbm1FOSDJ0&ab_channel=JamieDev
package main
import "fmt"
type stack struct{
	items []int
}


func (s *stack)push(i int){
	s.items=append(s.items,i)
}
func (s *stack)pop(){
	s.isEmpty()
	s.items=s.items[:len(s.items)-1]
}
func (s *stack)isEmpty(){
	l:=len(s.items)
	if l<0{
		fmt.Println("Non-existent stack")
		return
	}
	if l==0{
		fmt.Println("Empty stack")
	}
}
func (s *stack)peek(){
	s.isEmpty()
	top:=s.items[len(s.items)-1]
	fmt.Println("Top of the stack:",top)
}
func main(){
	var myStack stack
	//alternative declaration of stack-> myStack:=stack{}
	fmt.Println(myStack)
	myStack.push(8)
	myStack.push(9)
	myStack.push(10)
	fmt.Println(myStack)
	myStack.pop()
	fmt.Println(myStack)
	myStack.pop()
	fmt.Println(myStack)
	myStack.pop()
	fmt.Println(myStack)
	myStack.pop()
	
}