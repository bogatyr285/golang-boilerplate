package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
)

type Config struct {
	App `yaml:"app"`
}

type (
	App struct {
		// just app name
		Name    string   `yaml:"name"`
		Listen  Listen   `yaml:"listen"`
		Tracing *Tracing `yaml:"tracing"`
		Metrics *Metrics `yaml:"metrics"`
		// todo
		// Debug   *Debug   `yaml:"debug"`

		// Logging *Logging `yaml:"logging"`
	}

	// Listen configuration for address and port
	Listen struct {
		HTTP string `yaml:"http"`
		GRPC string `yaml:"grpc"`
	}

	Metrics struct {
		HTTP string `yaml:"http"`
	}
	Tracing struct {
		Endpoint string `yaml:"endpoint"`
	}
)

func ReadYaml(path string, target interface{}) error {
	confBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read config from path=%s: %w", path, err)
	}
	if len(confBytes) == 0 {
		return fmt.Errorf("empty config file")
	}

	if err := UnmarshalYaml(confBytes, target); err != nil {
		return fmt.Errorf(`parse config: %w`, err)
	}

	validate := validator.New()

	err = validate.Struct(target)
	if err != nil {
		return fmt.Errorf("validate config: %w", err)
	}

	return nil
}

func UnmarshalYaml(content []byte, target interface{}) error {
	return yaml.UnmarshalStrict([]byte(os.ExpandEnv(string(content))), target)
}
