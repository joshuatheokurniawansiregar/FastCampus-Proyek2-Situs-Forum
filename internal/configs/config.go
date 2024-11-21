package configs

import (
	"github.com/spf13/viper"
)

var config *Config

type option struct {
	configFolders []string
	configFile    string
	configType    string
}

func Init(optFuncs ...OptionFunc) error {
	var option_ *option = &option{
		configFolders: getDefaultConfigFolder(),
		configFile:    getDefaultConfigFile(),
		configType:    getDefaultConfigType(),
	}

	for _, optFunc := range optFuncs {
		optFunc(option_)
	}

	for _, configFolder := range option_.configFolders {
		viper.AddConfigPath(configFolder)
	}

	viper.SetConfigName(option_.configFile)
	viper.SetConfigType(option_.configType)
	viper.AutomaticEnv()

	config = new(Config)

	var err error = viper.ReadInConfig()
	if err != nil {
		return err
	}
	return viper.Unmarshal(&config)
}

type OptionFunc func(*option)

func getDefaultConfigFolder() []string {
	return []string{"./internal/configs/"}
}

func getDefaultConfigFile() string {
	return "config"
}

func getDefaultConfigType() string {
	return "yaml"
}

func WithConfigFolder(configFolders []string) OptionFunc {
	return func(opt *option) {
		opt.configFolders = configFolders
	}
}

func WithConfigFile(configFile string) OptionFunc {
	return func(opt *option) {
		opt.configFile = configFile
	}
}

func WithConfigType(configType string) OptionFunc {
	return func(opt *option) {
		opt.configType = configType
	}
}

func GetConfig() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}
