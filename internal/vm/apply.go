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
		}
	}
	return nil
}
