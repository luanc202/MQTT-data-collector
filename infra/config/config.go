package config

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

var (
	writeAPI           api.WriteAPIBlocking
	queryAPI           api.QueryAPI
	logger             *Logger
	StandardDateLayout = "2006-01-02"
	influxdb2Client    influxdb2.Client
	mqttClient         mqtt.Client
)

func InitConfigs() error {
	env := GetEnvConfig()

	options := influxdb2.DefaultOptions()
	influxdb2Client = influxdb2.NewClientWithOptions(env.DBUrl, env.InfluxDBToken, options)
	writeAPI = influxdb2Client.WriteAPIBlocking(env.InfluxDB_Org, env.InfluxDB_Bucket)
	logger.Info(fmt.Printf("Initialized writeAPI on Bucket %s and Org %s", env.InfluxDB_Bucket, env.InfluxDB_Org))
	queryAPI = influxdb2Client.QueryAPI(env.InfluxDB_Org)

	mqttOptions := mqtt.NewClientOptions().AddBroker(env.MQTT_Broker_url)
	mqttClient = mqtt.NewClient(mqttOptions)

	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

func GetWriteAPI() api.WriteAPIBlocking {
	return writeAPI
}

func GetQueryAPI() api.QueryAPI {
	return queryAPI
}

func GetMqttClient() mqtt.Client {
	return mqttClient
}

func GetInfluxDBClient() influxdb2.Client {
	return influxdb2Client
}
