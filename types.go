package chipmunk

import (
	"github.com/raimialiu/chipmunk.git/sources"
)

type (
	Chipmunk struct {
		_option  Options
		_sources []sources.Source
		_target  interface{}
	}

	Options struct {
		Prefix            string
		CaseSensitive     bool
		Separator         string
		NestedSeparator   string
		TagOption         StructTagOptions
		StrictMode        bool
		VariableExpansion bool
		RemoteLoading     bool
	}

	StructTagOptions struct {
		Alias      string
		Validation bool
		Required   bool
		Default    bool
		Expand     bool
		File       bool
	}
)
