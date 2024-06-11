package config

import "github.com/spf13/viper"

var env *envconfig

type envconfig struct {
	DBUrl           string `mapstructure:"DATABASE_URL"`
	PORT            string `mapstructure:"PORT"`
	ENV             string `mapstructure:"ENVIRONMENT"`
	InfluxDBToken   string `mapstructure:"INFLUXDB_TOKEN"`
	InfluxDB_Org    string `mapstructure:"DOCKER_INFLUXDB_INIT_ORG"`
	InfluxDB_Bucket string `mapstructure:"DOCKER_INFLUXDB_INIT_BUCKET"`
	MQTT_Broker_url string `mapstructure:"MQTT_BROKER_URL"`
	MQTT_Topic      string `mapstructure:"MQTT_TOPIC"`
}

func GetEnvConfig() *envconfig {
	return env
}

func LoadEnv(path string) (*envconfig, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		panic(err)
	}

	return env, nil
}
