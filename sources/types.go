package sources

import "os"

type (
	SourceType string

	SourceInterface interface {
		Load()
		Reload() error
		Read() []map[string]interface{}
		GetOrDefault(key string, value interface{}) interface{}
		GetOrSet(key string, value interface{}) *SourceInterface
		Remove(key string) bool
		Set(key string, value interface{}) *SourceInterface
	}

	MemorySource struct {
		_alias  string
		_values []map[string]interface{}
		_type   SourceType
	}

	EnvSource struct {
		_alias      string
		_env        []map[string]interface{}
		_type       SourceType
		_readFromOS bool
	}

	Source interface {
		GetAlias() string
		GetType() SourceType
		GetSourceFile() os.File
		GetSourceEnv() []map[string]interface{}
	}
)
