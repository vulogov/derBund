package vm

import (
	"github.com/gammazero/deque"
	"github.com/pieterclaerhout/go-log"
)

func CblockMatch(q *deque.Deque, e *Elem) bool {
	for i := 0; i < q.Len(); i++ {
		val := q.At(i).(*Elem)
		switch val.Type {
		case "int", "uint", "uflt", "str", "bool":
			if val.Type != e.Type {
				continue
			}
			eh, err := GetType(val.Type)
			if err != nil {
				log.Errorf("Matching error: %v", err)
				return false
			}
			if eh.Compare(val, e) == Eq {
				return true
			}
		case "glob":
			if e.Type == "str" {
				eh, err := GetType("glob")
				if err != nil {
					log.Errorf("Matching error: %v", err)
					return false
				}
				if eh.Compare(val, e) == Eq {
					return true
				}
			}
		default:
			return false
		}
	}
	return false
}
