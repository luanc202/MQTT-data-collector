package entity

type SensorData struct {
  Temperature float32 
  Luminosity int32
}

func NewSensorData(temperature float32, luminosity int32) *SensorData {
  return &SensorData{
    Temperature: temperature,
    Luminosity: luminosity,
  }
}