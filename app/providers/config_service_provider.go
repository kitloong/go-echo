package providers

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
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
	p.injectEnv()
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

// Inject value from dotenv into yaml file
// Find ${PATTERN} and replace with dotenv value
// If ${PATTERN} is not defined in dotenv, set key to empty
func (p *ConfigServiceProvider) injectEnv() {
	for _, key := range viper.AllKeys() {
		rgx := regexp.MustCompile(`\${(.*)}`)
		rs := rgx.FindStringSubmatch(viper.GetString(key))
		if len(rs) > 0 {
			if viper.Get(rs[1]) != nil {
				viper.Set(key, viper.Get(rs[1]))
			} else {
				viper.Set(key, "")
			}
		}
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
