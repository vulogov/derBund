package bund

import (
	"github.com/pieterclaerhout/go-log"

	// "github.com/vulogov/derBund/internal/conf"
	tlog "github.com/vulogov/derBund/internal/log"
	"github.com/vulogov/derBund/internal/signal"
	"github.com/vulogov/derBund/internal/vm"
)

func Init() {
	tlog.Init()
	log.Debug("[ BUND ] tsak.Init() is reached")
	signal.InitSignal()
	vm.InitVM()
}
