// Code generated by "stringer -type=NodeType"; DO NOT EDIT.

package ast

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
}

const _NodeType_name = "Unknown"

var _NodeType_index = [...]uint8{0, 7}

func (i NodeType) String() string {
	if i >= NodeType(len(_NodeType_index)-1) {
		return "NodeType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _NodeType_name[_NodeType_index[i]:_NodeType_index[i+1]]
}
