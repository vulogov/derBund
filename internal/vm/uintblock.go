package vm

import (
	"github.com/gammazero/deque"
	"github.com/pieterclaerhout/go-log"
	"gonum.org/v1/gonum/floats"
)

func UIblockToString(e *Elem) string {
	if e.Type == "uiblock" {
		res := "(uint "
		for i := 0; i < e.Value.(*deque.Deque).Len(); i++ {
			_e := e.Value.(*deque.Deque).At(i)
			eh, err := GetType(_e.(*Elem).Type)
			if err != nil {
				log.Errorf("Can not get type: %v", _e.(*Elem).Type)
				continue
			}
			res = res + eh.ToString(_e.(*Elem))
			res = res + " "
		}
		res = res + " )"
		return res
	}
	log.Errorf("trying to convert an UIblock and it is not an UIblock: %v %T", e.Type, e.Value)
	return "<?>"
}

func UIblockFromString(d string) *Elem {
	res := new(deque.Deque)
	return &Elem{Type: "uiblock", Value: &res}
}

func UIblockCompare(e1 *Elem, e2 *Elem) int {
	ar1 := BlockToArray(e1)
	ar2 := BlockToArray(e2)
	if floats.Equal(ar1, ar2) == true {
		return Eq
	}
	return Ne
}

func UIblockDup(e *Elem) *Elem {
	if e.Type != "uiblock" {
		return nil
	}
	return &Elem{Type: "uiblock", Value: e.Value}
}

func RegisterUIblock() {
	RegisterType("uiblock", UIblockToString, UIblockFromString, UIblockCompare, UIblockDup)
}
