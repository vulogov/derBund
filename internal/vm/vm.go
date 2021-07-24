package vm

import (
	"fmt"

	"github.com/gammazero/deque"
	"github.com/lrita/cmap"
	"github.com/pieterclaerhout/go-log"
)

type VM struct {
	Name        string
	Mode        bool
	IsIgnore    deque.Deque
	NSStack     deque.Deque
	NS          cmap.Cmap
	Current     *deque.Deque
	CurrentNS   *NS
	CurrentElem *Elem
}

func NewVM(name string) *VM {
	log.Debugf("Creating VM: %v", name)
	res := VM{Name: name, Current: nil, CurrentNS: nil, CurrentElem: nil, Mode: true}
	res.IsIgnore.PushBack(false)
	return &res
}

func (vm *VM) GetNS(name string) *NS {
	var res *NS
	if _res, ok := vm.NS.Load(name); ok {
		log.Debugf("Returning NAMESPACE: %v", name)
		res = _res.(*NS)
	} else {
		res = NewNS(name)
		vm.NS.Store(name, res)
	}
	if vm.CurrentNS != nil {
		vm.NSStack.PushBack(vm.CurrentNS)
	} else {
		log.Debugf("Skip sending an empty state to a NAMESPACE stack")
	}
	vm.CurrentNS = res
	vm.Current = &res.Stack
	return res
}

func (vm *VM) EndNS() *NS {
	var res *NS
	if vm.NSStack.Len() > 0 {
		_res := vm.NSStack.PopBack()
		if _res != nil {
			res = _res.(*NS)
			vm.CurrentNS = res
			vm.Current = &res.Stack
		} else {
			log.Debugf("NAMESPACE stack returns End_of_Stack")
			res = nil
			vm.CurrentNS = nil
			vm.Current = nil
		}
		log.Debugf("NAMESPACE stack %v size: %v", vm.CurrentNS.Name, vm.NSStack.Len())
	} else {
		log.Debugf("NAMESPACE stack is empty")
		res = nil
		vm.CurrentNS = nil
		vm.Current = nil
	}
	return res
}

func (vm *VM) IsStack() bool {
	if vm.CurrentNS == nil {
		return false
	}
	if vm.Current == nil {
		return false
	}
	return true
}

func (vm *VM) CanGet() bool {
	if !vm.IsStack() {
		return false
	}
	if vm.Current.Len() == 0 {
		return false
	}
	return true
}

func (vm *VM) Put(e *Elem) bool {
	if !vm.IsStack() {
		log.Errorf("Attempt to Put() but Stack doesn't exists")
		return false
	}
	if vm.Mode {
		vm.Current.PushBack(e)
	} else {
		vm.Current.PushFront(e)
	}
	return true
}

func (vm *VM) Get() *Elem {
	var res interface{}

	if !vm.CanGet() {
		log.Errorf("Attempt to Get() but Stack doesn't exists or empty")
		return nil
	}
	if vm.Mode {
		res = vm.Current.Back()
	} else {
		res = vm.Current.Front()
	}
	return res.(*Elem)
}

func (vm *VM) Take() *Elem {
	var res interface{}

	if !vm.CanGet() {
		log.Errorf("Attempt to Take() but Stack doesn't exists or empty")
		return nil
	}
	if vm.Mode {
		res = vm.Current.PopBack()
	} else {
		res = vm.Current.PopFront()
	}
	return res.(*Elem)
}

func (vm *VM) Apply(name string) error {
	err := Apply(name, vm)
	if err != nil {
		log.Errorf("Apply(): %v", err)
	}
	return err
}

func (vm *VM) Op(name string) error {
	if HasUserFunction(name, vm) {
		return vm.Apply(name)
	}
	if vm.Current.Len() < 2 {
		return fmt.Errorf("Attempt to call an operation: %v on empty stack", name)
	}
	f, err := GetOperator(name)
	if err != nil {
		return err
	}
	val1 := vm.Take()
	val2 := vm.Take()
	res, err := f(vm, val1, val2)
	if err != nil {
		return err
	}
	if !vm.Put(res) {
		return fmt.Errorf("Attempt to call a function: %v and error storing results", name)
	}
	return nil
}

func (vm *VM) Exec(name string) error {
	if HasUserFunction(name, vm) {
		return vm.Apply(name)
	}
	if vm.Current.Len() < 1 {
		return fmt.Errorf("Attempt to call a function: %v on empty stack", name)
	}
	f, err := GetFunction(name)
	if err != nil {
		return err
	}
	val := vm.Take()
	res, err := f(vm, val)
	if err != nil {
		return err
	}
	if !vm.Put(res) {
		return fmt.Errorf("Attempt to call a function: %v and error storing results", name)
	}
	return nil
}

func (vm *VM) Ignore() {
	vm.IsIgnore.PushBack(true)
}

func (vm *VM) NotIgnore() {
	vm.IsIgnore.PushBack(false)
}

func (vm *VM) MustIgnore() bool {
	if vm.IsIgnore.Len() == 0 {
		return false
	}
	res := vm.IsIgnore.PopBack()
	return res.(bool)
}

func (vm *VM) CheckIgnore() bool {
	if vm.IsIgnore.Len() == 0 {
		return false
	}
	res := vm.IsIgnore.Back()
	return res.(bool)
}

func (vm *VM) CurrentLambda() *deque.Deque {
	if !vm.IsStack() {
		log.Errorf("Attempt to CurrentLambda() but NS doesn't exists")
		return nil
	}
	return vm.CurrentNS.CurrentLambda()
}

func (vm *VM) InLambda() bool {
	if !vm.IsStack() {
		log.Debugf("Attempt to InLambda() but Stack doesn't exists")
		return false
	}
	if vm.CurrentNS.LambdasStack.Len() > 0 {
		log.Debugf("We are in InLambda(%v)", vm.CurrentNS.LambdasStack.Back().(string))
		return true
	}
	return false
}
