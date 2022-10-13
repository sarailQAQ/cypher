# cypher

对标准库的懒人封装

## Cypher

实现了加密算法的接口，定义为：

```go
// Encryptor 加密器
type Encryptor interface {
	Encrypt(src []byte) (dst []byte, err error)
}

// Decryptor 解密器
type Decryptor interface {
	Decrypt(dst []byte) (src []byte, err error)
}

type Cypher interface {
	Encryptor
	Decryptor
}
```

由于大部门的加密/解密场景都会将string转为[]byte，因此数据的入参和出参都是[]byte。

加密算法都有其默认配置，如默认的密钥、加密模式等，若需自定义加密配置，请通过在New函数中传入Config对象指定。

由于部分场景需要将加密后的结果放入url或是http请求的响应中，因此加密结果默认进行URL安全的Base64编码。若无需进行编码，也请通过传入Config对象指定。

使用：

```go
	cypher, err := NewAesCypher()
	if err != nil {
		panic(err)
	}

	dst, err := cypher.Encrypt([]byte("abcdefghijklmn1234567890!@#$%^&*()?><"))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dst))

	src, err := cypher.Decrypt(dst)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(src))
```

## Pool	

基于sync.Pool封装的Cypher对象池，同样也是并非安全，且不必再额外手写New()方法和类型断言。

使用：

```go
	// 第一个参数是Cypher的New函数，第二个参数传递相关配置
	pool, err := NewCypherPool(NewAesCypher, NewConfig("key", "1234567812345678"))
	if err != nil {
		panic(err)
	}

	ch := make(chan []byte)
	var wg sync.WaitGroup

	for i := 0; i < 100000; i++ {
		go func(i int) {
			wg.Add(1)
			_cypher := pool.Get()
			dst, err := _cypher.Encrypt([]byte(strconv.FormatInt(int64(i), 10)))
			if err != nil {
				panic(err)
			}
			ch <- dst
		}(i)
	}

	for i := 0; i < 100000; i++ {
		go func() {
			defer wg.Done()

			src := <- ch
			_cypher := pool.Get()
			_, err := _cypher.Decrypt(src)
			if err != nil {
				panic(err)
			}
		}()
	}
	wg.Wait()
```

