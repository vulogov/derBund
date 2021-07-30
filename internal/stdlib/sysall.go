package stdlib

import (
	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/derBund/internal/vm"
)

func AllSys(vm *vm.VM, n string) error {
	return nil
}

func InitAllSys() {
	log.Debug("[ BUND ] bund.InitAllSys() reached")
	vm.AddSys("all", AllSys)
}
