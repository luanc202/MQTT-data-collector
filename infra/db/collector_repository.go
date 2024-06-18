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
	client   influxdb2.Client
}

func NewCollectorRepository(writeAPIconnection api.WriteAPIBlocking, queryAPIconnection api.QueryAPI, clientConnection influxdb2.Client) interfaces.CollectorRepository {
	return &CollectorRepository{
		writeAPI: writeAPIconnection,
		queryAPI: queryAPIconnection,
		client:   clientConnection,
	}
}

func (cr *CollectorRepository) Save(measuresDto *dto.SensorDataDto) error {
	dataToSave := entity.NewSensorData(measuresDto.Distance, measuresDto.Luminosity)
	point := influxdb2.NewPoint("sensor",
		map[string]string{"device": "ESP32"},
		map[string]interface{}{
			"distance":   dataToSave.Distance,
			"luminosity": dataToSave.Luminosity,
		},
		time.Now())

	err := cr.writeAPI.WritePoint(context.Background(), point)
	defer cr.client.Close()

	if err != nil {
		return fmt.Errorf("error on saving data: %w", err)
	}

	logger.Info(fmt.Sprintf("Saved new entry with values for distance: %d and luminosity: %d", measuresDto.Distance, measuresDto.Luminosity))

	return nil
}
