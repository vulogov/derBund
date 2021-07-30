package vm

import (
	"github.com/pieterclaerhout/go-log"
)

func SepToString(e *Elem) string {
	if e.Type == "SEPARATE" {
		return "|"
	}
	log.Errorf("trying to convert a SEPARATE and it is not an SEPARATE: %v %T", e.Type, e.Value)
	return "<?>"
}

func SepFromString(d string) *Elem {
	return &Elem{Type: "SEPARATE", Value: nil}
}

func SepCompare(e1 *Elem, e2 *Elem) int {
	if e1.Type == "SEPARATE" && e2.Type == "SEPARATE" {
		return Eq
	}
	return IDK
}

func SepDup(e *Elem) *Elem {
	return &Elem{Type: "SEPARATE", Value: nil}
}

func RegisterSep() {
	RegisterType("SEPARATE", SepToString, SepFromString, SepCompare, SepDup)
}
