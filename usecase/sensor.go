package usecase

import (
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/luanc202/timeseriesSensorCollector/entity/dto"
	"github.com/luanc202/timeseriesSensorCollector/infra/config"
	"github.com/luanc202/timeseriesSensorCollector/interfaces"
)



var loggerUsecase = config.GetLogger("sensor-usecase")

type SensorUsecase struct {
  repo interfaces.CollectorRepository
}

func NewSensorUseCase(repo interfaces.CollectorRepository) *SensorUsecase {
  return &SensorUsecase{repo: repo}
}

func (c *SensorUsecase) Save(msg mqtt.Message) (error) {

  var data dto.SensorDataDto

  err := json.Unmarshal(msg.Payload(), &data)
  if err!= nil {
		return fmt.Errorf("error decoding message payload: %v", err)
	}

  err = data.Validate()

  if err!= nil {
		return fmt.Errorf("error validating data dto: %v", err)
	}

  err = c.repo.Save(&data)

  if err != nil {
    loggerUsecase.Error("error saving data", err)
		return fmt.Errorf("failed to save data: %w", err)
  }

  return nil
}