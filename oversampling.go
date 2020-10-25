package bmp280

type Oversampling uint8

const (
	OversampleSkip Oversampling = 0x00
	OversampleX1   Oversampling = 0x01
	OversampleX2   Oversampling = 0x02
	OversampleX4   Oversampling = 0x03
	OversampleX8   Oversampling = 0x04
	OversampleX16  Oversampling = 0x05
)
