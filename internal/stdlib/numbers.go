package stdlib

import (
	"fmt"
	"math"
	"math/big"

	"github.com/pieterclaerhout/go-log"
	"gonum.org/v1/gonum/floats"

	"github.com/vulogov/derBund/internal/vm"
)

func SpanFblock(v *vm.VM, e *vm.Elem) (*vm.Elem, error) {
	if e.Type != "dblock" {
		return nil, fmt.Errorf("For numbers/Span expecting (* ) in the stack")
	}
	N, err := vm.BlockAt(e, "int", 0)
	if err != nil {
		return nil, err
	}
	l, err := vm.BlockAt(e, "flt", 1)
	if err != nil {
		return nil, err
	}
	u, err := vm.BlockAt(e, "flt", 2)
	if err != nil {
		return nil, err
	}
	res := make([]float64, int(N.(int64)))
	floats.Span(res, l.(float64), u.(float64))
	return vm.ArrayToBlock("fblock", res), nil
}

func LogSpanFblock(v *vm.VM, e *vm.Elem) (*vm.Elem, error) {
	if e.Type != "dblock" {
		return nil, fmt.Errorf("For numbers/Span expecting (* ) in the stack")
	}
	N, err := vm.BlockAt(e, "int", 0)
	if err != nil {
		return nil, err
	}
	l, err := vm.BlockAt(e, "flt", 1)
	if err != nil {
		return nil, err
	}
	u, err := vm.BlockAt(e, "flt", 2)
	if err != nil {
		return nil, err
	}
	res := make([]float64, int(N.(int64)))
	floats.LogSpan(res, l.(float64), u.(float64))
	return vm.ArrayToBlock("fblock", res), nil
}

func ReverseBlock(v *vm.VM, e *vm.Elem) (*vm.Elem, error) {
	switch e.Type {
	case "iblock", "uiblock", "fblock":
		ar1 := vm.BlockToArray(e)
		floats.Reverse(ar1)
		return vm.ArrayToBlock(e.Type, ar1), nil
	}
	return nil, fmt.Errorf("I do not know how to do numbers/Reverse for %v", e.Type)
}

func SumBlock(v *vm.VM, e *vm.Elem) (*vm.Elem, error) {
	switch e.Type {
	case "iblock", "uiblock", "fblock":
		ar1 := vm.BlockToArray(e)
		res := floats.SumCompensated(ar1)
		return &vm.Elem{Type: "flt", Value: res}, nil
	}
	return nil, fmt.Errorf("I do not know how to do numbers/Sum for %v", e.Type)
}

func MaxBlock(v *vm.VM, e *vm.Elem) (*vm.Elem, error) {
	var res float64
	switch e.Type {
	case "iblock", "uiblock", "fblock":
		ar1 := vm.BlockToArray(e)
		res = floats.Max(ar1)
	}
	switch e.Type {
	case "iblock":
		return &vm.Elem{Type: "int", Value: int64(res)}, nil
	case "uiblock":
		return &vm.Elem{Type: "uint", Value: uint64(res)}, nil
	case "fblock":
		return &vm.Elem{Type: "int", Value: float64(res)}, nil
	}
	return nil, fmt.Errorf("I do not know how to do numbers/Max for %v", e.Type)
}

func MinBlock(v *vm.VM, e *vm.Elem) (*vm.Elem, error) {
	var res float64
	switch e.Type {
	case "iblock", "uiblock", "fblock":
		ar1 := vm.BlockToArray(e)
		res = floats.Min(ar1)
	}
	switch e.Type {
	case "iblock":
		return &vm.Elem{Type: "int", Value: int64(res)}, nil
	case "uiblock":
		return &vm.Elem{Type: "uint", Value: uint64(res)}, nil
	case "fblock":
		return &vm.Elem{Type: "int", Value: float64(res)}, nil
	}
	return nil, fmt.Errorf("I do not know how to do numbers/Max for %v", e.Type)
}

func NumSqrt(v *vm.VM, e *vm.Elem) (*vm.Elem, error) {
	switch e.Type {
	case "int":
		d := e.Value.(int64)
		return &vm.Elem{Type: "flt", Value: math.Sqrt(float64(d))}, nil
	case "uint":
		d := e.Value.(uint64)
		return &vm.Elem{Type: "flt", Value: math.Sqrt(float64(d))}, nil
	case "flt":
		d := e.Value.(float64)
		return &vm.Elem{Type: "flt", Value: math.Sqrt(d)}, nil
	case "big":
		res := big.NewInt(0)
		d := e.Value.(*big.Int)
		return &vm.Elem{Type: "big", Value: res.Sqrt(d)}, nil
	case "iblock", "uiblock", "fblock":
		res := vm.ApplyFloatFunc(e, math.Sqrt)
		if res != nil {
			return res, nil
		}
	}
	return nil, fmt.Errorf("I do not know how to do numbers/Sqrt for %v", e.Type)
}

func InitNumbersFunctions() {
	log.Debug("[ BUND ] bund.InitNumbersFunctions() reached")
	vm.AddFunction("numbers/Span", SpanFblock)
	vm.AddFunction("numbers/Logspan", LogSpanFblock)
	vm.AddFunction("numbers/Reverse", ReverseBlock)
	vm.AddFunction("numbers/Sum", SumBlock)
	vm.AddFunction("numbers/Max", MaxBlock)
	vm.AddFunction("numbers/Min", MinBlock)
	vm.AddFunction("numbers/Sqrt", NumSqrt)
}
