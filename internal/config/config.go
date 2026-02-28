package config

import (
	"sync"
)

// Singleton конфигурации

var (
	cfg  *Config // Глобальный конфиг
	once sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		cfg = &Config{} // Инициализируем один раз
		cfg.load()      // Заполняем значения
	})
	return cfg
}
