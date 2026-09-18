package provider

import (
	"github.com/swp391-group3/ai-interview-practice/api/internal/config"
)

func ProvideConfig(configPath string) (*config.Config, error) {
	return config.Load(configPath)
}
