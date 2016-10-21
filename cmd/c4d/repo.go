package main

import (
	"fmt"
)

var currentId int

var assets Assets

// Give us some seed data
func init() {
	RepoCreateAsset(Asset{Name: "Write presentation"})
	RepoCreateAsset(Asset{Name: "Host meetup"})
}

func RepoFindAsset(id int) Asset {
	for _, t := range assets {
		if t.Id == id {
			return t
		}
	}
	return Asset{}
}

func RepoCreateAsset(t Asset) Asset {
	currentId += 1
	t.Id = currentId
	assets = append(assets, t)
	return t
}

func RepoDestroyAsset(id int) error {
	for i, t := range assets {
		if t.Id == id {
			assets = append(assets[:i], assets[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Asset with id of %d to delete", id)
}
