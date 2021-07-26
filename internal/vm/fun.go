package vm

import (
	"fmt"

	"github.com/pieterclaerhout/go-log"
)

func FunToString(e *Elem) string {
	if e.Type == "CALL" {
		return fmt.Sprintf("[%v] ... .", e.Value.(string))
	}
	log.Errorf("trying to convert a Lambda and it is not an Lambda: %v %T", e.Type, e.Value)
	return "<?>"
}

func FunFromString(d string) *Elem {
	return &Elem{Type: "CALL", Value: d}
}

func FunCompare(e1 *Elem, e2 *Elem) int {
	return IDK
}

func FunDup(e *Elem) *Elem {
	if e.Type != "CALL" {
		return nil
	}
	return &Elem{Type: "CALL", Value: e.Value}
}

func RegisterFun() {
	RegisterType("CALL", FunToString, FunFromString, FunCompare, FunDup)
}
