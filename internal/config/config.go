package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

const (
	prefix = "chameleon"
)

type Config interface {

	// SetDefaults will set the default value for any values which have not already been set and return an error when
	//	a required value has not been set
	SetDefaults() error
}

func LoadConfig(appName string, cfg Config) (err error) {

	// Config Priority:
	// 1. Command-line flags (Configured in SetupFlagsForConfig)
	// 2. Environment variables
	// 3. Config file

	// Environment variables
	viper.SetEnvPrefix(fmt.Sprintf("%s_%s", prefix, appName))
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.AutomaticEnv()

	// Config file
	viper.SetConfigName(appName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/chameleon")
	viper.AddConfigPath("../configs")
	viper.AddConfigPath(".")

	// Load in the configuration
	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error
			fmt.Println("Config file not found")
		} else {
			return err
		}
	}

	// Unmarshal
	err = viper.Unmarshal(cfg)
	if err != nil {
		return err
	}

	// Ensure defaults are set
	err = cfg.SetDefaults()
	if err != nil {
		return err
	}

	return nil
}