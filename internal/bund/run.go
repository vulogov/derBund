package bund

import (
	"github.com/pieterclaerhout/go-log"

	// "github.com/vulogov/derBund/internal/conf"
	"github.com/vulogov/derBund/internal/signal"
)

func Run() {
	Init()
	log.Debug("[ BUND ] tsak.Run() is reached")
	signal.ExitRequest()
}
