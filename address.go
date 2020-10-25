package bmp280

const (
	// DefaultAddress the default I2C address BMP280 uses when the SDO pin is grounded
	DefaultAddress uint8 = 0x76

	// LegacyAddress the default I2C address for BMP180. BMP280 can also have this I2C address when SDO pin is connected to Vddio
	LegacyAddress uint8 = 0x77
)
