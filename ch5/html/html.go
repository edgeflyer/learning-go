package html

import "io"

// HTML DOM书中的一个节点
type Node struct {
	Type                    NodeType    // 节点的类型
	Data                    string      // 节点的标签名或文本内容
	Attr                    []Attribute // 存储结点的属性，每个都包含KV
	FirstChild, NextSibling *Node       // 指向该节点的第一个子节点和下一个兄弟节点
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
