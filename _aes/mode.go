package _aes

import "errors"

const (
	ModeCBC = "cbc"
	ModeECB = "ecb"
)

var ErrNoSuchMode = errors.New("no such mode")
