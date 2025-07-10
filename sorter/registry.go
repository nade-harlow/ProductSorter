package sorter

type Registry struct {
	sorters map[SorterType]Sorter
}

func NewRegistry() *Registry {
	return &Registry{
		sorters: map[SorterType]Sorter{},
	}
}

func (r *Registry) Register(key SorterType, s Sorter) {
	r.sorters[key] = s
}

func (r *Registry) Get(key SorterType) Sorter {
	return r.sorters[key]
}
