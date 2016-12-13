package asset

import (
	"bytes"
	"fmt"
	"io"
	"math/big"
)

func (id *ID) RawBytes() (b []byte) {
	bignum := big.Int(*id)
	b_raw := (&bignum).Bytes()
	bytes64 := make([]byte, 64)

	padding := 64 - len(b_raw)
	// Can't use copy!
	// It doesn't properly handle leading zeros
	// copy(bytes64, b_raw)
	for i, bb := range b_raw {
		bytes64[padding+i] = bb
	}
	return bytes64[:]
}

func (i *ID) Sum(j *ID) (*ID, error) {
	left := i.RawBytes()
	right := j.RawBytes()

	if bytes.Compare(left, right) < 0 {
		left, right = right, left
	}
	e := NewIDEncoder()
	_, err := io.Copy(e, bytes.NewReader(append(left, right...)))
	// fmt.Printf("Sum: %d, %d, n: %d\n", left[63], right[63], n)
	return e.ID(), err
}

func Tree(ids IDSlice) chan *ID {
	out := make(chan *ID)

	//sort
	ids.Sort()
	fmt.Printf("\nSorted:\n")
	for _, id := range ids {
		fmt.Printf("%s\n", id)
	}
	temp_ids := []*ID(ids)
	var next_ids []*ID

	row := 1

	fmt.Printf("\n\nLoop: %d\n", row)

	go func() {
		//iterate
		i := 0
		for {
			if len(temp_ids) == 1 {
				out <- temp_ids[0]
				break
			}
			if i >= len(temp_ids) {
				temp_ids = make([]*ID, len(next_ids))
				copy(temp_ids, next_ids)
				if len(temp_ids) == 1 {
					out <- temp_ids[0]
					break
				}

			}
			id1 := temp_ids[i]
			i++
			fmt.Printf("\t%s\n", id1)

			if i >= len(temp_ids) {
				fmt.Printf("= %s\n", id1)

				next_ids = append(next_ids, id1)
				temp_ids = make([]*ID, len(next_ids))
				copy(temp_ids, next_ids)
				next_ids = make([]*ID, 0, len(temp_ids)/2+1)
				i = 0
				row += 1

				fmt.Printf("\n\nLoop: %d\n", row)
				continue
			}
			id2 := temp_ids[i]
			fmt.Printf("\t%s\n", id2)
			i++

			temp_id, err := id1.Sum(id2)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("= %s\n", temp_id)

			next_ids = append(next_ids, temp_id)
			// out <- temp_id
		}
		close(out)
	}()
	return out

}
