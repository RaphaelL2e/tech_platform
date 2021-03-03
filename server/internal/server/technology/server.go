package technology

import (
	"context"
	"tech_platform/server/internal/model/technology"
	"tech_platform/server/internal/pkg/response"
	"tech_platform/server/internal/store"
	technologystore "tech_platform/server/internal/store/technology"
)

type Handler struct {
}

func (h Handler) AddTechnology(c context.Context,req technology.Technology) response.ServerResponse {

	tech_store := store.FromContext(c)
	at,err :=technologystore.Add(tech_store,req)
	if err!=nil{
		return response.CreateByErrorMessage(err)
	}
	return response.CreateBySuccessData(at)
}

func NewHandler() *Handler {
	return &Handler{}
}
