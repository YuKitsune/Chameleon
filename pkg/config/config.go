package config

type Value struct {
	ConfigKey string
	Flag string
}

type ConfigStore interface {
	GetString(key string) string
	GetStringSlice(key string) []string
	GetBool(key string) bool
	GetInt(key string) int
	GetInt64(key string) int64
}
