package bmp280

type StandbyTime uint8

const (
	Standby500u StandbyTime = 0x00
	Standby062  StandbyTime = 0x01
	Standby125  StandbyTime = 0x02
	Standby250  StandbyTime = 0x03
	Standby500  StandbyTime = 0x04
	Standby1s   StandbyTime = 0x05
	Standby2s   StandbyTime = 0x06
	Standby4s   StandbyTime = 0x07
)
