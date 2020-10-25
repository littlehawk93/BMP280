package bmp280

// FilteringCoefficient options for filtering data to smooth sudden spikes in sensor readings
type FilteringCoefficient uint8

const (
	filteringOff FilteringCoefficient = 0x00
	filtering2   FilteringCoefficient = 0x02
	filtering4   FilteringCoefficient = 0x03
	filtering8   FilteringCoefficient = 0x04
	filtering16  FilteringCoefficient = 0x05
)
