package bmp280

type PowerMode uint8

const (
	ModeSleep  PowerMode = 0x00
	ModeForced PowerMode = 0x01
	ModeNormal PowerMode = 0x03
)
