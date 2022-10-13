package cypher

// Config 加密算法需要的一些配置
// 比如：密钥（公钥私钥）、加密模式、加密（解密）次数等
type Config struct {
	Key   string
	Value interface{}
}

func NewConfig(key string, val interface{}) Config {
	return Config{
		Key:   key,
		Value: val,
	}
}
