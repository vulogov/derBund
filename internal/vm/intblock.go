package vm

import (
	"github.com/gammazero/deque"
	"github.com/pieterclaerhout/go-log"
	"gonum.org/v1/gonum/floats"
)

func IblockToString(e *Elem) string {
	if e.Type == "iblock" {
		res := "(int "
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
	log.Errorf("trying to convert an Iblock and it is not an Iblock: %v %T", e.Type, e.Value)
	return "<?>"
}

func IblockFromString(d string) *Elem {
	res := new(deque.Deque)
	return &Elem{Type: "iblock", Value: &res}
}

func IblockCompare(e1 *Elem, e2 *Elem) int {
	ar1 := BlockToArray(e1)
	ar2 := BlockToArray(e2)
	if floats.Equal(ar1, ar2) == true {
		return Eq
	}
	return Ne
}

func IblockDup(e *Elem) *Elem {
	if e.Type != "iblock" {
		return nil
	}
	return &Elem{Type: "iblock", Value: e.Value}
}

func RegisterIblock() {
	RegisterType("iblock", IblockToString, IblockFromString, IblockCompare, IblockDup)
}
