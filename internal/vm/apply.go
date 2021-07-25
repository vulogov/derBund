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
		}
	}
	return nil
}
