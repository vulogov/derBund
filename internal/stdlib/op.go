package stdlib

import (
	"github.com/pieterclaerhout/go-log"
)

func InitEmbeddedOperators() {
	log.Debug("[ BUND ] bund.InitEmbeddedOperators() reached")
	InitCMPOperators()
	InitGPMOperators()
	InitMathOperators()
}
