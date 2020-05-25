package g

import "github.com/goees/yago"

const (
	// 1000 - 9999 业务公共错误, 10000 - .... 业务错误
	ErrDemo = yago.Err("10001=error demo")
)
