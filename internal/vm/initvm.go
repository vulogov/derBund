package vm

import (
	"github.com/pieterclaerhout/go-log"
)

func InitVM() {
	log.Debug("[ BUND ] Initialization")
	RegisterDataTypes()
}

func RegisterDataTypes() {
	log.Debug("[ BUND ] Registering internal data types")
	RegisterBoolean()
}
