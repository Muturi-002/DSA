//from the youtube video:https://youtu.be/Id099WXOKNM
//creating a binary tree
package main
import "fmt"
var count int//counts the number of nodes in the tree transversed to find a number
type TreeNode struct{
	key int
	left *TreeNode//left child
	right *TreeNode//right child
}
func (node *TreeNode) Insert(value int){
	if node.key > value{
		//move to the left
		if node.left==nil{
			node.left = &TreeNode{key:value}
		}else{
			node.left.Insert(value) //recursion
		}
	}
	if node.key<value{
		//move to the right
		if node.right==nil{
			node.right = &TreeNode{key:value}
		}else{
			node.right.Insert(value)
		}
	}
}
func (node TreeNode)Search(value int) bool{
	count++
	if node.key==value{
		return true
	}
	if node.key>value{
		//move to the left
		if node.left==nil{
			return false
		}
		return node.left.Search(value)
	}
	if node.key<value{
		//move to the right
		if node.right==nil{
			return false
		}
		return node.right.Search(value)
	}
	return false
} 
func main(){
	defer fmt.Println("Binary Tree")
	tree:= &TreeNode{key:57}
	fmt.Println(tree)
	/* tree.Insert(60)
	tree.Insert(50) */
	fmt.Println(tree)
	fmt.Println("Adding new random values---")
	values:=[]int{45, 67, 23, 89, 78, 34, 56, 90, 12, 45, 67, 23, 89, 78, 34, 56, 90, 12}
	for _,v:=range values{
		tree.Insert(v)
		fmt.Println("Cheking position of ",v," in the tree: ",tree)
	}
	var v int
	fmt.Printf("Enter a value to search in the tree: ")
	fmt.Scanf("%d",&v)
	fmt.Println("Searching for",v," in the tree: ",tree.Search(v))
}