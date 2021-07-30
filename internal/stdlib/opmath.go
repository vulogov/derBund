package stdlib

import (
	"fmt"
	"math"
	"math/big"

	"github.com/pieterclaerhout/go-log"
	"gonum.org/v1/gonum/floats"

	"github.com/vulogov/derBund/internal/vm"
)

func AddOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type != e2.Type {
		return nil, fmt.Errorf("Datatypes for operation '+' must be the same")
	}
	switch e1.Type {
	case "int":
		return &vm.Elem{Type: "int", Value: (e1.Value.(int64) + e2.Value.(int64))}, nil
	case "uint":
		return &vm.Elem{Type: "uint", Value: (e1.Value.(uint64) + e2.Value.(uint64))}, nil
	case "big":
		d1 := e1.Value.(*big.Int)
		d2 := e2.Value.(*big.Int)
		d1.Add(d1, d2)
		return &vm.Elem{Type: "big", Value: d1}, nil
	case "flt":
		return &vm.Elem{Type: "flt", Value: (e1.Value.(float64) + e2.Value.(float64))}, nil
	case "uflt":
		return &vm.Elem{Type: "uflt", Value: (e1.Value.(float64) + e2.Value.(float64))}, nil
	case "str":
		return &vm.Elem{Type: "str", Value: (e1.Value.(string) + e2.Value.(string))}, nil
	case "fblock":
		ar1 := vm.BlockToArray(e1)
		ar2 := vm.BlockToArray(e2)
		if len(ar1) != len(ar2) {
			return nil, fmt.Errorf("Datatypes for operation '+' must be the same arity")
		}
		floats.Add(ar1, ar2)
		return vm.ArrayToBlock("fblock", ar1), nil
	case "iblock":
		ar1 := vm.BlockToArray(e1)
		ar2 := vm.BlockToArray(e2)
		if len(ar1) != len(ar2) {
			return nil, fmt.Errorf("Datatypes for operation '+' must be the same arity")
		}
		floats.Add(ar1, ar2)
		return vm.ArrayToBlock("iblock", ar1), nil
	case "uiblock":
		ar1 := vm.BlockToArray(e1)
		ar2 := vm.BlockToArray(e2)
		if len(ar1) != len(ar2) {
			return nil, fmt.Errorf("Datatypes for operation '+' must be the same arity")
		}
		floats.Add(ar1, ar2)
		return vm.ArrayToBlock("uiblock", ar1), nil
	}
	return nil, fmt.Errorf("I do not know how to perform '+' for this data")
}

func AddEqualOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type == e2.Type {
		return AddOperator(v, e1, e2)
	}
	switch e1.Type {
	case "iblock":
		if e2.Type == "int" {
			ar1 := vm.BlockToArray(e1)
			floats.AddConst(float64(e2.Value.(int64)), ar1)
			return vm.ArrayToBlock("iblock", ar1), nil
		}
	case "uiblock":
		if e2.Type == "uint" {
			ar1 := vm.BlockToArray(e1)
			floats.AddConst(float64(e2.Value.(uint64)), ar1)
			return vm.ArrayToBlock("uiblock", ar1), nil
		}
	case "fblock":
		if e2.Type == "flt" {
			ar1 := vm.BlockToArray(e1)
			floats.AddConst(e2.Value.(float64), ar1)
			return vm.ArrayToBlock("fblock", ar1), nil
		}
	}
	return nil, fmt.Errorf("I do not know how to perform '+=' for this data: %v %v", e1.Type, e2.Type)
}

func MulOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type != e2.Type {
		return nil, fmt.Errorf("Datatypes for operation '×' must be the same")
	}
	switch e1.Type {
	case "int":
		return &vm.Elem{Type: "int", Value: (e1.Value.(int64) * e2.Value.(int64))}, nil
	case "uint":
		return &vm.Elem{Type: "uint", Value: (e1.Value.(uint64) * e2.Value.(uint64))}, nil
	case "big":
		d1 := e1.Value.(*big.Int)
		d2 := e2.Value.(*big.Int)
		d1.Mul(d1, d2)
		return &vm.Elem{Type: "big", Value: d1}, nil
	case "flt":
		return &vm.Elem{Type: "flt", Value: (e1.Value.(float64) * e2.Value.(float64))}, nil
	case "uflt":
		return &vm.Elem{Type: "uflt", Value: (e1.Value.(float64) * e2.Value.(float64))}, nil
	case "fblock":
		ar1 := vm.BlockToArray(e1)
		ar2 := vm.BlockToArray(e2)
		if len(ar1) != len(ar2) {
			return nil, fmt.Errorf("Datatypes for operation '*' must be the same arity")
		}
		floats.Mul(ar1, ar2)
		return vm.ArrayToBlock("fblock", ar1), nil
	case "iblock":
		ar1 := vm.BlockToArray(e1)
		ar2 := vm.BlockToArray(e2)
		if len(ar1) != len(ar2) {
			return nil, fmt.Errorf("Datatypes for operation '*' must be the same arity")
		}
		floats.Mul(ar1, ar2)
		return vm.ArrayToBlock("iblock", ar1), nil
	case "uiblock":
		ar1 := vm.BlockToArray(e1)
		ar2 := vm.BlockToArray(e2)
		if len(ar1) != len(ar2) {
			return nil, fmt.Errorf("Datatypes for operation '*' must be the same arity")
		}
		floats.Mul(ar1, ar2)
		return vm.ArrayToBlock("uiblock", ar1), nil
	}
	return nil, fmt.Errorf("I do not know how to perform '×' for this data")
}

func MulEqualOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type == e2.Type {
		return MulOperator(v, e1, e2)
	}
	switch e1.Type {
	case "iblock":
		if e2.Type == "int" {
			ar1 := vm.BlockToArray(e1)
			vm.MulConst(float64(e2.Value.(int64)), ar1)
			return vm.ArrayToBlock("iblock", ar1), nil
		}
	case "uiblock":
		if e2.Type == "uint" {
			ar1 := vm.BlockToArray(e1)
			vm.MulConst(float64(e2.Value.(uint64)), ar1)
			return vm.ArrayToBlock("uiblock", ar1), nil
		}
	case "fblock":
		if e2.Type == "flt" {
			ar1 := vm.BlockToArray(e1)
			vm.MulConst(e2.Value.(float64), ar1)
			return vm.ArrayToBlock("fblock", ar1), nil
		}
	}
	return nil, fmt.Errorf("I do not know how to perform '+=' for this data: %v %v", e1.Type, e2.Type)
}

func DivOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type != e2.Type {
		return nil, fmt.Errorf("Datatypes for operation '÷' must be the same")
	}
	switch e1.Type {
	case "int":
		return &vm.Elem{Type: "int", Value: (e1.Value.(int64) / e2.Value.(int64))}, nil
	case "uint":
		return &vm.Elem{Type: "uint", Value: (e1.Value.(uint64) / e2.Value.(uint64))}, nil
	case "big":
		d1 := e1.Value.(*big.Int)
		d2 := e2.Value.(*big.Int)
		d1.Div(d1, d2)
		return &vm.Elem{Type: "big", Value: d1}, nil
	case "flt":
		return &vm.Elem{Type: "flt", Value: (e1.Value.(float64) / e2.Value.(float64))}, nil
	case "uflt":
		return &vm.Elem{Type: "uflt", Value: (e1.Value.(float64) / e2.Value.(float64))}, nil
	case "fblock":
		ar1 := vm.BlockToArray(e1)
		ar2 := vm.BlockToArray(e2)
		if len(ar1) != len(ar2) {
			return nil, fmt.Errorf("Datatypes for operation '÷' must be the same arity")
		}
		floats.Div(ar1, ar2)
		return vm.ArrayToBlock("fblock", ar1), nil
	case "iblock":
		ar1 := vm.BlockToArray(e1)
		ar2 := vm.BlockToArray(e2)
		if len(ar1) != len(ar2) {
			return nil, fmt.Errorf("Datatypes for operation '÷' must be the same arity")
		}
		floats.Div(ar1, ar2)
		return vm.ArrayToBlock("iblock", ar1), nil
	case "uiblock":
		ar1 := vm.BlockToArray(e1)
		ar2 := vm.BlockToArray(e2)
		if len(ar1) != len(ar2) {
			return nil, fmt.Errorf("Datatypes for operation '÷' must be the same arity")
		}
		floats.Div(ar1, ar2)
		return vm.ArrayToBlock("uiblock", ar1), nil
	}
	return nil, fmt.Errorf("I do not know how to perform '÷' for this data")
}

func DivEqualOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type == e2.Type {
		return DivOperator(v, e1, e2)
	}
	switch e1.Type {
	case "iblock":
		if e2.Type == "int" {
			ar1 := vm.BlockToArray(e1)
			vm.DivConst(float64(e2.Value.(int64)), ar1)
			return vm.ArrayToBlock("iblock", ar1), nil
		}
	case "uiblock":
		if e2.Type == "uint" {
			ar1 := vm.BlockToArray(e1)
			vm.DivConst(float64(e2.Value.(uint64)), ar1)
			return vm.ArrayToBlock("uiblock", ar1), nil
		}
	case "fblock":
		if e2.Type == "flt" {
			ar1 := vm.BlockToArray(e1)
			vm.DivConst(e2.Value.(float64), ar1)
			return vm.ArrayToBlock("fblock", ar1), nil
		}
	}
	return nil, fmt.Errorf("I do not know how to perform '+=' for this data: %v %v", e1.Type, e2.Type)
}

func SubOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type != e2.Type {
		return nil, fmt.Errorf("Datatypes for operation '-' must be the same")
	}
	switch e1.Type {
	case "int":
		return &vm.Elem{Type: "int", Value: (e1.Value.(int64) - e2.Value.(int64))}, nil
	case "uint":
		return &vm.Elem{Type: "uint", Value: (e1.Value.(uint64) - e2.Value.(uint64))}, nil
	case "big":
		d1 := e1.Value.(*big.Int)
		d2 := e2.Value.(*big.Int)
		d1.Sub(d1, d2)
		return &vm.Elem{Type: "big", Value: d1}, nil
	case "flt":
		return &vm.Elem{Type: "flt", Value: (e1.Value.(float64) - e2.Value.(float64))}, nil
	case "uflt":
		return &vm.Elem{Type: "uflt", Value: (e1.Value.(float64) - e2.Value.(float64))}, nil
	case "fblock":
		ar1 := vm.BlockToArray(e1)
		ar2 := vm.BlockToArray(e2)
		if len(ar1) != len(ar2) {
			return nil, fmt.Errorf("Datatypes for operation '-' must be the same arity")
		}
		floats.Sub(ar1, ar2)
		return vm.ArrayToBlock("fblock", ar1), nil
	case "iblock":
		ar1 := vm.BlockToArray(e1)
		ar2 := vm.BlockToArray(e2)
		if len(ar1) != len(ar2) {
			return nil, fmt.Errorf("Datatypes for operation '-' must be the same arity")
		}
		floats.Sub(ar1, ar2)
		return vm.ArrayToBlock("iblock", ar1), nil
	case "uiblock":
		ar1 := vm.BlockToArray(e1)
		ar2 := vm.BlockToArray(e2)
		if len(ar1) != len(ar2) {
			return nil, fmt.Errorf("Datatypes for operation '-' must be the same arity")
		}
		floats.Sub(ar1, ar2)
		return vm.ArrayToBlock("uiblock", ar1), nil
	}
	return nil, fmt.Errorf("I do not know how to perform '-' for this data")
}

func SubEqualOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type == e2.Type {
		return SubOperator(v, e1, e2)
	}
	switch e1.Type {
	case "iblock":
		if e2.Type == "int" {
			ar1 := vm.BlockToArray(e1)
			vm.SubConst(float64(e2.Value.(int64)), ar1)
			return vm.ArrayToBlock("iblock", ar1), nil
		}
	case "uiblock":
		if e2.Type == "uint" {
			ar1 := vm.BlockToArray(e1)
			vm.SubConst(float64(e2.Value.(uint64)), ar1)
			return vm.ArrayToBlock("uiblock", ar1), nil
		}
	case "fblock":
		if e2.Type == "flt" {
			ar1 := vm.BlockToArray(e1)
			vm.SubConst(e2.Value.(float64), ar1)
			return vm.ArrayToBlock("fblock", ar1), nil
		}
	}
	return nil, fmt.Errorf("I do not know how to perform '+=' for this data: %v %v", e1.Type, e2.Type)
}

func ExpOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type == e2.Type {
		switch e1.Type {
		case "int":
			r1 := e1.Value.(int64)
			r2 := e2.Value.(int64)
			return &vm.Elem{Type: "int", Value: int64(math.Pow(float64(r1), float64(r2)))}, nil
		case "uint":
			r1 := e1.Value.(uint64)
			r2 := e2.Value.(uint64)
			return &vm.Elem{Type: "uint", Value: uint64(math.Pow(float64(r1), float64(r2)))}, nil
		case "flt":
			r1 := e1.Value.(float64)
			r2 := e2.Value.(float64)
			return &vm.Elem{Type: "flt", Value: math.Pow(r1, r2)}, nil
		case "big":
			res := big.NewInt(0)
			r1 := e1.Value.(*big.Int)
			r2 := e2.Value.(*big.Int)
			return &vm.Elem{Type: "big", Value: res.Exp(r1, r2, nil)}, nil
		}
	}
	return nil, fmt.Errorf("I do not know how to perform '+=' for this data: %v %v", e1.Type, e2.Type)
}

func InitMathOperators() {
	log.Debug("[ BUND ] bund.InitMathOperators() reached")
	vm.AddOperator("+", AddOperator)
	vm.AddOperator("*", MulOperator)
	vm.AddOperator("×", MulOperator)
	vm.AddOperator("÷", DivOperator)
	vm.AddOperator("\\", DivOperator)
	vm.AddOperator("-", SubOperator)
	vm.AddOperator("+=", AddEqualOperator)
	vm.AddOperator("*=", MulEqualOperator)
	vm.AddOperator("×=", MulEqualOperator)
	vm.AddOperator("\\=", DivEqualOperator)
	vm.AddOperator("÷=", DivEqualOperator)
	vm.AddOperator("-=", SubEqualOperator)
	vm.AddOperator("**", ExpOperator)
	vm.AddOperator("↑", ExpOperator)
}
