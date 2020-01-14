// Code generated by "stringer -type=FileStatus"; DO NOT EDIT.

package vci

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Untracked-0]
	_ = x[Stored-1]
	_ = x[Modified-2]
	_ = x[Added-3]
	_ = x[Deleted-4]
	_ = x[Conflicted-5]
	_ = x[Updated-6]
	_ = x[FileStatusN-7]
}

const _FileStatus_name = "UntrackedStoredModifiedAddedDeletedConflictedUpdatedFileStatusN"

var _FileStatus_index = [...]uint8{0, 9, 15, 23, 28, 35, 45, 52, 63}

func (i FileStatus) String() string {
	if i < 0 || i >= FileStatus(len(_FileStatus_index)-1) {
		return "FileStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FileStatus_name[_FileStatus_index[i]:_FileStatus_index[i+1]]
}

func (i *FileStatus) FromString(s string) error {
	for j := 0; j < len(_FileStatus_index)-1; j++ {
		if s == _FileStatus_name[_FileStatus_index[j]:_FileStatus_index[j+1]] {
			*i = FileStatus(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: FileStatus")
}