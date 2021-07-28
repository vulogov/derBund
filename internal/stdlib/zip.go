package stdlib

import (
	"fmt"

	"github.com/gammazero/deque"
	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/derBund/internal/vm"
)

func mergeQueue(q *deque.Deque, e1 *vm.Elem, e2 *vm.Elem) {
	q1 := e1.Value.(*deque.Deque)
	q2 := e2.Value.(*deque.Deque)
	for i := 0; i < q1.Len(); i++ {
		q.PushBack(q1.At(i).(*vm.Elem))
	}
	for i := 0; i < q2.Len(); i++ {
		q.PushBack(q2.At(i).(*vm.Elem))
	}
}

func ZipOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	var e *vm.Elem
	e = new(vm.Elem)
	q := new(deque.Deque)
	e.Value = q
	if e1.Type == e2.Type {
		switch e1.Type {
		case "int", "iblock":
			e.Type = "iblock"
		case "uint", "uiblock":
			e.Type = "uiblock"
		case "flt", "fblock":
			e.Type = "fblock"
		case "bool", "str", "dblock":
			e.Type = "dblock"
		default:
			return nil, fmt.Errorf("I do not know how to perform '++' for this data: %v %v", e1.Type, e2.Type)
		}
		switch e1.Type {
		case "int", "uint", "flt", "bool", "str":
			q.PushBack(e1)
			q.PushBack(e2)
			return e, nil
		case "iblock", "uiblock", "fblock", "dblock":
			mergeQueue(q, e1, e2)
			return e, nil
		}
	}
	switch e1.Type {
	case "iblock":
		if e2.Type == "int" {
			e1.Value.(*deque.Deque).PushBack(e2)

		} else {
			return nil, fmt.Errorf("Attempt to merge data of type %v to (int ) failed", e2.Type)
		}
	case "uiblock":
		if e2.Type == "uint" {
			e1.Value.(*deque.Deque).PushBack(e2)
		} else {
			return nil, fmt.Errorf("Attempt to merge data of type %v to (uint ) failed", e2.Type)
		}
	case "fblock":
		if e2.Type == "flt" {
			e1.Value.(*deque.Deque).PushBack(e2)
		} else {
			return nil, fmt.Errorf("Attempt to merge data of type %v to (float ) failed", e2.Type)
		}
	case "dblock":
		e1.Value.(*deque.Deque).PushBack(e2)
	default:
		return nil, fmt.Errorf("I finally do not know how to perform '+=' for this data: %v %v", e1.Type, e2.Type)
	}
	return e1, nil
}

func InitZipOperators() {
	log.Debug("[ BUND ] bund.InitZipOperators() reached")
	vm.AddOperator("++", ZipOperator)
}
