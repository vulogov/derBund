pre:
	@echo "=== BUND === [Preinstallation of some stuff]"
	go get gotest.tools/gotestsum
	go get github.com/client9/misspell/cmd/misspell
	go get golang.org/x/tools/cmd/godoc
	go get github.com/llorllale/go-gitlint/cmd/go-gitlint
	go get github.com/psampaz/go-mod-outdated
	go get golang.org/x/tools/cmd/goimports
	go get gopkg.in/alecthomas/kingpin.v2
	go get gotest.tools/gotestsum
	go get github.com/stretchr/testify/assert@v1.6.1
	go get github.com/rs/xid
	go get github.com/lrita/cmap
	go get github.com/gammazero/deque
	go get github.com/common-nighthawk/go-figure
	go get -u gonum.org/v1/gonum/...
	go get github.com/pieterclaerhout/go-formatter@v1.0.4
	go get github.com/elliotchance/orderedmap
	go mod download github.com/stretchr/testify
	go get github.com/mgutz/ansi
	go get github.com/pieterclaerhout/go-log
	go get github.com/tomlazar/table
	go get github.com/antlr/antlr4/runtime/Go/antlr
	go get github.com/google/uuid
	go get github.com/cstockton/go-conv
