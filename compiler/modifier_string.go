// Code generated by "stringer -type=modifier"; DO NOT EDIT.

package compiler

import "fmt"

const (
	_modifier_name_0 = "LCTRLLSHIFT"
	_modifier_name_1 = "LALT"
	_modifier_name_2 = "LMETA"
	_modifier_name_3 = "RCTRL"
	_modifier_name_4 = "RSHIFT"
	_modifier_name_5 = "RALT"
	_modifier_name_6 = "RMETA"
	_modifier_name_7 = "NOMOD"
)

var (
	_modifier_index_0 = [...]uint8{0, 5, 11}
	_modifier_index_1 = [...]uint8{0, 4}
	_modifier_index_2 = [...]uint8{0, 5}
	_modifier_index_3 = [...]uint8{0, 5}
	_modifier_index_4 = [...]uint8{0, 6}
	_modifier_index_5 = [...]uint8{0, 4}
	_modifier_index_6 = [...]uint8{0, 5}
	_modifier_index_7 = [...]uint8{0, 5}
)

func (i modifier) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _modifier_name_0[_modifier_index_0[i]:_modifier_index_0[i+1]]
	case i == 4:
		return _modifier_name_1
	case i == 8:
		return _modifier_name_2
	case i == 16:
		return _modifier_name_3
	case i == 32:
		return _modifier_name_4
	case i == 64:
		return _modifier_name_5
	case i == 128:
		return _modifier_name_6
	case i == 256:
		return _modifier_name_7
	default:
		return fmt.Sprintf("modifier(%d)", i)
	}
}
