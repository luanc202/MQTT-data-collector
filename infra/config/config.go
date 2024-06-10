package config

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var (
	db                 influxdb2.Client
	logger             *Logger
	StandardDateLayout = "2006-01-02"
)

func InitConfigs() error {
	var err error
	env := GetEnvConfig()

	db = influxdb2.NewClient(env.DBUrl, env.InfluxDBToken)
	if err != nil {
		panic(err)
	}

	return nil
}

func GetDB() influxdb2.Client {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
