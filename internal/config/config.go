package config

import (
	"log"
	"path/filepath"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
)

type Env string

const (
	Dev  = "dev"
	Prod = "prod"
	Test = "test"
)

type Config struct {
	DBConfig     DBConfig     `json:"db"`
	ServerConfig ServerConfig `json:"server"`
	Env          Env          `json:"env"`
}

type ServerConfig struct {
	Port int16 `json:"port"`
}

type DBConfig struct {
	Migrate struct {
		Enable bool   `json:"enable"`
		Dir    string `json:"dir"`
	} `json:"migrate"`
	Path string `json:"path"`
}

func Load(configPath string) (*Config, error) {
	k := koanf.New(".")

	// load from default config
	err := k.Load(confmap.Provider(defaultConfig, "."), nil)
	if err != nil {
		log.Printf("failed to load default config. err: %v", err)
		return nil, err
	}

	if configPath != "" {
		path, err := filepath.Abs(configPath)
		if err != nil {
			log.Printf("failed to get absoulute config path. configPath:%s, err: %v", configPath, err)
			return nil, err
		}
		log.Printf("load config file from %s", path)
		if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
			log.Printf("failed to load config from file. err: %v", err)
			return nil, err
		}
	}

	var cfg Config
	if err := k.UnmarshalWithConf("", &cfg, koanf.UnmarshalConf{Tag: "json", FlatPaths: false}); err != nil {
		log.Printf("failed to unmarshal with conf. err: %v", err)
		return nil, err
	}
	return &cfg, err
}
