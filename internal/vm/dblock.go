package vm

import (
	"github.com/gammazero/deque"
	"github.com/pieterclaerhout/go-log"
)

func DblockToString(e *Elem) string {
	if e.Type == "dblock" {
		res := "(* "
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
	log.Errorf("trying to convert an Dblock and it is not an Dblock: %v %T", e.Type, e.Value)
	return "<?>"
}

func DblockFromString(d string) *Elem {
	res := new(deque.Deque)
	return &Elem{Type: "dblock", Value: &res}
}

func DblockCompare(e1 *Elem, e2 *Elem) int {
	return IDK
}

func DblockDup(e *Elem) *Elem {
	if e.Type != "dblock" {
		return nil
	}
	return &Elem{Type: "dblock", Value: e.Value}
}

func RegisterDblock() {
	RegisterType("dblock", DblockToString, DblockFromString, DblockCompare, DblockDup)
}
