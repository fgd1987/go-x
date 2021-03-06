package main

import (
	"github.com/fananchong/go-x/common/discovery"
)

var (
	xnode *Node = NewNode()
)

type Node struct {
	discovery.Node
}

func NewNode() *Node {
	return &Node{}
}

func (this *Node) OnNodeUpdate(nodeType int, id string, data []byte) {
}

func (this *Node) OnNodeJoin(nodeType int, id string, data []byte) {
}

func (this *Node) OnNodeLeave(nodeType int, id string) {
}
