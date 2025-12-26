package sources

import (
	"os"
)

type MemorySource struct {
	_alias  string
	_values []map[string]interface{}
	_type   SourceType
}

func (m MemorySource) Load() {
	//TODO implement me
	panic("implement me")
}

func (m MemorySource) Keys() []string {
	//TODO implement me
	panic("implement me")
}

func (m MemorySource) Reload() error {
	//TODO implement me
	panic("implement me")
}

func (m MemorySource) Read() []map[string]interface{} {
	//TODO implement me
	panic("implement me")
}

func (m MemorySource) GetOrDefault(key string, value interface{}) interface{} {
	//TODO implement me
	panic("implement me")
}

func (m MemorySource) GetOrSet(key string, value interface{}) *SourceInterface {
	//TODO implement me
	panic("implement me")
}

func (m MemorySource) Remove(key string) bool {
	//TODO implement me
	panic("implement me")
}

func (m MemorySource) Set(key string, value interface{}) *SourceInterface {
	//TODO implement me
	panic("implement me")
}

func (m MemorySource) GetAlias() string {
	//TODO implement me
	panic("implement me")
}

func (m MemorySource) GetType() SourceType {
	//TODO implement me
	panic("implement me")
}

func (m MemorySource) GetSourceFile() os.File {
	//TODO implement me
	panic("implement me")
}

func (m MemorySource) GetSourceEnv() []map[string]interface{} {
	//TODO implement me
	panic("implement me")
}

func NewMemorySource() *MemorySource {
	return &MemorySource{
		_alias:  "memory",
		_values: make([]map[string]interface{}, 0),
		_type:   MEM,
	}
}
