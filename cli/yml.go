package cli

import (
	"os"

	"gopkg.in/yaml.v3"
)

var defaultConfigYml = map[string]string{
	"auth":       "token",
	"model":      "text-davinci-003",
	"max_tokens": "256",
}

var configFolderName = "/cligpt/"
var configFileName = "cligpt.yml"

func getConfigDir() (string, error) {
	var userConfigDir, err = os.UserConfigDir()
	return userConfigDir + configFolderName, err
}

func CreateConfigDirectory() error {
	var err error
	var cligptConfigDir string
	if cligptConfigDir, err = getConfigDir(); err != nil {
		return err
	}
	var cligptConfigFile = cligptConfigDir + configFileName

	os.MkdirAll(cligptConfigDir, 0755)

	if _, err := os.Stat(cligptConfigFile); os.IsExist(err) {
		return err
	}
	file, err := os.Create(cligptConfigFile)
	if err != nil {
		return err
	}
	defer file.Close()

	var bffYml = yaml.NewEncoder(file)
	defer bffYml.Close()
	return bffYml.Encode(defaultConfigYml)
}

func ReadYml() (map[string]string, error) {
	var err error
	var cligptConfigDir string
	if cligptConfigDir, err = getConfigDir(); err != nil {
		return nil, err
	}
	var cligptConfigFile = cligptConfigDir + configFileName

	file, err := os.Open(cligptConfigFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var config map[string]string
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}
