package cache

// A MemCached caches the results of calling a Func.
type MemCached struct {
	f      Func
	cached map[string]result
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *MemCached {
	return &MemCached{f: f, cached: make(map[string]result)}
}

func (mem *MemCached) Get(key string) (interface{}, error) {
	res, ok := mem.cached[key]
	if !ok {
		res.value, res.err = mem.f(key)
		mem.cached[key] = res
	}
	return res.value, res.err
}
