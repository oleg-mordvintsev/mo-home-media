package main

import (
	"go-server/internal/config"
	"go-server/internal/pages"
	"log"
	"net/http"
)

type App struct {
	Cfg *config.Config
}

func main() {
	app := &App{Cfg: config.GetConfig()}

	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/", app.commonPage)

	log.Fatal(http.ListenAndServe(":"+app.Cfg.Port(), nil))
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/static/favicon.ico", http.StatusMovedPermanently)
}

func (app *App) commonPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	pages.CommonPage(w, r, app.Cfg)
}
