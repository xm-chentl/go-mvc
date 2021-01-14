package enum

// RouteMode 路由
type RouteMode int

const (
	// BinaryMode 二段模式
	BinaryMode RouteMode = iota
	// ThreeMode 三段模式
	ThreeMode
)
