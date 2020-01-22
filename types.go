package bml

type Node struct {
	ID string
	Type Token
	Children []Node
}


type Block struct {
	Node
	
}