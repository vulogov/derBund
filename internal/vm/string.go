package vm

import (
	"unsafe"

	"github.com/cstockton/go-conv"
	"github.com/pieterclaerhout/go-log"
)

func StringDeepCopy(s string) string {
	b := make([]byte, len(s))
	copy(b, s)
	return *(*string)(unsafe.Pointer(&b))
}

func StringToString(e *Elem) string {
	if e.Type == "str" {
		switch v := e.Value.(type) {
		case string:
			return string(v)
		}
	}
	log.Errorf("trying to convert a String and it is not a String: %v %T", e.Type, e.Value)
	return "<?>"
}

func StringFromString(d string) *Elem {
	res, err := conv.String(d)
	if err != nil {
		return nil
	}
	return &Elem{Type: "str", Value: res}
}

func StringCompare(e1 *Elem, e2 *Elem) int {
	if e1.Type == "str" && e2.Type == "str" {
		r1 := e1.Value.(string)
		r2 := e2.Value.(string)
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

func StringDup(e *Elem) *Elem {
	if e.Type != "str" {
		return nil
	}
	return &Elem{Type: "str", Value: StringDeepCopy(e.Value.(string))}
}

func RegisterString() {
	RegisterType("str", StringToString, StringFromString, StringCompare, StringDup)
}
