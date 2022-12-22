package cli

import (
	"os"

	"gopkg.in/yaml.v3"
)

func CreateConfigDirectory() {
	os.MkdirAll(os.Getenv("HOME")+"/.config/cligpt/", 0755)
	if _, err := os.Stat(os.Getenv("HOME") + "/.config/cligpt/cligpt.yml"); os.IsNotExist(err) {
		file, err := os.Create(os.Getenv("HOME") + "/.config/cligpt/cligpt.yml")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		config := `auth: token
model: text-davinci-003
max_tokens: 256`
		file.WriteString(config)
	}
}

func ReadFromYml(variable string) string {
	file, err := os.Open(os.Getenv("HOME") + "/.config/cligpt/cligpt.yml")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var config map[string]string
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	return string(config[variable])
}
