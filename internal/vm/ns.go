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
	return res
}

func (ns *NS) InLambda(name string) bool {
	if _, ok := ns.Fun.Load(name); ok {
		ns.LambdasStack.PushBack(name)
		log.Debugf("We are go8ing in Lambda(%v)", name)
		return true
	}
	log.Errorf("Attempt to go in Lambda(%v) failed", name)
	return false
}

func (ns *NS) CurrentLambda() *deque.Deque {
	var res *deque.Deque
	if ns.LambdasStack.Len() < 1 {
		log.Errorf("Attempt to select Lambda fuction on empty Lambdas stack")
		return nil
	}
	name := ns.LambdasStack.Back().(string)
	if _res, ok := ns.Fun.Load(name); ok {
		log.Debugf("Returning LAMBDA from %v: %v", ns.Name, name)
		res = _res.(*deque.Deque)
		return res
	}
	log.Errorf("Something is seriously wrong, current lambda %v not found", name)
	return nil
}

func (ns *NS) CloseLambda() bool {
	if ns.LambdasStack.Len() < 1 {
		log.Errorf("Attempt to close Lambda fuction on empty Lambdas stack")
		return false
	}
	ln := ns.LambdasStack.PopBack().(string)
	log.Debugf("Closing lambda %v. Stack size: %v", ln, ns.LambdasStack.Len())
	return true
}
