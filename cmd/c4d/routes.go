package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"AssetsIndex",
		"GET",
		"/assets",
		AssetsIndex,
	},
	Route{
		"AssetShow",
		"GET",
		"/assets/{ID}",
		AssetShow,
	},
	Route{
		"AssetCreate",
		"POST",
		"/assets",
		AssetCreate,
	},
}
