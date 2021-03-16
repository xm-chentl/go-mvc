package verify

import (
	"strconv"
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
	if len(args) == 0 {
		return true
	}

	min, _ := strconv.Atoi(args[0])
	str := value.(string)
	if len(args) == 2 {
		// 区间 0 min, 1 max
		max, _ := strconv.Atoi(args[1])
		return len(str) > min && len(str) <= max
	}
	if len(args) == 1 {
		return len(str) >= min
	}

	return true
}
