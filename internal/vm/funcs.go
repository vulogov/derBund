package vm

import (
	"fmt"

	"github.com/lrita/cmap"
	"github.com/pieterclaerhout/go-log"
)

var BundFunctions cmap.Cmap
var BundOperators cmap.Cmap

type BundFunction func(vm *VM, e *Elem) (*Elem, error)
type BundOperator func(vm *VM, e1 *Elem, e2 *Elem) (*Elem, error)

func AddFunction(name string, fn BundFunction) bool {
	if _, ok := BundFunctions.Load(name); ok {
		log.Warnf("Function %v already registered", name)
		return true
	}
	BundFunctions.Store(name, fn)
	log.Debugf("Register BUND function: %v", name)
	return true
}

func AddOperator(name string, fn BundOperator) bool {
	if _, ok := BundOperators.Load(name); ok {
		log.Warnf("Operator %v already registered", name)
		return true
	}
	BundFunctions.Store(name, fn)
	log.Debugf("Register BUND operator: %v", name)
	return true
}

func HasUserFunction(name string, vm *VM) bool {
	return false
}

func GetFunction(name string) (BundFunction, error) {
	if res, ok := BundFunctions.Load(name); ok {
		return res.(BundFunction), nil
	}
	return nil, fmt.Errorf("BUND do not have function: %v", name)
}

func GetOperator(name string) (BundOperator, error) {
	if res, ok := BundOperators.Load(name); ok {
		return res.(BundOperator), nil
	}
	return nil, fmt.Errorf("BUND do not have operator: %v", name)
}
