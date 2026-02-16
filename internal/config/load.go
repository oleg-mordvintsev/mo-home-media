package config

import (
	"log"
	"os"
	"path/filepath"
	"sync"
)

const (
	defaultPort         = "1777"
	defaultExtPort      = "777"
	defaultTemplatePath = "templates"
	defaultDataPath     = "data"
)

type Path string

// Paths
//
//	/app 			- корень проекта		- cfg.paths.project		- cfg.Project()
//	/app/template	- директория шаблонов	- cfg.paths.template	- cfg.Template()
//	/data 			- медийные данные		- cfg.paths.data		- cfg.Data()
type Paths struct {
	project, template, data Path
}
type Config struct {
	port             string
	paths            Paths
	protocol         string
	host             string
	onceProtocolHost sync.Once
}

func (cfg *Config) load() {
	cfg.port = cfg.setPort()

	// Порядок установки директорий важен
	cfg.setProject()
	cfg.setTemplate()
	cfg.setData()
}

// Сеттеры и Геттеры

func (cfg *Config) setPort() string {
	return defaultPort
}

func (cfg *Config) Port() string {
	return cfg.port
}

func (cfg *Config) setProject() {
	src, err := os.Getwd()
	if err != nil {
		log.Fatalf("Директория запуска недоступна: %v", err)
	}

	cfg.paths.project = Path(src)
}

func (cfg *Config) Project() Path {
	return cfg.paths.project
}

func (cfg *Config) setTemplate() {
	defaultDir := os.Getenv("DIR_TEMPLATES")
	if defaultDir == "" {
		defaultDir = defaultTemplatePath
	}

	dir := filepath.Join(string(cfg.Project()), defaultDir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Fatalf("Директория шаблонов недоступна: %v", err)
	}

	cfg.paths.template = Path(dir)
}

func (cfg *Config) Template() Path {
	return cfg.paths.template
}

func (cfg *Config) setData() {
	defaultDir := os.Getenv("DIR_DATA")
	if defaultDir == "" {
		defaultDir = defaultDataPath
	}

	dir := filepath.Join(filepath.Dir(string(cfg.Project())), defaultDir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Fatalf("Директория медиаданных недоступна: %v", err)
	}

	cfg.paths.data = Path(dir)
}

func (cfg *Config) Data() Path {
	return cfg.paths.data
}

func (cfg *Config) SetOnceProtocolHost(protocol string, host string) {
	cfg.onceProtocolHost.Do(func() {
		extPort := os.Getenv("EXTERNAL_PORT")
		if extPort == "" {
			extPort = defaultExtPort
		}

		cfg.protocol = protocol
		cfg.host = host + ":" + extPort
	})
}

func (cfg *Config) Protocol() string {
	return cfg.protocol
}

func (cfg *Config) Host() string {
	return cfg.host
}
