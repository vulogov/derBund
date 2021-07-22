package vm

import (
	"github.com/gammazero/deque"
	"github.com/pieterclaerhout/go-log"
)

func FblockToString(e *Elem) string {
	if e.Type == "fblock" {
		res := "(float "
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
	log.Errorf("trying to convert an Fblock and it is not an Fblock: %v %T", e.Type, e.Value)
	return "<?>"
}

func FblockFromString(d string) *Elem {
	res := new(deque.Deque)
	return &Elem{Type: "fblock", Value: &res}
}

func FblockCompare(e1 *Elem, e2 *Elem) int {
	return IDK
}

func FblockDup(e *Elem) *Elem {
	if e.Type != "fblock" {
		return nil
	}
	return &Elem{Type: "fblock", Value: e.Value}
}

func RegisterFblock() {
	RegisterType("fblock", FblockToString, FblockFromString, FblockCompare, FblockDup)
}
