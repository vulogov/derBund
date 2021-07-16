package vm

import (
	"fmt"

	"github.com/lrita/cmap"
	"github.com/pieterclaerhout/go-log"
)

type Elem struct {
	Type  string
	Value interface{}
}

const (
	Eq  = 0
	Gt  = 1
	Ls  = -1
	IDK = -2
)

type ToStringFun func(e *Elem) string
type FromStringFun func(d string) *Elem
type CompareFun func(e1 *Elem, e2 *Elem) int
type DupFun func(e *Elem) *Elem

type ElemHandler struct {
	Type       string
	ToString   ToStringFun
	FromString FromStringFun
	Compare    CompareFun
	Duplicate  DupFun
}

var BundTypes cmap.Cmap

func RegisterType(t string, ts ToStringFun, fs FromStringFun, cf CompareFun, df DupFun) bool {
	if _, ok := BundTypes.Load(t); ok {
		return true
	}
	BundTypes.Store(t, &ElemHandler{Type: t, ToString: ts, FromString: fs, Compare: cf})
	log.Debugf("Register BUND datatype: %v", t)
	return true
}

func GetType(t string) (*ElemHandler, error) {
	if res, ok := BundTypes.Load(t); ok {
		return res.(*ElemHandler), nil
	}
	return nil, fmt.Errorf("BUND do not have datatype: %v", t)
}
