package chipmunk

import "github.com/raimialiu/chipmunk.git/sources"

func Load() *Chipmunk {
	mem_source := sources.NewMemorySource()
	env_sources := make([]*sources.Source, 0)
	return &Chipmunk{
		_sources: env_sources,
	}
}
