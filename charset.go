package c4

import (
	"errors"
	"strconv"
	"strings"
)

const (
	OldCharset = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
	Charset    = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

var (
	oldnewlut map[byte]byte
	newoldlut map[byte]byte
)

func init() {
	oldnewlut = make(map[byte]byte)
	newoldlut = make(map[byte]byte)
	for i := 0; i < base; i++ {
		oldnewlut[OldCharset[i]] = Charset[i]
		newoldlut[Charset[i]] = OldCharset[i]
	}
}

// Returns true if all digits are the same between two IDs.
func sameNumbers(a *ID, b *ID) bool {
	if a == nil || b == nil {
		return false
	}
	x := a.String()
	y := b.String()
	for i, c := range x {
		_, err := strconv.ParseInt(string(c), 10, 8)
		if err != nil {
			continue
		}
		if rune(y[i]) != c {
			return false
		}
	}
	return true
}

func CheckCharacterSet(a *ID, b *ID) (*ID, error) {
	if !sameNumbers(a, b) {
		return nil, errors.New("not the same id")
	}
	var x, y string
	remove_set := "123456789"
	x = a.String()[2:]
	y = b.String()[2:]
	for _, c := range remove_set {
		x = strings.Replace(x, string(c), "", -1)
		y = strings.Replace(y, string(c), "", -1)
	}
	newer := 0 // -1, 0, 1: a is newer, not the same, b is newer
	for i := range x {
		xc := x[i]
		yc := y[i]
		if oldnewlut[xc] == yc {
			if newer == -1 {
				return nil, errors.New("not the same id 2")
			}
			newer = 1
			continue
		}
		if newoldlut[xc] != yc || newer == 1 {
			return nil, errors.New("not the same id 2")
		}
		newer = -1
	}
	if newer == -1 {
		return a, nil
	}
	return b, nil
}

func OldCharsetIDToNew(id *ID) *ID {
	if id == nil {
		return nil
	}
	id_str := id.String()
	newid := "c4"

	for i := 2; i < 90; i++ {
		newid = newid + string(oldnewlut[id_str[i]])
	}
	idout, _ := ParseID(newid)
	return idout
}

func NewCharsetIDToOld(id *ID) *ID {
	if id == nil {
		return nil
	}
	id_str := id.String()
	oldid := "c4"

	for i := 2; i < 90; i++ {
		oldid = oldid + string(newoldlut[id_str[i]])
	}
	idout, _ := ParseID(oldid)
	return idout
}
