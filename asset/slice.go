package asset

import (
	"bytes"
	"io"
	"sort"
)

type IDSlice []*ID

func (s IDSlice) Len() int           { return len(s) }
func (s IDSlice) Less(i, j int) bool { return s[i].Cmp(s[j]) < 0 }
func (s IDSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// Sort is a convenience method.
func (s IDSlice) Sort() {
	sort.Sort(s)
}

// Append id to slice.
func (s *IDSlice) Push(id *ID) {
	*s = append(*s, id)
}

type IDSliceIndex struct {
	I  int
	Id *ID
}

// Iterator over each id in an *IDSlice
func (s *IDSlice) Iterator(cancel <-chan struct{}) <-chan *IDSliceIndex {
	out := make(chan *IDSliceIndex)
	go func() {
		for i, id := range *s {
			newID := *id
			idx := IDSliceIndex{i, &newID}
			select {
			case out <- &idx:
			case <-cancel:
				break
			}
		}
		close(out)
	}()
	return out
}

func (s *IDSlice) Index(i int) *ID {
	ss := ([]*ID)(*s)
	return ss[i]
}

//String returns the slice of c4ids concatenated together without spaces or newlines.
func (s *IDSlice) String() string {
	str := ""
	for _, id := range *s {
		str += id.String()
	}
	return str
}

// SearchIDs searches for x in a sorted slice of *ID and returns the index
// as specified by sort.Search. The slice must be sorted in ascending order.
func SearchIDs(a IDSlice, x *ID) int {
	return sort.Search(len(a), func(i int) bool { return a[i].Cmp(x) >= 0 })
}

// ID of a sorted slice of IDs
func (s IDSlice) ID() (*ID, error) {
	s.Sort()
	encoder := NewIDEncoder()
	for _, bigID := range s {
		_, err := io.Copy(encoder, bytes.NewReader(bigID.Bytes()))
		if err != nil {
			return nil, err
		}
	}
	return encoder.ID(), nil
}

// ParseBytesIDSlice parses a C4 ID as []byte into an ID.
func ParseBytesIDSlice(src []byte) (*IDSlice, error) {
	l := len(src)
	if l == 0 || l%90 != 0 {
		return nil, idSliceErr("ParseBytesIDSlice input must be multiple of 90 characters long.")
	}
	var ids IDSlice
	end := 90
	for end <= l {
		c4id := src[end-90 : end]
		// fmt.Printf("parsing: %s\n", string(c4id))
		end += 90
		// id, err := Identify(bytes.NewReader(c4id))
		id, err := ParseID(string(c4id))
		if err != nil {
			return nil, err
		}
		ids.Push(id)
	}
	return &ids, nil
}
