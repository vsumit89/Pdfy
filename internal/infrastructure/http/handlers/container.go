package handlers

import (
	"pdfy/internal/domain/service"
	"pdfy/internal/infrastructure/http/handlers/templates"
)

// HandlerDeps contains the dependencies required for all the handlers
type HandlerDeps struct {
	TemplateSvc service.ITemplateService
}

// HandlerContainer holds references to handler dependencies
type HandlerContainer struct {
	TemplateHandler *templates.TemplateHandler
}

func NewHandlerContainer(deps *HandlerDeps) *HandlerContainer {
	// instantiating the HandlerContainer object which will contain all the handler references
	container := &HandlerContainer{}

	container.TemplateHandler = templates.NewTemplateHandler(deps.TemplateSvc)

	return container
}
