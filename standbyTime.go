package bmp280

// StandbyTime defines how long between measurements during NormalMode operation
type StandbyTime uint8

const (
	// Standby500u 500 microseconds
	Standby500u StandbyTime = 0x00

	// Standby062 62.5 milliseconds
	Standby062 StandbyTime = 0x01

	// Standby125 125 milliseconds
	Standby125 StandbyTime = 0x02

	// Standby250 250 milliseconds
	Standby250 StandbyTime = 0x03

	// Standby500 500 milliseconds
	Standby500 StandbyTime = 0x04

	// Standby1s 1 second
	Standby1s StandbyTime = 0x05

	// Standby2s 2 seconds
	Standby2s StandbyTime = 0x06

	// Standby4s 4 seconds
	Standby4s StandbyTime = 0x07
)
