package stdlib

import (
	"fmt"
	"math/big"

	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/derBund/internal/vm"
)

func BigAbs(v *vm.VM, e *vm.Elem) (*vm.Elem, error) {
	if e.Type != "big" {
		return nil, fmt.Errorf("All big/* lambdas require Big data to work with")
	}
	d := e.Value.(*big.Int)
	d.Abs(d)
	return &vm.Elem{Type: "big", Value: d}, nil
}

func BigNeg(v *vm.VM, e *vm.Elem) (*vm.Elem, error) {
	if e.Type != "big" {
		return nil, fmt.Errorf("All big/* lambdas require Big data to work with")
	}
	d := e.Value.(*big.Int)
	d.Neg(d)
	return &vm.Elem{Type: "big", Value: d}, nil
}

func BigFromInt(v *vm.VM, e *vm.Elem) (*vm.Elem, error) {
	if e.Type != "int" {
		return nil, fmt.Errorf("big/FromInt lambdas require Int data to work with")
	}
	d := e.Value.(int64)
	res := big.NewInt(0)
	res.SetInt64(d)
	return &vm.Elem{Type: "big", Value: res}, nil
}

func BigToInt(v *vm.VM, e *vm.Elem) (*vm.Elem, error) {
	if e.Type != "big" {
		return nil, fmt.Errorf("All big/* lambdas require Big data to work with")
	}
	d := e.Value.(*big.Int)
	if !d.IsInt64() {
		return nil, fmt.Errorf("Can not convert  Big to Int")
	}
	return &vm.Elem{Type: "int", Value: d.Int64()}, nil
}

func BigFromUInt(v *vm.VM, e *vm.Elem) (*vm.Elem, error) {
	if e.Type != "uint" {
		return nil, fmt.Errorf("big/FromUInt lambdas require UInt data to work with")
	}
	d := e.Value.(uint64)
	res := big.NewInt(0)
	res.SetUint64(d)
	return &vm.Elem{Type: "big", Value: res}, nil
}

func BigToUInt(v *vm.VM, e *vm.Elem) (*vm.Elem, error) {
	if e.Type != "big" {
		return nil, fmt.Errorf("All big/* lambdas require Big data to work with")
	}
	d := e.Value.(*big.Int)
	if !d.IsInt64() {
		return nil, fmt.Errorf("Can not convert  Big to UInt")
	}
	return &vm.Elem{Type: "uint", Value: d.Uint64()}, nil
}

func InitBigFunctions() {
	log.Debug("[ BUND ] bund.InitBigFunctions() reached")
	vm.AddFunction("big/Abs", BigAbs)
	vm.AddFunction("big/Neg", BigNeg)
	vm.AddFunction("big/FromInt", BigFromInt)
	vm.AddFunction("big/FromUInt", BigFromUInt)
	vm.AddFunction("big/ToInt", BigToInt)
	vm.AddFunction("big/ToUInt", BigToUInt)
}
