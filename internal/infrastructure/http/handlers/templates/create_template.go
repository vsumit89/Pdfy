package templates

import "net/http"

func (t *TemplateHandler) CreateTemplate(res http.ResponseWriter, req *http.Request) {
	t.templateSvc.CreateTemplate()
}
