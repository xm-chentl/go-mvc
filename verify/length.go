package verify

import (
	"strings"
)

type lengthHandler struct {
	argsConfig string
}

func (h *lengthHandler) Args(argsConfig string) IVerify {
	h.argsConfig = argsConfig
	return h
}

func (h lengthHandler) Execute(value interface{}) bool {
	// desc: 解析argsConfig
	args := strings.Split(h.argsConfig, ",")
	str = value.(string)
	if len(args) == 2 {
		// 区间 0 min, 1 max
		return len(str) > args[0].(int) && len(str) <= args[1].(int)
	}
	if len(args) == 1 {
		return len(str) >= args[0].(int)
	}

	return true
}
