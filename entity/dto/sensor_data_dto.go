package dto

import "fmt"

type SensorDataDto struct {
  Temperature float32 `json:"temperature"`
  Luminosity int32 `json:"luminosity"`
}

func (s *SensorDataDto) Validate() error {
  if s.Luminosity > 4066 {
    return fmt.Errorf("luminosity value above maximum expected")
  }

  return nil
}