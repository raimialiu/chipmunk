package chipmunk

import "github.com/raimialiu/chipmunk.git/sources"

func Load(target interface{}) *Chipmunk {
	mem_source := sources.NewMemorySource()
	env_sources := make([]*sources.SourceInterface, 0)
	env_sources = append(env_sources, mem_source)
	return &Chipmunk{
		_sources: env_sources,
		_target:  target,
	}
}
