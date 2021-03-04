package technology

import "tech_platform/server/internal/model/technology"

type Store interface {
	Add(t technology.Technology)(technology.AddTechnology,error)
	GetOneById(id int64)(technology.Technology,error)
	Update(t technology.Technology)(technology.Technology,error)
	Delete(t technology.DeleteTechnology)(bool,error)
	List(t technology.ListModel)([]technology.ListTechnology,error)
}

func Add(store Store,technology technology.Technology)(technology.AddTechnology,error)  {
	return store.Add(technology)
}

func GetOneById(store Store,id int64)(technology.Technology,error){
	return store.GetOneById(id)
}

func Update(store Store,technology technology.Technology)(technology.Technology,error)  {
	return store.Update(technology)
}

func Delete(store Store,technology technology.DeleteTechnology)(bool,error)  {
	return store.Delete(technology)
}

func List(store Store,technology technology.ListModel)([]technology.ListTechnology,error)  {
	return store.List(technology)
}

