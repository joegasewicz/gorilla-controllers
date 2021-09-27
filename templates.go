package gorillacontrollers

type GTemplate struct {
	BaseTemplates []string
}

func (t *GTemplate) GTemplates(routeTemplateName ...string) []string {
	for _, template := range routeTemplateName {
		t.BaseTemplates = append(t.BaseTemplates, template)
	}
	return t.BaseTemplates
}
