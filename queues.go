//from https://www.youtube.com/watch?v=fsbm1FOSDJ0&ab_channel=JamieDev
package main
import "fmt"
type queue struct {
	items []int
}
func (q *queue) Enqueue(i int){
	q.items = append(q.items, i)
}
func (q *queue) Dequeue(){
	removed:= q.items[0]
	q.items = q.items[1:]
	/* _,err:=recover();if err != nil{
		return
	} */
	fmt.Println("removed: ",removed)

}
func (q *queue)IsEmpty(){
	if len(q.items)==0{
		fmt.Println("Empty queue")
	}
	if len(q.items)<0{
		panic("Non-existent queue!!!!")
	}
}
func main(){
	var newQueue queue
	//alternative declaration of queue-> newQueue:=queue{}
	fmt.Println("Queue: ",newQueue.items)
	newQueue.Enqueue(1)
	fmt.Println("Queue: ",newQueue.items)
	newQueue.Enqueue(2)
	fmt.Println("Queue: ",newQueue.items)
	newQueue.Enqueue(3)
	newQueue.Dequeue()
	fmt.Println("Queue: ",newQueue.items)
	newQueue.Dequeue()
	fmt.Println("Queue: ",newQueue.items)
	newQueue.Dequeue()
	fmt.Println("Queue: ",newQueue.items)
}