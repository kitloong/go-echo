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
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Boot service
func (p *ConfigServiceProvider) Boot(e *echo.Echo) {
	viper.SetConfigType("yaml")

	files := p.getConfigFiles()

	for _, file := range files {
		content := p.getConfigFileContent(file)

		err := viper.MergeConfig(bytes.NewBuffer(content))
		check(err)
	}
}

func (p *ConfigServiceProvider) getConfigFiles() []os.FileInfo {
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

func (p *ConfigServiceProvider) getConfigFileContent(file os.FileInfo) []byte {
	content, err := ioutil.ReadFile(configPath + string(os.PathSeparator) + file.Name())
	check(err)

	return content
}
