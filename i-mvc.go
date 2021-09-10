package mvc

type IMvc interface {
	// Method todo: 扩展其它协议 PUT, DELETE
	Run(port int)
}

// 5500 + 农信8500 +
