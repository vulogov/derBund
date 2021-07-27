package stdlib

import (
	"fmt"

	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/derBund/internal/vm"
)

func EqOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type != e2.Type {
		return nil, fmt.Errorf("Datatypes for operation '=' must be the same")
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

func MoreOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type != e2.Type {
		return nil, fmt.Errorf("Datatypes for operation '>' must be the same")
	}
	eh, err := vm.GetType(e1.Type)
	if err != nil {
		return nil, err
	}
	res := eh.Compare(e1, e2)
	if res == vm.Gt {
		return &vm.Elem{Type: "bool", Value: true}, nil
	} else {
		return &vm.Elem{Type: "bool", Value: false}, nil
	}
}

func LessOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type != e2.Type {
		return nil, fmt.Errorf("Datatypes for operation '>' must be the same")
	}
	eh, err := vm.GetType(e1.Type)
	if err != nil {
		return nil, err
	}
	res := eh.Compare(e1, e2)
	if res == vm.Ls {
		return &vm.Elem{Type: "bool", Value: true}, nil
	} else {
		return &vm.Elem{Type: "bool", Value: false}, nil
	}
}

func MoreEqOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type != e2.Type {
		return nil, fmt.Errorf("Datatypes for operation '=' must be the same")
	}
	eh, err := vm.GetType(e1.Type)
	if err != nil {
		return nil, err
	}
	res := eh.Compare(e1, e2)
	if res == vm.Eq || res == vm.Gt {
		return &vm.Elem{Type: "bool", Value: true}, nil
	} else {
		return &vm.Elem{Type: "bool", Value: false}, nil
	}
}

func LessEqOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type != e2.Type {
		return nil, fmt.Errorf("Datatypes for operation '=' must be the same")
	}
	eh, err := vm.GetType(e1.Type)
	if err != nil {
		return nil, err
	}
	res := eh.Compare(e1, e2)
	if res == vm.Eq || res == vm.Ls {
		return &vm.Elem{Type: "bool", Value: true}, nil
	} else {
		return &vm.Elem{Type: "bool", Value: false}, nil
	}
}

func NeqOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type != e2.Type {
		return nil, fmt.Errorf("Datatypes for operation '=' must be the same")
	}
	eh, err := vm.GetType(e1.Type)
	if err != nil {
		return nil, err
	}
	res := eh.Compare(e1, e2)
	if res != vm.Eq {
		return &vm.Elem{Type: "bool", Value: true}, nil
	} else {
		return &vm.Elem{Type: "bool", Value: false}, nil
	}
}

func InitCMPOperators() {
	log.Debug("[ BUND ] bund.InitCMPOperators() reached")
	vm.AddOperator("=", EqOperator)
	vm.AddOperator(">", MoreOperator)
	vm.AddOperator("=>", MoreEqOperator)
	vm.AddOperator("<", LessOperator)
	vm.AddOperator("<=", LessEqOperator)
	vm.AddOperator("<>", NeqOperator)
}
