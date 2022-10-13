package _coding

import (
	"encoding/base64"
)

// DecodeBase64 base解码
func DecodeBase64(src []byte) (dst []byte, err error) {
	dst = make([]byte, base64.StdEncoding.DecodedLen(len(src)))
	n, err := base64.StdEncoding.Decode(dst, src)
	return dst[:n], err
}

// EncodeBase64 base64编码
func EncodeBase64(src []byte) (dst []byte) {
	dst = make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return
}

func DecodeBase64Url(src []byte) (dst []byte, err error) {
	dst = make([]byte, base64.URLEncoding.DecodedLen(len(src)))
	n, err := base64.URLEncoding.Decode(dst, src)
	return dst[:n], err
}

func EncodeBase64Url(src []byte) (dst []byte) {
	dst = make([]byte, base64.URLEncoding.EncodedLen(len(src)))
	base64.URLEncoding.Encode(dst, src)
	return
}
