package cypher

import (
	"strconv"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	pool, err := NewCypherPool(NewAesCypher, NewConfig("key", "1234567812345678"))
	if err != nil {
		panic(err)
	}

	ch := make(chan []byte, 10)
	var wg sync.WaitGroup

	for i := 0; i < 100000; i++ {
		go func(i int) {
			wg.Add(1)
			_cypher := pool.Get()
			dst, err := _cypher.Encrypt([]byte(strconv.FormatInt(int64(i), 10)))
			if err != nil {
				panic(err)
			}
			pool.Put(_cypher)
			ch <- dst
		}(i)
	}

	for i := 0; i < 100000; i++ {
		go func() {
			defer wg.Done()

			src := <- ch
			_cypher := pool.Get()
			pool.Put(_cypher)
			_, err := _cypher.Decrypt(src)
			if err != nil {
				panic(err)
			}
		}()
	}
	wg.Wait()
}
