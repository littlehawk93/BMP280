package bmp280

// Oversampling defines how much sampling do to for a measurement. Higher sampling means higher precision, but slower measurement speeds
type Oversampling uint8

const (
	// OversampleSkip skip this measurement
	OversampleSkip Oversampling = 0x00

	// OversampleX1 16 bit resolution (2.62 Pa or 0.0050 C)
	OversampleX1 Oversampling = 0x01

	// OversampleX2 17 bit resolution (1.31 Pa or 0.0025 C)
	OversampleX2 Oversampling = 0x02

	// OversampleX4 18 bit resolution (0.66 Pa or 0.0012 C)
	OversampleX4 Oversampling = 0x03

	// OversampleX8 19 bit resolution (20.33 Pa or 0.0006 C)
	OversampleX8 Oversampling = 0x04

	// OversampleX16 20 bit resolution (0.16 Pa or 0.0003 C)
	OversampleX16 Oversampling = 0x05
)
