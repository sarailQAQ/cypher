package _aes

import (
	"errors"
	"github.com/sarailQAQ/cypher/_coding"
)

func NewAesCypher(key []byte, mode, coding string) (*Cypher, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errors.New("invalid key length, Aes`s key`s length must be 16,24 or 32")
	}

	return &Cypher{
		Key:    key,
		Mode:   mode,
		coding: coding,
	}, nil
}

// Cypher 实现Aes算法的Cypher接口
type Cypher struct {
	Key    []byte
	Mode   string
	coding string
}

func (c *Cypher) Encrypt(src []byte) (dst []byte, err error) {
	defer func() {
		if i := recover(); i != nil {
			if e, ok := i.(error); ok {
				err = e
				return
			}
			err = errors.New("meet unknown error in Aes Encrypting")
		}
	}()

	switch c.Mode {
	case ModeCBC:
		dst, err = encryptCBC(src, c.Key)
	default:
		return nil, ErrNoSuchMode
	}

	switch c.coding {
	case _coding.CodeNull:
	case _coding.CodeBase64:
		dst = _coding.EncodeBase64(dst)
	case _coding.CodeBase64URL:
		dst = _coding.EncodeBase64Url(dst)
	}

	return dst, err
}

func (c *Cypher) Decrypt(src []byte) (dst []byte, err error) {
	defer func() {
		if i := recover(); i != nil {
			if e, ok := i.(error); ok {
				err = e
				return
			}
			err = errors.New("meet unknown error in Aes Decrypting")
		}
	}()

	switch c.coding {
	case _coding.CodeNull:
	case _coding.CodeBase64:
		src, err = _coding.DecodeBase64(src)
	case _coding.CodeBase64URL:
		src, err = _coding.DecodeBase64Url(src)
	}
	if err != nil {
		return nil, err
	}

	switch c.Mode {
	case ModeCBC:
		return decryptCBC(src, c.Key)
	default:
		return nil, ErrNoSuchMode
	}
}
