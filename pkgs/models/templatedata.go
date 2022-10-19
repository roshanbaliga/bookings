package models

type TemplateData struct {
	StringsMap  map[string]string
	IntegersMap map[string]int
	FloatMap    map[string]float32
	Data        map[string]interface{}
	CSRFToken   string
	Flash       string
	Warning     string
	Error       string
}
