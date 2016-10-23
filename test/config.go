package test

import (
	"io/ioutil"
	"os"

	"github.com/cheekybits/is"
	"github.com/etcenter/c4/db"
	"github.com/etcenter/c4/env"
)

func TempDir(is is.I) string {
	dir, err := ioutil.TempDir("/tmp", "c4test_")
	is.NoErr(err)
	return dir
}

func DeleteDir(dir *string) {
	os.RemoveAll(*dir)
}

func TestConfig(is is.I) *env.Config {
	dir := TempDir(is)
	temp_db := dir + "/c4.db"
	test_db, err := db.Open(temp_db)
	is.NoErr(err)
	return env.NewConfig().WithRoot(dir).WithDb(test_db)
}

func TestDeleteConfig(cfg *env.Config) {
	cfg.Db.Close()
	DeleteDir(cfg.Root)
}
