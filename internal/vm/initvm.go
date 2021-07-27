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
	RegisterString()
	RegisterInt()
	RegisterUInt()
	RegisterFloat()
	RegisterUFloat()
	RegisterCpx()
	RegisterDblock()
	RegisterFblock()
	RegisterIblock()
	RegisterUIblock()
	RegisterFun()
	RegisterFuna()
}
