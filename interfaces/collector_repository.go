package interfaces

import (
	"github.com/luanc202/timeseriesSensorCollector/entity/dto"
)

type CollectorRepository interface {
  Save(*dto.SensorDataDto) (error)
}