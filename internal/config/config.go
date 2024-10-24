package config

import "os"

type Config struct {
	MongoURL string
	MongoDB  string
}

func LoadConfig() Config {
	return Config{
		MongoURL: os.Getenv("MONGO_URL"),
		MongoDB:  os.Getenv("MONGO_DB"),
	}
}
