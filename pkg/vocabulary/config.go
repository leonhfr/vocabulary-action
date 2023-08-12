package vocabulary

import "github.com/spf13/viper"

const (
	configFile       = "vocabulary"
	defaultLanguage  = "default"
	defaultDirectory = "todo/vocabulary"
)

type config map[string]string

func (cfg config) languageDirectory(language string) string {
	directory, ok := cfg[language]
	if ok {
		return directory
	}

	directory, ok = cfg[defaultLanguage]
	if ok {
		return directory
	}

	return defaultDirectory
}

func newConfig(workspace string) config {
	viper.AddConfigPath(workspace)
	viper.SetConfigName(configFile)

	err := viper.ReadInConfig()
	if err != nil {
		return config{}
	}

	languages := viper.GetStringMapString("languages")

	return config(languages)
}
