package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/wire"
)

var ConfigProvider = wire.NewSet(
	ProvideApplicationConfig,
	ProvideFlowableConfig,
)

type FlowableConfig struct {
	URL string `json:"url"`
}

type ApplicationConfig struct {
	*FlowableConfig `json:"flowable"`
}

func ProvideApplicationConfig() (*ApplicationConfig, error) {
	fp, ok := os.LookupEnv("CONFIG")
	if !ok {
		fp = "./application.config.json"
	}
	data, err := os.ReadFile(fp)
	if err != nil {
		return nil, fmt.Errorf("read config of application failed: %w", err)
	}

	var config ApplicationConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("unmarshal config of application failed: %w", err)
	}
	return &config, nil
}

func ProvideFlowableConfig(applicationConfig *ApplicationConfig) (*FlowableConfig, error) {
	if applicationConfig.FlowableConfig == nil {
		return nil, fmt.Errorf("config of flowable is nil")
	}

	return applicationConfig.FlowableConfig, nil
}
