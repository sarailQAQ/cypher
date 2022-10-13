package cypher

import (
	"fmt"
	"testing"
)

func TestNewAesCypher(t *testing.T) {
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
}

func BenchmarkNewAesCypher(b *testing.B) {
	cypher, err := NewAesCypher()
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		dst, err := cypher.Encrypt([]byte("MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAm/rz+mIH8QKXnvy3cO6pG33fGBGj" +
			"MQDPhrKRmB1xFwrlYIIZA8hxrQTkT1g3NfQFV5p0HygO1fp323u6PfkPSMJAAmRegXnaRXctG6Jkum/Opdk8tuILMxiLgwgkIIJxKQ" +
			"6Sc8TyVznZ0wZx2tGwdBZsADBdKPUiHiUVykgeD48PvgWS6+mimLOmJNyJEsb+EHHx6YBtdBBgg6oSjZXPKkUoikdKoCDOP3o1JVRq" +
			"qpw17VWxOxZtzMwXPi2ECYkBwPDhTGLyOWEYpDYc0jXceEl6pJM27NrVuOHwYJvY7AVuvAfHc8kNTDHEm47heO+piY9nIYmnuiUL2E" +
			"YY2y65TwIDAQAB\n\n2月5日 05:25\nhttps://be-prod.redrock.cqupt.edu.cn/magicloop-sso/callback?state=ef505" +
			"1ec3063d78a2875aafdc7a248f1&session_state=fbd8185f-b044-4e1f-a416-3df7991a0ee4&code=79cc3db0-0e49-4301-" +
			"b17c-5f2acf42e153.fbd8185f-b044-4e1f-a416-3df7991a0ee4.872ba2d7-50dc-45a6-b08b-03d618d48683"))
		if err != nil {
			panic(err)
		}

		_, err = cypher.Decrypt(dst)
		if err != nil {
			panic(err)
		}
	}
}
