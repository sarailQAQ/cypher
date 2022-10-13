package cypher

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
