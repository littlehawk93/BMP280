package bmp280

const (
	oversamplingMask uint8 = 0x07
	powerModeMask    uint8 = 0x03
)

// MeasurementControl defines sampling and control options for how the BMP280 makes measurements
type MeasurementControl struct {
	TemperatureOversampling Oversampling
	PressureOversampling    Oversampling
	Mode                    PowerMode
}

// UInt8 returns the unsigned byte value for this measurement mode's properties
func (me MeasurementControl) UInt8() uint8 {

	return uint8(((uint8(me.TemperatureOversampling) & oversamplingMask) << 5) | ((uint8(me.PressureOversampling) & oversamplingMask) << 2) | (uint8(me.Mode) & powerModeMask))
}
