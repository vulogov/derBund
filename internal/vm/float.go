package vm

import (
	"fmt"

	"github.com/cstockton/go-conv"
	"github.com/pieterclaerhout/go-log"
)

func FloatToString(e *Elem) string {
	if e.Type == "flt" {
		switch v := e.Value.(type) {
		case float64:
			return fmt.Sprint(v)
		}
	}
	log.Errorf("trying to convert an Float and it is not an Float: %v %T", e.Type, e.Value)
	return "<?>"
}

func FloatFromString(d string) *Elem {
	res, err := conv.Float64(d)
	if err != nil {
		return nil
	}
	return &Elem{Type: "flt", Value: res}
}

func asfloat(v interface{}) float64 {
	switch e := v.(type) {
	case int64:
		return float64(e)
	case float64:
		return e
	}
	return 0.0
}

func FloatCompare(e1 *Elem, e2 *Elem) int {
	var fe1, fe2 float64
	if e1.Type == "flt" || e1.Type == "int" {
		fe1 = asfloat(e1.Value)
	} else {
		return IDK
	}
	if e2.Type == "flt" || e2.Type == "int" {
		fe2 = asfloat(e2.Value)
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

func FloatDup(e *Elem) *Elem {
	if e.Type != "flt" {
		return nil
	}
	return &Elem{Type: "flt", Value: e.Value}
}

func RegisterFloat() {
	RegisterType("flt", FloatToString, FloatFromString, FloatCompare, FloatDup)
}
