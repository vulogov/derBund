package vm

import (
	"github.com/gammazero/deque"
)

func BlockToArray(d *Elem) []float64 {
	var res []float64
	q := d.Value.(*deque.Deque)
	for i := 0; i < q.Len(); i++ {
		e := q.At(i).(*Elem)
		switch d.Type {
		case "iblock":
			res = append(res, float64(e.Value.(int64)))
		case "uiblock":
			res = append(res, float64(e.Value.(uint64)))
		case "fblock":
			res = append(res, float64(e.Value.(float64)))
		}
	}
	return res
}

func ArrayToBlock(t string, d []float64) *Elem {
	q := new(deque.Deque)
	res := Elem{Type: t, Value: q}
	for _, v := range d {
		switch t {
		case "iblock":
			q.PushBack(&Elem{Type: "int", Value: int64(v)})
		case "uiblock":
			q.PushBack(&Elem{Type: "uint", Value: uint64(v)})
		case "fblock":
			q.PushBack(&Elem{Type: "flt", Value: float64(v)})
		}
	}
	return &res
}

func MulConst(s float64, d []float64) {
	for n, v := range d {
		d[n] = v * s
	}
}

func DivConst(s float64, d []float64) {
	for n, v := range d {
		d[n] = v / s
	}
}

func SubConst(s float64, d []float64) {
	for n, v := range d {
		d[n] = v - s
	}
}
