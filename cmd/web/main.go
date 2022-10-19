package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/roshanbaliga/bookings/pkgs/config"
	"github.com/roshanbaliga/bookings/pkgs/handlers"
	"github.com/roshanbaliga/bookings/pkgs/render"
)

const portNumber = ":8080"

var session *scs.SessionManager

func main() {
	var app config.AppConfig

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	app.UseCache = true
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc

	render.SetupRender(&app)

	repo := handlers.CreateRepositry(&app)
	handlers.SetupHandlers(repo)

	fmt.Println("Listening on port " + portNumber)
	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		panic("Error in launching http listener:" + err.Error())
	}
}
