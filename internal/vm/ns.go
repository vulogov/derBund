package vm

import (
	"github.com/gammazero/deque"
	"github.com/lrita/cmap"
	"github.com/pieterclaerhout/go-log"
)

type NS struct {
	Name         string
	Stack        deque.Deque
	Fun          cmap.Cmap
	LambdasStack deque.Deque
}

func NewNS(name string) *NS {
	log.Debugf("Creating NAMESPACE: %v", name)
	return &NS{Name: name}
}

func (ns *NS) HasLambda(name string) bool {
	if _, ok := ns.Fun.Load(name); ok {
		return true
	}
	return false
}

func (ns *NS) GetLambda(name string) *deque.Deque {
	var res *deque.Deque
	if _res, ok := ns.Fun.Load(name); ok {
		log.Debugf("Returning LAMBDA from %v: %v", ns.Name, name)
		res = _res.(*deque.Deque)
	} else {
		log.Debugf("Creating LAMBDA in %v: %v", ns.Name, name)
		res = new(deque.Deque)
		ns.Fun.Store(name, res)
	}
	ns.LambdasStack.PushBack(name)
	return res
}

func (ns *NS) CurrentLambda() *deque.Deque {
	if ns.LambdasStack.Len() < 1 {
		log.Errorf("Attempt to select Lambda fuction on empty Lambdas stack")
		return nil
	}
	return ns.GetLambda(ns.LambdasStack.Back().(string))
}

func (ns *NS) CloseLambda() bool {
	if ns.LambdasStack.Len() < 1 {
		log.Errorf("Attempt to close Lambda fuction on empty Lambdas stack")
		return false
	}
	ns.LambdasStack.PopBack()
	return true
}
