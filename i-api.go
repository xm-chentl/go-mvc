package mvc

// IApi 接口规范
type IApi interface {
	// Code() string
	// Scope() scope.Value
	Call() (interface{}, error)
}
