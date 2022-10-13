package _coding

import (
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	str := "asdkjahjksdhjk-02198	1[]./.,.,m.\" `''阿迪健康`"
	fmt.Println([]byte(str))
	dst := EncodeBase64Url([]byte(str))
	fmt.Println(string(dst))

	dst, err := DecodeBase64Url(dst)
	if err != nil {
		panic(err)
	}
	fmt.Println(dst)
}
