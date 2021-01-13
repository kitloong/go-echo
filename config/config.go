package config

import "github.com/spf13/viper"

// Get configuration by key
func Get(key string) interface{} {
	return viper.Get(key)
}
