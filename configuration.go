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

// Bytes returns a byte array containing the configuration register and config data serialized for I2C
func (me Configuration) Bytes() []byte {

	return []byte{uint8(regConfig), me.UInt8()}
}

// UInt8 returns the unsigned byte value for this configuration's properties
func (me Configuration) UInt8() uint8 {

	return uint8(((standbyMask & uint8(me.StandbyTime)) << 5) | ((filteringMask & uint8(me.IIRFiltering)) << 2) | 0x00)
}
