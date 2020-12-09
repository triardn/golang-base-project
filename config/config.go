package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Provider the config provider
type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

var defaultConfig *viper.Viper

func init() {
	defaultConfig = readViperConfig()
}

func readViperConfig() *viper.Viper {
	v := viper.New()
	v.AddConfigPath(".")
	v.AddConfigPath("../../../params")
	v.AddConfigPath("./params")
	v.AddConfigPath("/opt/golang-base-project/params")
	v.SetConfigName("golang-base-project")

	err := v.ReadInConfig()
	if err == nil {
		fmt.Printf("Using config file: %s \n\n", v.ConfigFileUsed())
	} else {
		panic(fmt.Errorf("ProviderConfig error: %s", err))
	}

	return v
}

// NewAppConfig initiate app config
func NewAppConfig(c *viper.Viper) (conf AppConfig, err error) {
	err = c.Unmarshal(&conf)
	if err != nil {
		return
	}

	return
}

// ProviderConfig return provider so that you can read config anywhere
func ProviderConfig() *viper.Viper {
	return defaultConfig
}

// AppConfig define app configuration
type AppConfig struct {
}
