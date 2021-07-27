package stdlib

import (
	"fmt"

	glob "github.com/ganbarodigital/go_glob"
	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/derBund/internal/vm"
)

func GlobPatternMatchOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type != "glob" {
		return nil, fmt.Errorf("Pattern for '===' not found")
	}
	if e2.Type != "str" {
		return nil, fmt.Errorf("Pattern matching with '===' done only with strings")
	}
	eh, err := vm.GetType(e1.Type)
	if err != nil {
		return nil, err
	}
	res := eh.Compare(e1, e2)
	if res == vm.Eq {
		return &vm.Elem{Type: "bool", Value: true}, nil
	} else {
		return &vm.Elem{Type: "bool", Value: false}, nil
	}
}

func GlobPatternMatchPrefixOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type != "glob" {
		return nil, fmt.Errorf("Pattern for '<===' not found")
	}
	if e2.Type != "str" {
		return nil, fmt.Errorf("Pattern matching with '<===' done only with strings")
	}
	g := glob.NewGlob(e1.Value.(string))
	_, res, err := g.MatchShortestPrefix(e2.Value.(string))
	if err != nil {
		return nil, err
	}
	if res == true {
		return &vm.Elem{Type: "bool", Value: true}, nil
	} else {
		return &vm.Elem{Type: "bool", Value: false}, nil
	}
}

func GlobPatternMatchSuffixOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type != "glob" {
		return nil, fmt.Errorf("Pattern for '===>' not found")
	}
	if e2.Type != "str" {
		return nil, fmt.Errorf("Pattern matching with '===>' done only with strings")
	}
	g := glob.NewGlob(e1.Value.(string))
	_, res, err := g.MatchShortestSuffix(e2.Value.(string))
	if err != nil {
		return nil, err
	}
	if res == true {
		return &vm.Elem{Type: "bool", Value: true}, nil
	} else {
		return &vm.Elem{Type: "bool", Value: false}, nil
	}
}

func GlobPatternMatchLongestPrefixOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type != "glob" {
		return nil, fmt.Errorf("Pattern for '=<==' not found")
	}
	if e2.Type != "str" {
		return nil, fmt.Errorf("Pattern matching with '=<==' done only with strings")
	}
	g := glob.NewGlob(e1.Value.(string))
	_, res, err := g.MatchLongestPrefix(e2.Value.(string))
	if err != nil {
		return nil, err
	}
	if res == true {
		return &vm.Elem{Type: "bool", Value: true}, nil
	} else {
		return &vm.Elem{Type: "bool", Value: false}, nil
	}
}

func GlobPatternMatchLongestSuffixOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type != "glob" {
		return nil, fmt.Errorf("Pattern for '==>=' not found")
	}
	if e2.Type != "str" {
		return nil, fmt.Errorf("Pattern matching with '==>=' done only with strings")
	}
	g := glob.NewGlob(e1.Value.(string))
	_, res, err := g.MatchLongestSuffix(e2.Value.(string))
	if err != nil {
		return nil, err
	}
	if res == true {
		return &vm.Elem{Type: "bool", Value: true}, nil
	} else {
		return &vm.Elem{Type: "bool", Value: false}, nil
	}
}

func GlobPatternMatchPrefixPosOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type != "glob" {
		return nil, fmt.Errorf("Pattern for '<===&' not found")
	}
	if e2.Type != "str" {
		return nil, fmt.Errorf("Pattern matching with '<===&' done only with strings")
	}
	g := glob.NewGlob(e1.Value.(string))
	pos, res, err := g.MatchShortestPrefix(e2.Value.(string))
	if err != nil {
		return nil, err
	}
	if res == true {
		return &vm.Elem{Type: "int", Value: int64(pos)}, nil
	} else {
		return &vm.Elem{Type: "int", Value: int64(-1)}, nil
	}
}

func GlobPatternMatchLongestPrefixPosOperator(v *vm.VM, e1 *vm.Elem, e2 *vm.Elem) (*vm.Elem, error) {
	if e1.Type != "glob" {
		return nil, fmt.Errorf("Pattern for '=<==&' not found")
	}
	if e2.Type != "str" {
		return nil, fmt.Errorf("Pattern matching with '=<==&' done only with strings")
	}
	g := glob.NewGlob(e1.Value.(string))
	pos, res, err := g.MatchLongestPrefix(e2.Value.(string))
	if err != nil {
		return nil, err
	}
	if res == true {
		return &vm.Elem{Type: "int", Value: int64(pos)}, nil
	} else {
		return &vm.Elem{Type: "int", Value: int64(-1)}, nil
	}
}

func InitGPMOperators() {
	log.Debug("[ BUND ] bund.InitGPMOperators() reached")
	vm.AddOperator("===", GlobPatternMatchOperator)
	vm.AddOperator("<===", GlobPatternMatchPrefixOperator)
	vm.AddOperator("===>", GlobPatternMatchSuffixOperator)
	vm.AddOperator("<===&", GlobPatternMatchPrefixPosOperator)
	vm.AddOperator("=<==", GlobPatternMatchLongestPrefixOperator)
	vm.AddOperator("==>=", GlobPatternMatchLongestSuffixOperator)
	vm.AddOperator("=<==&", GlobPatternMatchLongestPrefixPosOperator)

}
