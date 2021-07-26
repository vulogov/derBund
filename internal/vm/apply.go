package vm

import (
	"fmt"

	"github.com/pieterclaerhout/go-log"
)

func Apply(name string, vm *VM) error {
	if !vm.IsStack() {
		return fmt.Errorf("Attempt to execute lambda %v on empty context", name)
	}
	ls := vm.CurrentNS.GetLambda(name)
	if ls == nil {
		return fmt.Errorf("Lambda %v not exist in %v", name, vm.Name)
	}
	log.Debugf("Executing lambda %v in %v", name, vm.Name)
	for ls.Len() > 0 {
		cmd := ls.PopFront().(*Elem)
		fmt.Println(cmd)
		switch cmd.Type {
		case "bool":
			log.Debugf("Value: %v", cmd.Value.(bool))
			vm.Put(cmd)
		case "str":
			log.Debugf("String: %v", cmd.Value.(string))
			vm.Put(cmd)
		case "int":
			log.Debugf("Int64: %v", cmd.Value.(int64))
			vm.Put(cmd)
		case "uint":
			log.Debugf("UInt64: %v", cmd.Value.(uint64))
			vm.Put(cmd)
		case "flt":
			log.Debugf("Float64: %v", cmd.Value.(float64))
			vm.Put(cmd)
		case "uflt":
			log.Debugf("UFloat64: %v", cmd.Value.(float64))
			vm.Put(cmd)
		case "cpx":
			log.Debugf("Complex128: %v", cmd.Value.(complex128))
			vm.Put(cmd)
		case "NS":
			log.Debugf("ENTERING Namespace: %v", cmd.Value.(string))
			vm.GetNS(cmd.Value.(string))
		case "exitNS":
			log.Debugf("EXITING Namespace: %v", cmd.Value.(string))
			vm.EndNS()
		case "BLOCK":
			log.Debugf("ENTERING Block: %v", cmd.Value.(string))
			vm.GetNS(cmd.Value.(string))
		case "exitBLOCK":
			log.Debugf("EXITING Block")
			vm.EndNS()
		case "CALL":
			log.Debugf("Calling: %v", cmd.Value.(string))
			vm.Exec(cmd.Value.(string))
		case "OP":
			log.Debugf("Operator: %v", cmd.Value.(string))
			vm.Op(cmd.Value.(string))
		case "DROP":
			log.Debugf("STACK: Drop")
			if vm.IsStack() {
				if vm.Current.Len() == 0 {
					log.Warn("Attempt to Drop value from an empty stack")
				} else {
					vm.Take()
				}
			} else {
				log.Error("Attempt to Drop value with empty context")
			}
		case "MODE":
			log.Debugf("Stack MODE: %v", cmd.Value.(bool))
			vm.Mode = cmd.Value.(bool)
		case "RETURN":
			log.Debugf("STACK: Return")
			if vm.NSStack.Len() < 1 {
				log.Errorf("Namespace stack is too shallow for RETURN operation: %v", vm.NSStack.Len())
				break
			}
			nsr := vm.NSStack.Back().(*NS)
			log.Debugf("Return push to %v", nsr.Name)
			var e *Elem
			switch cmd.Value.(string) {
			case "$":
				e = vm.Take()
			case "$$":
				e = vm.Get()
			default:
				log.Errorf("I do not know how to do this RETURN operation: %v", cmd.Value.(string))
			}
			if vm.Mode {
				nsr.Stack.PushBack(e)
			} else {
				nsr.Stack.PushFront(e)
			}
		case "DBLOCK":
			log.Debugf("Initialize DBlock")
			vm.GetNS(cmd.Value.(string))
		case "exitDBLOCK":
			if vm.Current != nil {
				log.Debugf("Create Data Block. Stack size: %v", vm.Current.Len())
				res := new(Elem)
				res.Type = "dblock"
				res.Value = vm.Current
				vm.EndNS()
				if vm.IsStack() {
					vm.Put(res)
				}
			} else {
				log.Debugf("Close Data Block. No current stack")
				vm.EndNS()
			}
		case "FBLOCK":
			log.Debugf("ENTERING Float Block")
			vm.GetNS(cmd.Value.(string))
		case "exitFBLOCK":
			if vm.Current != nil {
				log.Debugf("EXITING Float Block. Stack size: %v", vm.Current.Len())
				res := new(Elem)
				res.Type = "fblock"
				res.Value = vm.Current
				vm.EndNS()
				if vm.IsStack() {
					vm.Put(res)
				}
			} else {
				log.Debugf("EXITING Float Block. No current stack")
				vm.EndNS()
			}
		default:
			log.Errorf("Unknown OP-block: %v", cmd)
		}
	}
	return nil
}
