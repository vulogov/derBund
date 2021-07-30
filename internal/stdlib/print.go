package stdlib

import (
	"fmt"

	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/derBund/internal/vm"
)

func PrintElement(v *vm.VM, e *vm.Elem) (*vm.Elem, error) {
	eh, err := vm.GetType(e.Type)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s", eh.ToString(e))
	return &vm.Elem{Type: "bool", Value: true}, nil
}

func PrintlnElement(v *vm.VM, e *vm.Elem) (*vm.Elem, error) {
	eh, err := vm.GetType(e.Type)
	if err != nil {
		log.Errorf("Error getting element handler in println: %v", err)
		return nil, err
	}
	fmt.Printf("%s\n", eh.ToString(e))
	return e, nil
}

func InitPrintFunctions() {
	log.Debug("[ BUND ] bund.InitPrintFunctions() reached")
	vm.AddFunction("print", PrintElement)
	vm.AddFunction("println", PrintlnElement)
}
