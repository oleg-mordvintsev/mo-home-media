package pages

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

func isExecuteResponse(w http.ResponseWriter, tmpl *template.Template, data *CommonData) bool {
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Ошибка выполнения шаблона: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}
	return true
}

func hasParsedTemplate(w http.ResponseWriter, templatePath string) (*template.Template, bool) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		errorMsg := fmt.Sprintf("Ошибка парсинга шаблона: %v", err)
		log.Printf(errorMsg)
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return nil, false
	}
	return tmpl, true
}

func isExistTemplate(w http.ResponseWriter, templatePath string) bool {
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		errorMsg := fmt.Sprintf("Шаблон не найден по пути: %s", templatePath)
		log.Printf(errorMsg)
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return false
	}
	return true
}
