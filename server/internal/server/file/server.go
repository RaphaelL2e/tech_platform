package file

import (
	"tech_platform/server/pkg/ossutil"
)

type Handler struct {
	OSSHelper ossutil.OSSHelper
}

func (h Handler) upload() {

}

func NewHandler(helper ossutil.OSSHelper) *Handler {
	return &Handler{OSSHelper: helper}
}