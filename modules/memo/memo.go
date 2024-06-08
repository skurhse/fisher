package memo

// Package memo provides a `Memoizer` server
// for concurrency-safe caching of `Calculator` results.

type Memoizer struct {
	requests chan request
}

type Calculator func(key string) (interface{}, error)

type request struct {
	key    string
	result chan<- result
}

type result struct {
	val interface{}
	err error
}

type entry struct {
	result result
	ready  chan struct{}
}

func New(calc Calculator) (memo *Memoizer) {
	memo = &Memoizer{
		requests: make(chan request),
	}

	go memo.serve(calc)
	return
}

func (memo *Memoizer) Get(key string) (interface{}, error) {
	ch := make(chan result)

	memo.requests <- request{key, ch}
	res := <-ch

	return res.val, res.err
}

func (memo *Memoizer) Close() {
	close(memo.requests)
}

func (memo *Memoizer) serve(calc Calculator) {
	cache := make(map[string]*entry)

	for req := range memo.requests {
		e := cache[req.key]

		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e

			go e.call(calc, req.key)
		}

		go e.deliver(req.result)
	}
}

func (e *entry) call(calc Calculator, key string) {
	e.result.val, e.result.err = calc(key)
	close(e.ready)
}

func (e *entry) deliver(res chan<- result) {
	<-e.ready
	res <- e.result
}
