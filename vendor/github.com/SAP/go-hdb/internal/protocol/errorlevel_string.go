// Code generated by "stringer -type=ErrorLevel"; DO NOT EDIT.

package protocol

import "strconv"

const _ErrorLevel_name = "HdbWarningHdbErrorHdbFatalError"

var _ErrorLevel_index = [...]uint8{0, 10, 18, 31}

func (i ErrorLevel) String() string {
	if i < 0 || i >= ErrorLevel(len(_ErrorLevel_index)-1) {
		return "ErrorLevel(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ErrorLevel_name[_ErrorLevel_index[i]:_ErrorLevel_index[i+1]]
}
