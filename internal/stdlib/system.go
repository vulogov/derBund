package stdlib

import (
	"fmt"

	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/derBund/internal/vm"
)

func SwapOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e2.Type == "NONE" {
		return nil, fmt.Errorf("Stack is to shallow for <->")
	}
	v.Put(e1)
	return e2, nil
}

func InitSystemOperators() {
	log.Debug("[ BUND ] bund.InitSystemOperators() reached")
	vm.AddOperator("<->", SwapOperator)
}
