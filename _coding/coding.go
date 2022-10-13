package _coding

import (
	"errors"
	"strings"
)

const (
	CodeNull      = "null"
	CodeBase64    = "base64"
	CodeBase64URL = "base64url"
)

var ErrNoSuchCodeMethod = errors.New("no such code method")

func Encode(src []byte, codingMethod string) (dst []byte, err error) {
	codingMethod = strings.ToLower(codingMethod)
	switch codingMethod {
	case CodeNull:
		return src, nil
	case CodeBase64:
		return EncodeBase64(src), nil
	case CodeBase64URL:
		return EncodeBase64Url(src), nil
	default:
		return nil, ErrNoSuchCodeMethod
	}
}

func Decode(src []byte, codingMethod string) (dst []byte, err error) {
	codingMethod = strings.ToLower(codingMethod)
	switch codingMethod {
	case CodeNull:
		return src, nil
	case CodeBase64:
		return DecodeBase64(src)
	case CodeBase64URL:
		return DecodeBase64Url(src)
	default:
		return nil, ErrNoSuchCodeMethod
	}
}
