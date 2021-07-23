package vm

import (
	"fmt"

	"github.com/cstockton/go-conv"
	"github.com/pieterclaerhout/go-log"
)

func UIntToString(e *Elem) string {
	if e.Type == "uint" {
		switch v := e.Value.(type) {
		case uint64:
			return fmt.Sprint(v)
		}
	}
	log.Errorf("trying to convert an UInteger and it is not an UInteger: %v %T", e.Type, e.Value)
	return "<?>"
}

func UIntFromString(d string) *Elem {
	res, err := conv.Uint64(d)
	if err != nil {
		return nil
	}
	return &Elem{Type: "uint", Value: res}
}

func UIntCompare(e1 *Elem, e2 *Elem) int {
	if e1.Type == "uint" && e2.Type == "uint" {
		r1 := e1.Value.(uint64)
		r2 := e2.Value.(uint64)
		if r1 == r2 {
			return Eq
		} else if r1 > r2 {
			return Gt
		} else {
			return Ls
		}
	}
	return IDK
}

func UIntDup(e *Elem) *Elem {
	if e.Type != "uint" {
		return nil
	}
	return &Elem{Type: "uint", Value: e.Value}
}

func RegisterUInt() {
	RegisterType("uint", UIntToString, UIntFromString, UIntCompare, UIntDup)
}
