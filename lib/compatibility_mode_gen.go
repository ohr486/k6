// Code generated by "enumer -type=CompatibilityMode -transform=snake -trimprefix CompatibilityMode -output compatibility_mode_gen.go"; DO NOT EDIT.

//
package lib

import (
	"fmt"
)

const _CompatibilityModeName = "extendedbaseenhanced"

var _CompatibilityModeIndex = [...]uint8{0, 8, 12, 20}

func (i CompatibilityMode) String() string {
	i -= 1
	if i >= CompatibilityMode(len(_CompatibilityModeIndex)-1) {
		return fmt.Sprintf("CompatibilityMode(%d)", i+1)
	}
	return _CompatibilityModeName[_CompatibilityModeIndex[i]:_CompatibilityModeIndex[i+1]]
}

var _CompatibilityModeValues = []CompatibilityMode{1, 2, 3}

var _CompatibilityModeNameToValueMap = map[string]CompatibilityMode{
	_CompatibilityModeName[0:8]:   1,
	_CompatibilityModeName[8:12]:  2,
	_CompatibilityModeName[12:20]: 3,
}

// CompatibilityModeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CompatibilityModeString(s string) (CompatibilityMode, error) {
	if val, ok := _CompatibilityModeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to CompatibilityMode values", s)
}

// CompatibilityModeValues returns all values of the enum
func CompatibilityModeValues() []CompatibilityMode {
	return _CompatibilityModeValues
}

// IsACompatibilityMode returns "true" if the value is listed in the enum definition. "false" otherwise
func (i CompatibilityMode) IsACompatibilityMode() bool {
	for _, v := range _CompatibilityModeValues {
		if i == v {
			return true
		}
	}
	return false
}
