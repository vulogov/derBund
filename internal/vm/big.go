package vm

import (
	"fmt"

	"math/big"

	"github.com/pieterclaerhout/go-log"
)

func BigToString(e *Elem) string {
	if e.Type == "big" {
		i := e.Value.(*big.Int)
		return fmt.Sprintf("b%v", i.String())
	}
	log.Errorf("trying to convert an Big and it is not a Big: %v %T", e.Type, e.Value)
	return "<?>"
}

func BigFromString(d string) *Elem {
	i := big.NewInt(0)
	i.SetString(d, 0)
	res := Elem{Type: "big", Value: i}
	return &res
}

func BigCompare(e1 *Elem, e2 *Elem) int {
	if e1.Type == "big" && e2.Type == "big" {
		r1 := e1.Value.(*big.Int)
		r2 := e2.Value.(*big.Int)
		res := r1.Cmp(r2)
		if res == 0 {
			return Eq
		} else if res == -1 {
			return Ls
		} else if res == 1 {
			return Gt
		}
	}
	return IDK
}

func BigDup(e *Elem) *Elem {
	if e.Type != "big" {
		return nil
	}
	return BigFromString(BigToString(e))
}

func RegisterBig() {
	RegisterType("big", BigToString, BigFromString, BigCompare, BigDup)
}
