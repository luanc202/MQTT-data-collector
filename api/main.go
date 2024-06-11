package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/luanc202/timeseriesSensorCollector/infra/config"
	"github.com/luanc202/timeseriesSensorCollector/infra/db"
	"github.com/luanc202/timeseriesSensorCollector/usecase"
)

func main() {
	env, err := config.LoadEnv(".")
	if err != nil {
		panic(err)
	}

	err = config.InitConfigs()
	if err != nil {
		panic(err)
	}

	mqttClient := config.GetMqttClient()

	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	dbCollectorRepo := db.NewCollectorRepository(config.GetWriteAPI(), config.GetQueryAPI(), config.GetInfluxDBClient())

	sensorUseCase := usecase.NewSensorUseCase(dbCollectorRepo)

	token := mqttClient.Subscribe(env.MQTT_Topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Println(string(msg.Payload()))
		err = sensorUseCase.Save(msg)

		if err != nil {
			fmt.Printf("error on usecase: %s", err.Error())
		}
	})
	token.Wait()

	fmt.Printf("running on port %v \n", env.PORT)

	select {}
}
