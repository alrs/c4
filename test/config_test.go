package test_test

import (
	"testing"

	"github.com/etcenter/c4/test"

	"github.com/cheekybits/is"
)

func TestTestConfig(t *testing.T) {
	is := is.New(t)

	cfg := test.TestConfig(is)
	is.NotNil(cfg)
	defer test.TestDeleteConfig(cfg)
	is.NotNil(cfg.Db)
	is.NotNil(cfg.Root)
}
