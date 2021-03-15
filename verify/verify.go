package verify

var keyMapping = make(map[string]IVerify)

// New 获取验证实例
func New(cb func(string) (string, bool)) IVerify {
	if cb == nil {
		return nil
	}
	for key, handler := range keyMapping {
		if value, ok := cb(key); ok {
			return handler.Args(value)
		}
	}

	return nil
}

func init() {
	keyMapping["length"] = new(lengthHandler)
}
