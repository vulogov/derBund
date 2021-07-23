package vm

import (
	"fmt"

	"strconv"

	"github.com/pieterclaerhout/go-log"
)

func ComplexToString(e *Elem) string {
	if e.Type == "cpx" {
		switch v := e.Value.(type) {
		case complex128:
			return fmt.Sprint(v)
		}
	}
	log.Errorf("trying to convert an Complex and it is not an Complex: %v %T", e.Type, e.Value)
	return "<?>"
}

func ComplexFromString(d string) *Elem {
	res, err := strconv.ParseComplex(d, 128)
	if err != nil {
		return nil
	}
	return &Elem{Type: "cpx", Value: res}
}

func ComplexCompare(e1 *Elem, e2 *Elem) int {
	var fe1, fe2 complex128
	if e1.Type == "cpx" {
		fe1 = e1.Value.(complex128)
	} else {
		return IDK
	}
	if e2.Type == "cpx" {
		fe2 = e2.Value.(complex128)
	} else {
		return IDK
	}
	if fe1 == fe2 {
		return Eq
	} else {
		return IDK
	}
}

func ComplexDup(e *Elem) *Elem {
	if e.Type != "cpx" {
		return nil
	}
	return &Elem{Type: "cpx", Value: e.Value}
}

func RegisterCpx() {
	RegisterType("cpx", ComplexToString, ComplexFromString, ComplexCompare, ComplexDup)
}
