package technology

import "tech_platform/server/internal/model/technology"

type Store interface {
	Add(t technology.Technology)(technology.AddTechnology,error)
	GetOneById(id int64)(technology.Technology,error)
}

func Add(store Store,technology technology.Technology)(technology.AddTechnology,error)  {
	return store.Add(technology)
}

func GetOneById(store Store,id int64)(technology.Technology,error){
	return store.GetOneById(id)
}


