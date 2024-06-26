package memo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestSequential(t *testing.T) {
	memo := New(httpGetBody)
	defer memo.Close()

	testSequential(t, memo)
}

func TestConcurrent(t *testing.T) {
	memo := New(httpGetBody)
	defer memo.Close()

	testConcurrent(t, memo)
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func httpGetURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"https://go.dev",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"https://go.dev",
		} {
			ch <- url
		}
		close(ch)
	}()

	return ch
}

type M interface {
	Get(url string) (interface{}, error)
}

func testSequential(t *testing.T, m M) {
	for url := range httpGetURLs() {

		val, err := m.Get(url)

		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s %d bytes\n", url, len(val.([]byte)))
	}
}

func testConcurrent(t *testing.T, m M) {
	var n sync.WaitGroup

	for url := range httpGetURLs() {
		n.Add(1)

		go func(url string) {
			defer n.Done()

			start := time.Now()
			value, err := m.Get(url)

			if err != nil {
				log.Print(err)
				return
			}

			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}

	n.Wait()
}
