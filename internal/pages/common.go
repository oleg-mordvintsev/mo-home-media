package pages

import (
	"fmt"
	"go-server/internal/config"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
)

type MediaPath struct {
	Name string
	Path string
}

type CommonData struct {
	Title    string
	Protocol string
	Host     string
	once     sync.Once
	Dirs     []MediaPath
	Paths    []MediaPath
}

func CommonPage(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	cfg.SetOnceProtocolHost(strings.ToLower(strings.Split(r.Proto, "/")[0]), r.Host)

	setHeaders(w)

	query := r.URL.Query()

	switch {
	case query.Has("search"):
		searchPage(w, cfg, query.Get("search"))
	default:
		sourcePage(w, r, cfg, query.Get("source"))
	}

	return
}

func sourcePage(w http.ResponseWriter, r *http.Request, cfg *config.Config, source string) {
	sourcePath, done := formSourcePath(w, r, cfg, source)
	if done {
		return
	}

	tmpl, done := formTemplatePath(w, cfg)
	if done {
		return
	}

	// Подготовка к получению данных о медиафайлах
	var paths, dirs []MediaPath

	done = formDotDirs(w, r, cfg, sourcePath, &dirs)
	if done {
		return
	}

	entries, done := getEntries(w, r, sourcePath)
	if done {
		return
	}

	done = fillMediaPaths(w, r, cfg, entries, sourcePath, &dirs, &paths)
	if done {
		return
	}

	var title string
	baseDir := filepath.Base(source)

	if baseDir != "" {
		title = baseDir
	} else {
		title = "Корень"
	}

	// Данные для шаблона
	data := CommonData{
		Title:    title,
		Protocol: cfg.Protocol(),
		Host:     cfg.Host(),
		Dirs:     dirs,
		Paths:    paths,
	}

	// Выполняем шаблон
	if !isExecuteResponse(w, tmpl, &data) {
		return
	}
}

func searchPage(w http.ResponseWriter, cfg *config.Config, search string) {
	tmpl, done := formTemplatePath(w, cfg)
	if done {
		return
	}

	// Подготовка к получению данных о медиафайлах
	var paths []MediaPath
	var dirs []MediaPath

	err := fillFound(cfg, search, &dirs, &paths)
	if err != nil {
		errorMsg := fmt.Sprintf("Проблема при получении данных при поиске: %v", err)
		log.Printf(errorMsg)
		http.Error(w, errorMsg, http.StatusInternalServerError)
	}

	// Данные для шаблона
	data := CommonData{
		Title:    "Поиск",
		Protocol: cfg.Protocol(),
		Host:     cfg.Host(),
		Dirs:     dirs,
		Paths:    paths,
	}

	// Выполняем шаблон
	if !isExecuteResponse(w, tmpl, &data) {
		return
	}
}

func videoExtensions() *map[string]bool {
	return &map[string]bool{
		"avi":  true,
		"mkv":  true,
		"mp4":  true,
		"mpeg": true,
		"mpg":  true,
		"ts":   true,
		"vob":  true,
		"wmv":  true,
		"wvc":  true,
		"webm": true,
		"asf":  true,
		"3gp":  true,
		"3g2":  true,
		"mts":  true,
		"mov":  true,
		"flv":  true,
	}
}
