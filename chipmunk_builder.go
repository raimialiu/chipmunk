package chipmunk

import (
	"github.com/raimialiu/chipmunk.git/sources"
)

func (c *Chipmunk) WithPrefix(prefix string) *Chipmunk {
	c._option.Prefix = prefix
	return c
}

func (c *Chipmunk) WithSources(sources ...sources.SourceInterface) *Chipmunk {
	c._sources = append(c._sources, sources...)
	return c
}

func (c *Chipmunk) WithValidation(enable bool) *Chipmunk {
	c._option.TagOption.Validation = enable
	return c
}
