package cypher

import (
	"github.com/sarailQAQ/cypher/_aes"
	"github.com/sarailQAQ/cypher/_coding"
	"strings"
)

// NewAesCypher Aes加密算法
// 目前支持等配置项有：
// key：密钥，必须为16、24或32字节，支持string和[]byte类型
// mode：加密模式，目前支持cbc模式，支持string类型，建议使用_aes包中的常量
// coding: 加密结果是否进行编码目前支持"null"、"base64"、"base64url"。默认使用base64url编码
func NewAesCypher(configs ...Config) (Cypher, error) {
	mode := _aes.ModeCBC
	key := []byte("abccdefabccdef12")
	coding := _coding.CodeBase64URL
	for _, c := range configs {
		switch strings.ToLower(c.Key) {
		case "mode":
			m, ok := c.Value.(string)
			if !ok {
				break
			}
			mode = strings.ToLower(m)

		case "key":
			switch c.Value.(type) {
			case string:
				key = []byte(c.Value.(string))
			case []byte:
				key = c.Value.([]byte)
			}
		case "coding":
			if cc, ok := c.Value.(string); ok {
				coding = cc
			}
		}
	}

	return _aes.NewAesCypher(key, mode, coding)
}
