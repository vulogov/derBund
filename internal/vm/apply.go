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
		switch cmd.Type {
		case "bool":
			log.Debugf("Value: %v", cmd.Value.(bool))
			vm.Put(cmd)
		case "str":
			log.Debugf("String: %v", cmd.Value.(string))
			vm.Put(cmd)
		case "CALL":
			log.Debugf("Calling: %v", cmd.Value.(string))
			vm.Exec(cmd.Value.(string))
		}
	}
	return nil
}
