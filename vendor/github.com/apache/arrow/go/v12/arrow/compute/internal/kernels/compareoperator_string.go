// Code generated by "stringer -type=CompareOperator -linecomment"; DO NOT EDIT.

//go:build go1.18

package kernels

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CmpEQ-0]
	_ = x[CmpNE-1]
	_ = x[CmpGT-2]
	_ = x[CmpGE-3]
	_ = x[CmpLT-4]
	_ = x[CmpLE-5]
}

const _CompareOperator_name = "equalnot_equalgreatergreater_equallessless_equal"

var _CompareOperator_index = [...]uint8{0, 5, 14, 21, 34, 38, 48}

func (i CompareOperator) String() string {
	if i < 0 || i >= CompareOperator(len(_CompareOperator_index)-1) {
		return "CompareOperator(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CompareOperator_name[_CompareOperator_index[i]:_CompareOperator_index[i+1]]
}
