//https://www.youtube.com/watch?v=1S0_-VxPLJo&list=PL0q7mDmXPZm7s7weikYLpNZBKk5dCoWm6&index=2
package main
import "fmt"
type node struct{
	data int
	next *node//variable 'pointer' type should have the same name as its parent struct
}
type linkedList struct{
	head *node//points to the memory address
	length int
}
//this function implements a stack data structure
func (l *linkedList)insertAtBegin (n *node){//pointer receiver
	second:=l.head
	l.head=n
	l.head.next=second
	l.length++
}
func (l linkedList) printListData(){//value receiver
	fmt.Println("Linked list data")
	for l.length!=0{
		fmt.Print(" ",l.head.data,"\n")
		l.head=l.head.next
		l.length--
	}
}
func (l *linkedList) deleteWithValue(value int){
	if l.length <=0{
		return
	}
	if l.head.data==value{
		l.head=l.head.next
		l.length--
		return
	}
	previousToDelete:= l.head
	for previousToDelete.next.data!= value{
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next=previousToDelete.next.next
	l.length--
}
func main() {
	fmt.Println("boom!")
	myL_list:=linkedList{}
	node1:=&node{data:40}
	node2:=&node{data:30}
	node3:=&node{data:20}
	node4:=&node{data:10}
	node5:=&node{data:50}
	node6:=&node{data:60}
	myL_list.insertAtBegin(node1)
	myL_list.insertAtBegin(node2)
	myL_list.insertAtBegin(node3)
	//fmt.Println("next in line: ",node2.next.data)
	myL_list.insertAtBegin(node4)
	myL_list.insertAtBegin(node5)
	myL_list.insertAtBegin(node6)
	fmt.Println("Linked List: ", myL_list)
	myL_list.printListData()
	myL_list.deleteWithValue(30)
	myL_list.printListData()

}