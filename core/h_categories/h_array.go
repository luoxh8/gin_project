package h_categories

type Array []interface{}

func (a *Array) Append(obj interface{}) {
	*a = append(*a, obj)
}

func (a *Array) Pop(obj interface{}) {
	temp1 := *a
	temp2 := make([]interface{}, len(*a)-1, 2*(len(*a))-1)
	_index := 0
	for i := 0; i < len(*a); i++ {
		if temp1[i] != obj {
			temp2[_index] = temp1[i]
			_index += 1
		}
	}
	*a = temp2
}

func (a *Array) In(obj interface{}) bool {
	for _, value := range *a {
		if value == obj {
			return true
		}
	}
	return false
}

func (a *Array) Get(obj interface{}) interface{} {
	for _, value := range *a {
		if obj == value {
			return value
		}
	}
	return nil
}
