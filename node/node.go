package node

import (
	"github.com/gxlb/godi/service"
)

var instance = NewNode()

func Register(s ...service.IService) {

}

func Start() {

}

func Stop() {

}

func NewNode() *Node {

}

type Node struct {
	services []service.IService
}
