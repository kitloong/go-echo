package config

import (
	"bytes"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

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

// Load config files
func Load() {
	// Load server env
	viper.AutomaticEnv()

	loadDotEnv()
	loadYaml()
	injectEnv()
}

func loadDotEnv() {
	viper.SetConfigType("env")
	env := getFileContent(dotEnvPath)

	err := viper.MergeConfig(bytes.NewBuffer(env))
	check(err)
}

func loadYaml() {
	viper.SetConfigType("yaml")

	files := getYamlFiles()

	for _, file := range files {
		content := getFileContent(configPath + string(os.PathSeparator) + file.Name())

		err := viper.MergeConfig(bytes.NewBuffer(content))
		check(err)
	}
}

// Inject value from dotenv into yaml file
// Find ${PATTERN} and replace with dotenv value
// If ${PATTERN} is not defined in dotenv, set key to empty
func injectEnv() {
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
func getYamlFiles() []os.FileInfo {
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

func getFileContent(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	check(err)

	return content
}
