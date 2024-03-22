package dbs

type option struct {
	query    []Query
	order    any
	offset   int
	limit    int
	preloads []string
}
type optionFn func(*option)

func (f optionFn) apply(opt *option) {
	f(opt)
}

type FindOption interface {
	apply(*option)
}

func getOption(opts ...FindOption) option {
	opt := option{
		query:  []Query{},
		offset: 0,
		limit:  1000,
		order:  "id",
	}

	for _, o := range opts {
		o.apply(&opt)
	}

	return opt
}
