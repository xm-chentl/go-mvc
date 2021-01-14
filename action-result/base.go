package actionresult

type base struct {
	Err  int
	Data interface{}
}

func (b base) Exec() interface{} {
	return map[string]interface{}{
		"err":  b.Err,
		"data": b.Data,
	}
}
