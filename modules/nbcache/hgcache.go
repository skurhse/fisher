package hgcache

// Package hgcache provides a concurrency-safe
// memoization of a function of type HTTPGetFunction.

type HTTPGetCache struct {
	requests chan httpGetRequest
}

type HTTPGetFunction func(url string) (interface{}, error)

type httpGetRequest struct {
	url      string
	response chan<- httpGetResult
}

type httpGetResult struct {
	val interface{}
	err error
}

type httpGetEntry struct {
	result httpGetResult
	ready  chan struct{}
}

func New(getFunction HTTPGetFunction) *HTTPGetCache {
	cache := &HTTPGetCache{
		requests: make(chan httpGetRequest),
	}

	go cache.serve(getFunction)

	return cache
}

func (getCache *HTTPGetCache) Get(url string) (interface{}, error) {
	getResponse := make(chan httpGetResult)

	getCache.requests <- httpGetRequest{url, getResponse}

	response := <-getResponse

	return response.val, response.err
}

func (getCache *HTTPGetCache) Close() {
	close(getCache.requests)
}

func (getCache *HTTPGetCache) serve(getFunction HTTPGetFunction) {
	cache := make(map[string]*httpGetEntry)

	for request := range getCache.requests {
		entry := cache[request.url]

		if entry == nil {
			entry = &httpGetEntry{ready: make(chan struct{})}
			cache[request.url] = entry

			go entry.call(getFunction, request.url)
		}

		go entry.deliver(request.response)
	}
}

func (getEntry *httpGetEntry) call(getFunction HTTPGetFunction, url string) {
	getEntry.result.val, getEntry.result.err = getFunction(url)
	close(getEntry.ready)
}

func (getEntry *httpGetEntry) deliver(result chan<- httpGetResult) {
	<-getEntry.ready
	result <- getEntry.result
}
