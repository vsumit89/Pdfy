package service

import "fmt"

type ITemplateService interface {
	CreateTemplate() error
}

func NewTemplateSvc() ITemplateService {
	return &templateSvcImpl{}
}

// templateSvcImpl implements all the methods defined by ITemplateService
type templateSvcImpl struct{}

// CreateTemplate creates the html template which will be rendered as a definition
func (t *templateSvcImpl) CreateTemplate() error {
	fmt.Println("Creating Template Successfully")
	return nil
}
