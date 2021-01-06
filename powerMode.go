package bmp280

// PowerMode indicates what power mode the BMP sensor should operate in
type PowerMode uint8

const (
	// ModeSleep No measurements are performed. Power consumption is at a minimum. All registers are accessible
	ModeSleep PowerMode = 0x00

	// ModeForced Single measurements are performed according to the select measurement and filter options. Sensor returns to sleep mode after measurement.
	ModeForced PowerMode = 0x01

	// ModeNormal Continuously cycles between an active measurement period and inactive standby period (defined by Standby Time).
	ModeNormal PowerMode = 0x03
)
