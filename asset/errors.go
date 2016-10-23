package asset

import (
	"strconv"
)

type errBadChar int

func (e errBadChar) Error() string {
	return "non c4 id character at position " + strconv.Itoa(int(e))
}

type errBadLength int

func (e errBadLength) Error() string {
	return "c4 ids must be 90 characters long, input length " + strconv.Itoa(int(e))
}

type idSliceErr string

func (e idSliceErr) Error() string {
	return "C4 IDSlice error: " + string(e)
}
