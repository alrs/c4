package c4

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

type errAllocation string

func (e errAllocation) Error() string {
	return "memory allocation error in " + string(e)
}

type errKeyNotExist []byte

func (e errKeyNotExist) Error() string {
	return "key does not exist: " + string(e)
}
