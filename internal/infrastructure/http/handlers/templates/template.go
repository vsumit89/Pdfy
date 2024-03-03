package templates

import "pdfy/internal/domain/service"

type TemplateHandler struct {
	templateSvc service.ITemplateService
}

func NewTemplateHandler(svc service.ITemplateService) *TemplateHandler {
	return &TemplateHandler{
		templateSvc: svc,
	}
}


