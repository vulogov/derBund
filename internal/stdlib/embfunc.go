package stdlib

import (
	"github.com/pieterclaerhout/go-log"
)

func InitFUNCTIONS() {
	log.Debug("[ BUND ] bund.InitFUNCTIONS() reached")
	InitPrintFunctions()
	InitNumbersFunctions()
	InitBigFunctions()
}
