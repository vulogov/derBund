package vm

import (
	"fmt"

	"github.com/cstockton/go-conv"
	"github.com/pieterclaerhout/go-log"
)

func UFloatToString(e *Elem) string {
	if e.Type == "uflt" {
		switch v := e.Value.(type) {
		case float64:
			return fmt.Sprint(v)
		}
	}
	log.Errorf("trying to convert an UFloat and it is not an UFloat: %v %T", e.Type, e.Value)
	return "<?>"
}

func UFloatFromString(d string) *Elem {
	res, err := conv.Float64(d)
	if err != nil {
		return nil
	}
	return &Elem{Type: "uflt", Value: res}
}

func asUfloat(v interface{}) float64 {
	switch e := v.(type) {
	case uint64:
		return float64(e)
	case float64:
		return e
	}
	return 0.0
}

func UFloatCompare(e1 *Elem, e2 *Elem) int {
	var fe1, fe2 float64
	if e1.Type == "uflt" || e1.Type == "uint" {
		fe1 = asUfloat(e1.Value)
	} else {
		return IDK
	}
	if e2.Type == "uflt" || e2.Type == "uint" {
		fe2 = asUfloat(e2.Value)
	} else {
		return IDK
	}
	if fe1 == fe2 {
		return Eq
	} else if fe1 > fe2 {
		return Gt
	} else {
		return Ls
	}
	return IDK
}

func UFloatDup(e *Elem) *Elem {
	if e.Type != "uflt" {
		return nil
	}
	return &Elem{Type: "uflt", Value: e.Value}
}

func RegisterUFloat() {
	RegisterType("uflt", UFloatToString, UFloatFromString, UFloatCompare, UFloatDup)
}
