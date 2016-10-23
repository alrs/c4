package asset_test

import (
	// "bytes"

	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"testing"

	"github.com/cheekybits/is"
	"github.com/etcenter/c4/asset"
)

func TestIDMarshalJSON(t *testing.T) {
	is := is.New(t)

	type testType struct {
		Name string    `json:"name"`
		ID   *asset.ID `json:"id"`
	}

	big_empty := big.NewInt(0)
	for _, test := range []struct {
		In  testType
		Exp string
	}{
		{
			In:  testType{"Test", asset.NIL_ID},
			Exp: `{"name":"Test","id":"c459CSJESBh38BxDwwxNFKTXE4cC9HASGe3bhtN6z58GbwLqpCyRaKyZSvBAvTdF5NpSTPdUMH4hHRJ75geLsB1Sfs"}`,
		},
		{
			In:  testType{"Test", nil},
			Exp: `{"name":"Test","id":null}`,
		},
		{
			In:  testType{"Test", (*asset.ID)(big_empty)},
			Exp: `{"name":"Test","id":""}`,
		},
	} {
		actual, err := json.Marshal(test.In)
		is.NoErr(err)
		is.Equal(string(actual), test.Exp)
	}
}

func TestIDUnarshalJSON(t *testing.T) {
	is := is.New(t)

	type testType struct {
		Name string    `json:"name"`
		ID   *asset.ID `json:"id"`
	}

	for _, test := range []struct {
		In  string
		Exp testType
	}{
		{
			In:  `{"name":"Test","id":"c459CSJESBh38BxDwwxNFKTXE4cC9HASGe3bhtN6z58GbwLqpCyRaKyZSvBAvTdF5NpSTPdUMH4hHRJ75geLsB1Sfs"}`,
			Exp: testType{"Test", asset.NIL_ID},
		},
		{
			In:  `{"name":"Test","id":null}`,
			Exp: testType{"Test", nil},
		},
	} {

		testObject := testType{}
		err := json.Unmarshal([]byte(test.In), &testObject)
		is.NoErr(err)

		is.Equal(testObject, test.Exp)
	}
}

// Testing for IDSlice

func TestIDSliceMarshalJSON(t *testing.T) {
	is := is.New(t)

	type testType struct {
		Name string         `json:"name"`
		Ids  *asset.IDSlice `json:"ids"`
	}

	var someIDs asset.IDSlice
	testString := ""
	for i := 0; i < 10; i++ {
		num_str := strconv.Itoa(i)
		id, err := asset.Identify(strings.NewReader(num_str))
		is.NoErr(err)
		someIDs.Push(id)
	}
	someIDs.Sort()
	for idx := range someIDs.Iterator(nil) {
		testString += idx.Id.String()
	}

	emptyIDSlice := asset.IDSlice{}
	for _, test := range []struct {
		In  testType
		Exp string
	}{
		{
			In:  testType{"Test", &someIDs},
			Exp: fmt.Sprintf(`{"name":"Test","ids":"%s"}`, testString),
		},
		{
			In:  testType{"Test", nil},
			Exp: `{"name":"Test","ids":null}`,
		},
		{
			In:  testType{"Test", &emptyIDSlice},
			Exp: `{"name":"Test","ids":""}`,
		},
	} {
		actual, err := json.Marshal(test.In)
		is.NoErr(err)
		is.Equal(string(actual), test.Exp)
	}
}

func TestIDSliceUnarshalJSON(t *testing.T) {
	is := is.New(t)

	type testType struct {
		Name string         `json:"name"`
		Ids  *asset.IDSlice `json:"ids"`
	}

	var someIDs asset.IDSlice
	testString := ""
	for i := 0; i < 10; i++ {
		num_str := strconv.Itoa(i)
		id, err := asset.Identify(strings.NewReader(num_str))
		is.NoErr(err)
		someIDs.Push(id)
	}
	someIDs.Sort()
	for idx := range someIDs.Iterator(nil) {
		testString += idx.Id.String()
	}

	for _, test := range []struct {
		In  string
		Exp testType
	}{
		{
			In:  fmt.Sprintf(`{"name":"Test","ids":"%s"}`, testString),
			Exp: testType{"Test", &someIDs},
		},
		{
			In:  `{"name":"Test","ids":null}`,
			Exp: testType{"Test", nil},
		},
	} {
		testObject := testType{}
		err := json.Unmarshal([]byte(test.In), &testObject)
		is.NoErr(err)
		if test.Exp.Ids == nil {
			is.Equal(testObject.Ids, test.Exp.Ids)
		} else {
			is.Equal(*testObject.Ids, *test.Exp.Ids)
		}
	}
}
