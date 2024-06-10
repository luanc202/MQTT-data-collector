package db

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/luanc202/timeseriesSensorCollector/entity"
	"github.com/luanc202/timeseriesSensorCollector/entity/dto"
	"github.com/luanc202/timeseriesSensorCollector/infra/config"
	"github.com/luanc202/timeseriesSensorCollector/interfaces"
)

var logger = config.GetLogger("collector-repository")

type CollectorRepository struct {
	writeAPI api.WriteAPIBlocking
	queryAPI api.QueryAPI
	client influxdb2.Client
}

func NewCollectorRepository(writeAPIconnection api.WriteAPIBlocking, queryAPIconnection api.QueryAPI) interfaces.CollectorRepository {
	return &CollectorRepository{
		writeAPI: writeAPIconnection,
    queryAPI: queryAPIconnection,
	}
}

func (cr *CollectorRepository) Save(measuresDto *dto.SensorDataDto) error {

	dataToSave := entity.NewSensorData(measuresDto.Temperature, measuresDto.Luminosity)

	point := influxdb2.NewPoint("sensor",
	map[string]string{"device": "ESP32"},
	map[string]interface{}{
		"temperature": dataToSave.Temperature, 
		"luminosity": dataToSave.Luminosity,
	},
	time.Now())

	err := cr.writeAPI.WritePoint(context.Background(), point)

	if err != nil {
		return fmt.Errorf("error on saving data: %w", err)
	}

	defer cr.client.Close()

	logger.Info(fmt.Printf("Saved new entry with values for temperature: %f and luminosity: %d", measuresDto.Temperature, measuresDto.Luminosity))

	return nil
}
