package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v8"
)

var (
	config *Config
)

func C() *Config {
	if config == nil {
		panic("加载 config 文件失败!")
	}
	return config
}

// 从 env 环境变量中加载配置
func LoadConfigFromEnv() error {
	config = NewDefaultConfig()
	if err := env.Parse(config); err != nil {
		return err
	}
	return nil
}

// 从 toml 文件中加载配置
func LoadConfigFromToml(file_path string) error {
	config = NewDefaultConfig()
	_, err := toml.DecodeFile(file_path, config)
	if err != nil {
		return err
	}
	return nil
}
