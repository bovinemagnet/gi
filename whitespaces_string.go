// Code generated by "stringer -type=WhiteSpaces"; DO NOT EDIT.

package gi

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

const _WhiteSpaces_name = "WhiteSpaceNormalWhiteSpaceNowrapWhiteSpacePreWhiteSpacePreLineWhiteSpacePreWrapWhiteSpacesN"

var _WhiteSpaces_index = [...]uint8{0, 16, 32, 45, 62, 79, 91}

func (i WhiteSpaces) String() string {
	if i < 0 || i >= WhiteSpaces(len(_WhiteSpaces_index)-1) {
		return "WhiteSpaces(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _WhiteSpaces_name[_WhiteSpaces_index[i]:_WhiteSpaces_index[i+1]]
}

func (i *WhiteSpaces) FromString(s string) error {
	for j := 0; j < len(_WhiteSpaces_index)-1; j++ {
		if s == _WhiteSpaces_name[_WhiteSpaces_index[j]:_WhiteSpaces_index[j+1]] {
			*i = WhiteSpaces(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: WhiteSpaces")
}
