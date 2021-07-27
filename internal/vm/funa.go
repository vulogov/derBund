package vm

import (
	"fmt"

	"github.com/pieterclaerhout/go-log"
)

func AFunToString(e *Elem) string {
	if e.Type == "ACALL" {
		return fmt.Sprintf("[%v] ... .", e.Value.(string))
	}
	log.Errorf("trying to convert a Lambda and it is not an Lambda: %v %T", e.Type, e.Value)
	return "<?>"
}

func AFunFromString(d string) *Elem {
	return &Elem{Type: "ACALL", Value: d}
}

func AFunCompare(e1 *Elem, e2 *Elem) int {
	return IDK
}

func AFunDup(e *Elem) *Elem {
	if e.Type != "ACALL" {
		return nil
	}
	return &Elem{Type: "ACALL", Value: e.Value}
}

func RegisterFuna() {
	RegisterType("ACALL", AFunToString, AFunFromString, AFunCompare, AFunDup)
}
