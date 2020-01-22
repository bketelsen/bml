package bml

type Node struct {
	ID string
	Type Token

}


type Block struct {
	Node
	Raw string
	Children []Node

}

