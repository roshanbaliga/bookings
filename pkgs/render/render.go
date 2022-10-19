package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/roshanbaliga/bookings/pkgs/config"
	"github.com/roshanbaliga/bookings/pkgs/models"
)

var appConfig *config.AppConfig

func SetupRender(a *config.AppConfig) {
	appConfig = a
}

func RenderTemplate(w http.ResponseWriter, page string, td *models.TemplateData) {
	var templateCache map[string]*template.Template
	if appConfig.UseCache {
		templateCache = appConfig.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}
	parsedTemplate, found := templateCache[page]
	if !found {
		fmt.Println("Template for page: " + page + " not found")
	}
	err := parsedTemplate.Execute(w, td)
	if err != nil {
		fmt.Println("Error in launching executing template:" + err.Error())
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	fmt.Println("Creating template cache")
	templateCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return templateCache, err
	}
	for _, page := range pages {
		templates := []string{page}
		layouts, _ := filepath.Glob("./templates/*.layout.html")
		templates = append(templates, layouts...)
		renderedTemplate, err := template.ParseFiles(templates...)
		if err != nil {
			return templateCache, err
		}
		templateCache[filepath.Base(page)] = renderedTemplate
	}
	return templateCache, nil
}
