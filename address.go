package bmp280

const (
	// AddressDefault the default I2C address BMP280 uses when the SDO pin is grounded
	AddressDefault uint8 = 0x76

	// AddressLegacy the default I2C address for BMP180. BMP280 can also have this I2C address when SDO pin is connected to Vddio
	AddressLegacy uint8 = 0x77
)
