package main

type Asset struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Assets []Asset
