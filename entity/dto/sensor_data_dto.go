package dto

import "fmt"

type SensorDataDto struct {
	Distance   int32 `json:"distance"`
	Luminosity int32 `json:"luminosity"`
}

func (s *SensorDataDto) Validate() error {
	if s.Luminosity > 4096 {
		return fmt.Errorf("luminosity value above maximum expected")
	}

	return nil
}
