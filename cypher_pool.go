package cypher

import "sync"

// NewCypherPool 一个Cypher池，用于处理需要并发都场景
func NewCypherPool(newFunc func(configs ...Config) (Cypher, error), configs ...Config) (pool *Pool, err error) {
	// 先判断newFunc是否正确，防止开发者错误使用
	c, err := newFunc(configs...)
	if err != nil {
		return nil, err
	}

	pool = &Pool{
		p: &sync.Pool{
			New: func() interface{} {
				_cypher, _ := newFunc(configs...)
				return _cypher
			},
		},
	}

	// 都生成了，不如直接放进去
	pool.Put(c)
	return
}

type Pool struct {
	p *sync.Pool
}

func (pool *Pool) Get() Cypher {
	return pool.p.Get().(Cypher)
}

func (pool *Pool) Put(c Cypher) {
	pool.p.Put(c)
}
