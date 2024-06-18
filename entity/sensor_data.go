package entity

type SensorData struct {
	Distance   int32
	Luminosity int32
}

func NewSensorData(temperature int32, luminosity int32) *SensorData {
	return &SensorData{
		Distance:   temperature,
		Luminosity: luminosity,
	}
}
