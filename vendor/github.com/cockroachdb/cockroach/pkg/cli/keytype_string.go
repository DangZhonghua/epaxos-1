// Code generated by "stringer -type=keyType"; DO NOT EDIT.

package cli

import "fmt"

const _keyType_name = "rawhumanrangeID"

var _keyType_index = [...]uint8{0, 3, 8, 15}

func (i keyType) String() string {
	if i < 0 || i >= keyType(len(_keyType_index)-1) {
		return fmt.Sprintf("keyType(%d)", i)
	}
	return _keyType_name[_keyType_index[i]:_keyType_index[i+1]]
}
