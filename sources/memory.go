package sources

func NewMemorySource() *MemorySource {
	return &MemorySource{
		_alias:  "memory",
		_values: make([]map[string]interface{}, 0),
		_type:   MEM,
	}
}
