// generated by stringer -type=Move; DO NOT EDIT

package dep

import "fmt"

const _Move_name = "ShiftLeftRightMAXMOVE"

var _Move_index = [...]uint8{0, 5, 9, 14, 21}

func (i Move) String() string {
	if i >= Move(len(_Move_index)-1) {
		return fmt.Sprintf("Move(%d)", i)
	}
	return _Move_name[_Move_index[i]:_Move_index[i+1]]
}
