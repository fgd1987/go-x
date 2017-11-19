package discovery

import "github.com/fananchong/go-discovery"

type Node struct {
	godiscovery.Node
	Info ServerInfo
}

func (this *Node) OnNodeUpdate(nodeType int, id string, data []byte) {
}

func (this *Node) OnNodeJoin(nodeType int, id string, data []byte) {
}

func (this *Node) OnNodeLeave(nodeType int, id string) {
}

func (this *Node) GetBase() interface{} {
	return this
}
