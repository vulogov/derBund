package stdlib

import (
	"fmt"

	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/derBund/internal/vm"
)

func GlobPatternMatchOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type != "glob" {
		return nil, fmt.Errorf("Pattern for '===' not found")
	}
	if e2.Type != "str" {
		return nil, fmt.Errorf("Pattern matching with '===' done only with strings")
	}
	eh, err := vm.GetType(e1.Type)
	if err != nil {
		return nil, err
	}
	res := eh.Compare(e1, e2)
	if res == vm.Eq {
		return &vm.Elem{Type: "bool", Value: true}, nil
	} else {
		return &vm.Elem{Type: "bool", Value: false}, nil
	}
}

func InitGPMOperators() {
	log.Debug("[ BUND ] bund.InitGPMOperators() reached")
	vm.AddOperator("===", GlobPatternMatchOperator)
}
