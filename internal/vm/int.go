package vm

import (
	"fmt"

	"github.com/cstockton/go-conv"
	"github.com/pieterclaerhout/go-log"
)

func IntToString(e *Elem) string {
	if e.Type == "int" {
		switch v := e.Value.(type) {
		case int64:
			return fmt.Sprint(v)
		}
	}
	log.Errorf("trying to convert an Integer and it is not an Integer: %v %T", e.Type, e.Value)
	return "<?>"
}

func IntFromString(d string) *Elem {
	res, err := conv.Int64(d)
	if err != nil {
		return nil
	}
	return &Elem{Type: "int", Value: res}
}

func IntCompare(e1 *Elem, e2 *Elem) int {
	if e1.Type == "int" && e2.Type == "int" {
		r1 := e1.Value.(int64)
		r2 := e2.Value.(int64)
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

func IntDup(e *Elem) *Elem {
	if e.Type != "int" {
		return nil
	}
	return &Elem{Type: "int", Value: e.Value}
}

func RegisterInt() {
	RegisterType("int", IntToString, IntFromString, IntCompare, IntDup)
}
