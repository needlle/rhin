package pipeline

import "rhin/lib"

type Pipeline struct {
	items []any
}

func NewPipeline(items []any) *Pipeline {
	return &Pipeline{items}
}

func (pip Pipeline) Filter(condition func(item any) bool) *Pipeline {
	return NewPipeline(lib.Filter(pip.items, condition))
}

func (pip Pipeline) Map(transform func(item any) any) *Pipeline {
	return NewPipeline(lib.Map(pip.items, transform))
}
