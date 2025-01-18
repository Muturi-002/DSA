//from this youtube video: https://youtu.be/0AucCCwmpnw. Also from this youtube video: https://youtu.be/b_NjndniOqY
//Creating a binary tree key structure in Go
package main
import "fmt"

type Node struct{
	key string
	children []*Node
}
/* In an in-order traversal, the nodes are recursively visited in this order:
1. Traverse the left subtree by recursively calling the in-order function.
2. Visit the root node.
3. Traverse the right subtree by recursively calling the in-order function.
*/
func inOrderTransversal(node *Node){
	if node==nil{
		return
	}
	if len(node.children)>0{
		inOrderTransversal(node.children[0])//Only applicable to the left leaf node from the root node
	}
	fmt.Print(node.key," ")
	for i:=1;i<len(node.children);i++{
		inOrderTransversal(node.children[i])//Only applicable to the right leaf node from the root node, for a binary tree that is.
	}
}
/* In a pre-order traversal, the nodes are recursively visited in this order:
1. Visit the root node.
2. Traverse the left subtree by recursively calling the pre-order function.
3. Traverse the right subtree by recursively calling the pre-order function.
*/
func preOrderTransversal(node *Node){
	if node==nil{
		return
	}
	fmt.Print(node.key," ")//makes a difference from the pre-order transversal. The root is printed out first in this case.
	for i:=0;i<len(node.children);i++{
		preOrderTransversal(node.children[i])
	}
}
/* In a post-order traversal, the nodes are recursively visited in this order:
1. Traverse the left subtree by recursively calling the post-order function.
2. Traverse the right subtree by recursively calling the post-order function.
3. Visit the root node.
*/
func postOrderTransversal(node *Node){
	if node==nil{
		return
	}
	for i:=0;i<len(node.children);i++{
		postOrderTransversal(node.children[i])
	}
	fmt.Print(node.key," ")//makes the difference from the post-order transversal. The root is printed out last in this case
}
func main(){
	fmt.Println("BINARY TREES  STRUCTURE")
	tree1 := &Node{key: "root"}
	fmt.Println("Tree root node: ", tree1)
	child := []string{"a", "b"}
	for _, v := range child {
		tree1.children = append(tree1.children, &Node{key: v})
		fmt.Println("Appended first child node(s) to root node: ", v)
	}
	fmt.Println("Tree children: ", tree1.children)
	gc := []string{"c", "d", "e", "f"}
	var count int=0//counts the number of grand children nodes appended to a child node, explicitly limiting the number of grand children nodes to 2
	for d := 0; d < len(tree1.children); d++ {
		for _, v := range gc {
			if len(tree1.children[d].children) < 2 {
				count++
				tree1.children[d].children = append(tree1.children[d].children, &Node{key: v})
				fmt.Println("Appended grandchild node(s)", v, "to child node:", child[d])
				if count==2{
					d+=1
				}
			}
		}
	}
	fmt.Println("Tree grand children: ", tree1.children[0].children, tree1.children[1].children)
	fmt.Println("=======================")
	fmt.Print("InOrderTransversal: \n")
	inOrderTransversal(tree1)
	fmt.Println("\n=======================")	
	fmt.Print("PreOrderTransversal: \n")
	preOrderTransversal(tree1)
	fmt.Println("\n=======================")
	fmt.Print("PostOrderTransversal: \n")
	postOrderTransversal(tree1)
	fmt.Println("\n=======================")
}