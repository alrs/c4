package assets

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/etcenter/c4/asset"
	"github.com/etcenter/c4/env"
)

const (
	MaxItems uint32 = 1000
)

func ListAssets(w rest.ResponseWriter, r *rest.Request) {
	params := ListAssetsInput{}
	err := r.DecodeJsonPayload(&params)
	if err != nil {
		panic(err)
	}
	app := env.GetGlobalConfig()
	i := uint32(0)
	limit := MaxItems
	if *params.MaxIDs < limit {
		limit = *params.MaxIDs
	}
	var marker []byte
	if params.Marker != nil {
		marker = params.Marker.RawBytes()
	}
	ids := asset.IDSlice{}
	cancel := make(chan struct{})
	_ = marker

	for ele := range app.Db.Iterator("assets", marker, cancel) {
		i++
		if i > limit {
			close(cancel)
			break
		}
		id := asset.BytesToID(ele.Key())
		ids.Push(id)
	}
	response := ListAssetsOutput{&ids}
	w.WriteJson(&response)
}
