package scope

type Value int

const (
	// Server 服务访问级别
	Server Value = iota + 1
	// Platform 平台访问级别
	Platform
)
