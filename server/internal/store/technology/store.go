package technology

import "tech_platform/server/internal/model/technology"

type Store interface {
	Add(t technology.Technology)(technology.AddTechnology,error)
}

func Add(store Store,technology technology.Technology)(technology.AddTechnology,error)  {
	return store.Add(technology)
}


