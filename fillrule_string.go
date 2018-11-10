// Code generated by "stringer -type=FillRule"; DO NOT EDIT.

package gi

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

const _FillRule_name = "FillRuleNonZeroFillRuleEvenOddFillRuleN"

var _FillRule_index = [...]uint8{0, 15, 30, 39}

func (i FillRule) String() string {
	if i < 0 || i >= FillRule(len(_FillRule_index)-1) {
		return "FillRule(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FillRule_name[_FillRule_index[i]:_FillRule_index[i+1]]
}

func (i *FillRule) FromString(s string) error {
	for j := 0; j < len(_FillRule_index)-1; j++ {
		if s == _FillRule_name[_FillRule_index[j]:_FillRule_index[j+1]] {
			*i = FillRule(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: FillRule")
}
