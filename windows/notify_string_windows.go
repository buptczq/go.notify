// Code generated by "stringer -type IconType -output notify_string_windows.go"; DO NOT EDIT.

package windows

import "fmt"

const _IconType_name = "IconNoneIconInfoIconWarnIconErrorIconUser"

var _IconType_index = [...]uint8{0, 8, 16, 24, 33, 41}

func (i IconType) String() string {
	if i >= IconType(len(_IconType_index)-1) {
		return fmt.Sprintf("IconType(%d)", i)
	}
	return _IconType_name[_IconType_index[i]:_IconType_index[i+1]]
}
