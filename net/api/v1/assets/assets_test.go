package assets_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/Avalanche-io/c4/api/v1/assets"
	"github.com/etcenter/c4/asset"
	"github.com/etcenter/c4/env"
	c4_test "github.com/etcenter/c4/test"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"
	"github.com/cheekybits/is"
)

func TestListAssets(t *testing.T) {
	is := is.New(t)

	cfg := c4_test.TestConfig(is)
	is.NotNil(cfg)
	defer c4_test.TestDeleteConfig(cfg)
	env.SetGlobalConfig(cfg)

	err := cfg.Db.CreateBucket("assets")
	is.NoErr(err)

	expectedIds := asset.IDSlice{}
	for i := 0; i < 20; i++ {
		num_str := strconv.Itoa(i)
		id, err := asset.Identify(strings.NewReader(num_str))
		is.NoErr(err)
		cfg.Db.Put("assets", id.RawBytes(), nil)
		expectedIds.Push(id)
	}
	maxids := uint32(20)
	args := assets.ListAssetsInput{Marker: nil, MaxIDs: &maxids}

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Post("/assets", assets.ListAssets),
	)
	is.NoErr(err)
	api.SetApp(router)

	recorded := test.RunRequest(t, api.MakeHandler(), test.MakeSimpleRequest("POST", "http://1.2.3.4/assets", args))
	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
	expectedIds.Sort()
	resp := assets.ListAssetsOutput{}
	err = recorded.DecodeJsonPayload(&resp)
	is.NoErr(err)
	is.NotNil(resp.Ids)
	is.Equal(resp.Ids.Len(), 20)

	//  Test starting from a Marker (the 9th id)

	maxids = uint32(10)
	args2 := assets.ListAssetsInput{Marker: expectedIds.Index(9), MaxIDs: &maxids}
	recorded2 := test.RunRequest(t, api.MakeHandler(), test.MakeSimpleRequest("POST", "http://1.2.3.4/assets", args2))
	recorded2.CodeIs(200)
	recorded2.ContentTypeIsJson()
	resp2 := assets.ListAssetsOutput{}
	err = recorded2.DecodeJsonPayload(&resp2)
	is.NoErr(err)
	is.NotNil(resp2.Ids)
	is.Equal(resp2.Ids.Len(), 10)

	for idx := range resp2.Ids.Iterator(nil) {
		is.Equal(idx.Id.String(), expectedIds.Index(idx.I+9).String())
	}
}
