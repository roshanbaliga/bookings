package handlers

import (
	"net/http"

	"github.com/roshanbaliga/bookings/pkgs/config"
	"github.com/roshanbaliga/bookings/pkgs/models"
	"github.com/roshanbaliga/bookings/pkgs/render"
)

var Repo *Repository

type Repository struct {
	AppConfig *config.AppConfig
}

func CreateRepositry(a *config.AppConfig) *Repository {
	return &Repository{
		AppConfig: a,
	}
}

func SetupHandlers(r *Repository) {
	Repo = r
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteAddr := r.RemoteAddr
	repo.AppConfig.Session.Put(r.Context(), "remote_ip", remoteAddr)
	strMap := map[string]string{"greeting": "Hello world"}
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{StringsMap: strMap})
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	strMap := map[string]string{"remote_ip": repo.AppConfig.Session.GetString(r.Context(), "remote_ip")}
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{StringsMap: strMap})
}
