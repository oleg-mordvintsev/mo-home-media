package pages

import (
	"go-server/internal/config"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	w.Header().Set("Expires", "Thu, 19 Feb 1998 13:24:18 GMT")
	w.Header().Set("Last-Modified", time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT"))
	w.Header().Set("Cache-Control", "no-cache, must-revalidate")
	w.Header().Add("Cache-Control", "post-check=0,pre-check=0")
	w.Header().Add("Cache-Control", "max-age=0")
	w.Header().Set("Pragma", "no-cache")
}

func fillMediaPaths(w http.ResponseWriter, r *http.Request, cfg *config.Config, entries []os.DirEntry, sourcePath string, dirs *[]MediaPath, paths *[]MediaPath) bool {
	for _, entry := range entries {
		path := filepath.Join(sourcePath, entry.Name())

		httpDirPath, err := filepath.Rel(string(cfg.Data()), path)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return true
		}

		httpFilePath := strings.TrimLeft(path, "/")

		if entry.IsDir() {
			if path == sourcePath {
				http.Redirect(w, r, "/", http.StatusMovedPermanently)
				return true
			}

			dirname := entry.Name()
			*dirs = append(*dirs, MediaPath{Name: dirname, Path: httpDirPath})

			continue
		}

		*paths = append(*paths, MediaPath{Name: entry.Name(), Path: httpFilePath})
	}

	return false
}

func fillFound(cfg *config.Config, search string, dirs *[]MediaPath, paths *[]MediaPath) error {
	err := filepath.WalkDir(string(cfg.Data()), func(path string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			if os.IsPermission(err) {
				log.Printf("Ошибка доступа при поиске: %v", err)
				return nil
			}

			return err
		}

		if dirEntry.IsDir() {
			if path == string(cfg.Data()) {
				return nil
			}

			httpPath, err := filepath.Rel(string(cfg.Data()), path)
			if err != nil {
				log.Printf("Ошибка доступа к перебираемой директории в поиске: %v", err)
				return err
			}

			dirname := dirEntry.Name()

			if strings.Contains(strings.ToLower(dirname), strings.ToLower(search)) {
				*dirs = append(*dirs, MediaPath{Name: dirname, Path: httpPath})
			}

			return nil
		}

		filename := dirEntry.Name()

		ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(filename)), ".")

		// Если у файла нужны подстрока найдена и файл с нужным расширением
		if strings.Contains(strings.ToLower(filename), strings.ToLower(search)) && (*videoExtensions())[ext] {
			httpPath, err := filepath.Rel(string(cfg.Project()), path)
			if err != nil {
				return err
			}

			*paths = append(*paths, MediaPath{Name: filename, Path: httpPath})
		}

		return nil
	})

	return err
}

func getEntries(w http.ResponseWriter, r *http.Request, sourcePath string) ([]os.DirEntry, bool) {
	entries, err := os.ReadDir(sourcePath)

	if err != nil {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return nil, true
	}

	return entries, false
}

func formDotDirs(w http.ResponseWriter, r *http.Request, cfg *config.Config, sourcePath string, dirs *[]MediaPath) bool {
	twoDotDir, err := filepath.Rel(string(cfg.Data()), filepath.Dir(sourcePath))
	if twoDotDir == "." {
		twoDotDir = ""
	}

	if err != nil {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return true
	}

	*dirs = append(*dirs, MediaPath{Name: ".", Path: ""})
	*dirs = append(*dirs, MediaPath{Name: "..", Path: twoDotDir})

	return false
}

func formTemplatePath(w http.ResponseWriter, cfg *config.Config) (*template.Template, bool) {
	templatePath := filepath.Join(string(cfg.Template()), "common.xml")

	// Есть ли нужный шаблон
	if !isExistTemplate(w, templatePath) {
		return nil, true
	}

	// Парсится ли нужный шаблон?
	tmpl, done := hasParsedTemplate(w, templatePath)
	if !done {
		return nil, true
	}
	return tmpl, false
}

func formSourcePath(w http.ResponseWriter, r *http.Request, cfg *config.Config, source string) (string, bool) {
	// По умолчанию корневая
	sourcePath := string(cfg.Data())

	if source != "" {
		sourcePath = filepath.Join(string(cfg.Data()), source)
	}

	if !strings.Contains(sourcePath, string(cfg.Data())) {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return "", true
	}

	return sourcePath, false
}
