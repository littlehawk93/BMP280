package bmp280

const (
	filteringMask uint8 = 0x03
	standbyMask   uint8 = 0x03
)

// Configuration additional configuration options for the BMP280 sensor
type Configuration struct {
	StandbyTime  StandbyTime
	IIRFiltering FilteringCoefficient
}

// UInt8 returns the unsigned byte value for this configuration's properties
func (me Configuration) UInt8() uint8 {

	return uint8(((standbyMask & uint8(me.StandbyTime)) << 5) | ((filteringMask & uint8(me.IIRFiltering)) << 2) | 0x00)
}
