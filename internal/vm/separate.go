package vm

import (
	"github.com/gammazero/deque"
	"github.com/pieterclaerhout/go-log"
)

func mergeBlocksWithSeparator(e1 *Elem, e2 *Elem) {
	switch e1.Type {
	case "dblock":
		switch e2.Type {
		case "dblock", "iblock", "uiblock", "fblock":
			q1 := e1.Value.(*deque.Deque)
			q2 := e2.Value.(*deque.Deque)
			q1.PushBack(&Elem{Type: "SEPARATE", Value: nil})
			for q2.Len() > 0 {
				e := q2.PopFront()
				q1.PushBack(e)
			}
			return
		case "int", "uint", "flt", "bool", "str":
			q1 := e1.Value.(*deque.Deque)
			q1.PushBack(&Elem{Type: "SEPARATE", Value: nil})
			q1.PushBack(e2)
			return
		}
		log.Errorf("I can not merge %v to dblock", e2.Type)
		return
	}
	log.Errorf("%v is not mergeable", e1.Type)
}

func Separate(vm *VM) {
	if vm.Current.Len() < 2 {
		log.Errorf("Your stack for the | command is too shallow")
		return
	}
	e1 := vm.Take()
	e2 := vm.Take()
	if e1.Type == e2.Type {
		switch e1.Type {
		case "int", "uint", "flt", "bool", "str":
			res := new(Elem)
			res.Type = "dblock"
			q := new(deque.Deque)
			res.Value = q
			q.PushBack(e2)
			q.PushBack(&Elem{Type: "SEPARATE", Value: nil})
			q.PushBack(e1)
			vm.Put(res)
			return
		case "dblock":
			mergeBlocksWithSeparator(e1, e2)
			return
		}
	} else {
		if e1.Type == "dblock" {
			mergeBlocksWithSeparator(e1, e2)
			return
		}
	}
	log.Errorf("Merge failed. Stack uncertain")
}
