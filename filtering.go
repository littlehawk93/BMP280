package bmp280

// FilteringCoefficient options for filtering data to smooth sudden spikes in sensor readings
type FilteringCoefficient uint8

const (
	// FilteringOff no filtering. Each value is the last raw value read from the sensor
	FilteringOff FilteringCoefficient = 0x00

	// Filtering2 Slight filtering. Takes about 2 samples to start seeing significant changes in sensor values
	Filtering2 FilteringCoefficient = 0x02

	// Filtering4 Some filtering. Takes about 5 samples to start seeing significant changes in sensor values
	Filtering4 FilteringCoefficient = 0x03

	// Filtering8 High filtering. Takes about 11 samples to start seeing significant changes in sensor values
	Filtering8 FilteringCoefficient = 0x04

	// Filtering16 Very high filtering. Takes about 22 samples to start seeing significant changes in sensor values
	Filtering16 FilteringCoefficient = 0x05
)
