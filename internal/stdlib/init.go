package stdlib

import (
	"github.com/pieterclaerhout/go-log"
)

func InitSTDLIB() {
	log.Debug("[ BUND ] bund.InitSTDLIB() reached")
	InitPrintFunctions()
}
