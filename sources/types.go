package sources

import "os"

type (
	SourceType string

	SourceInterface interface {
		Load()
		Keys() []string
		Reload() error
		Read() []map[string]interface{}
		GetOrDefault(key string, value interface{}) interface{}
		GetOrSet(key string, value interface{}) *SourceInterface
		Remove(key string) bool
		Set(key string, value interface{}) *SourceInterface
		GetAlias() string
		GetType() SourceType
		GetSourceFile() os.File
		GetSourceEnv() []map[string]interface{}
	}

	EnvSource struct {
		_alias      string
		_env        []map[string]interface{}
		_type       SourceType
		_readFromOS bool
	}
)
