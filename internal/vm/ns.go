package vm

import (
	"github.com/gammazero/deque"
	"github.com/lrita/cmap"
	"github.com/pieterclaerhout/go-log"
)

type NS struct {
	Name  string
	Stack deque.Deque
	Fun   cmap.Cmap
}

func NewNS(name string) *NS {
	log.Debugf("Creating NAMESPACE: %v", name)
	return &NS{Name: name}
}
