package providers

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ConfigServiceProvider read yaml file and load into config.Get
type ConfigServiceProvider struct {
}

const (
	configPath      = "./config"
	configExtension = ".yaml"
	dotEnvPath      = "./.env"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Boot service
func (p *ConfigServiceProvider) Boot(e *echo.Echo) {
	// Load server env
	viper.AutomaticEnv()

	p.loadDotEnv()

	p.loadYaml()
}

func (p *ConfigServiceProvider) loadDotEnv() {
	viper.SetConfigType("env")
	env := p.getFileContent(dotEnvPath)

	err := viper.MergeConfig(bytes.NewBuffer(env))
	check(err)
}

func (p *ConfigServiceProvider) loadYaml() {
	viper.SetConfigType("yaml")

	files := p.getYamlFiles()

	for _, file := range files {
		content := p.getFileContent(configPath + string(os.PathSeparator) + file.Name())

		err := viper.MergeConfig(bytes.NewBuffer(content))
		check(err)
	}
}

// Get all yaml files from configPath
func (p *ConfigServiceProvider) getYamlFiles() []os.FileInfo {
	var files []os.FileInfo
	err := filepath.Walk(configPath, func(path string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if file.IsDir() {
			return nil
		}
		if filepath.Ext(path) == configExtension {
			files = append(files, file)
		}
		return nil
	})
	if err != nil {
		return nil
	}

	return files
}

func (p *ConfigServiceProvider) getFileContent(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	check(err)

	return content
}
