package env

import (
	"github.com/etcenter/c4/db"
)

type Config struct {
	// Root path for c4 configuration and storage
	Root *string
	Db   *db.DB
}

var global_config *Config

// NewConfig returns a new Config pointer that can be chained with builder methods to
// set multiple configuration values inline without using pointers.
//
//     svc := storage.NewAttributeStore(c4.NewConfig().WithRoot("/mnt/c4"))
//
func NewConfig() *Config {
	return &Config{}
}

func SetGlobalConfig(config *Config) {
	global_config = config
}

func GetGlobalConfig() *Config {
	return global_config
}

// WithRoot sets the Root path and returns a Config pointer for chaining.
func (c *Config) WithRoot(path string) *Config {
	c.Root = &path
	return c
}

// WithDB sets the c4 db handle and returns a Config pointer for chaining.
func (c *Config) WithDb(db *db.DB) *Config {
	c.Db = db
	return c
}

// Merge merges an array of configs.
func (c *Config) Merge(cfgs ...*Config) {
	for _, src := range cfgs {
		mergeConfigs(c, src)
	}
}

func mergeConfigs(dst *Config, src *Config) {
	if src == nil {
		return
	}

	if src.Root != nil {
		dst.Root = src.Root
	}
}
