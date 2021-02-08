package store

import "context"

const key = "_store"

type Setter interface {
	Set(string, interface{})
}

func FromContext(c context.Context) Store {
	return c.Value(key).(Store)
}

func ToContext(c Setter,store Store){
	c.Set(key,store)
}
