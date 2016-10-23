package assets

import "github.com/etcenter/c4/asset"

type ListAssetsInput struct {
	Marker *asset.ID `json:"marker"`
	MaxIDs *uint32   `json:"maxids"`
}

type ListAssetsOutput struct {
	Ids *asset.IDSlice `json:"ids"`
}
