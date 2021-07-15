package bund

import (
	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/derBund/internal/banner"
)

func Fin() {
	banner.Banner("[ Zay Gezunt ]")
	log.Debug("[ BUND ] tsak.Fin() is reached")
}
