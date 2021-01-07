package bmp280

const (
	resetCode  uint8 = 0xB6
	spiDisable uint8 = 0

	spiMask          uint8 = 0x03
	filteringMask    uint8 = 0x03
	standbyMask      uint8 = 0x03
	oversamplingMask uint8 = 0x07
	powerModeMask    uint8 = 0x03
)

// Configuration additional configuration options for the BMP280 sensor
type Configuration struct {
	StandbyTime             StandbyTime
	IIRFiltering            FilteringCoefficient
	TemperatureOversampling Oversampling
	PressureOversampling    Oversampling
	Mode                    PowerMode
}

func (me Configuration) bytes() []byte {

	return []byte{uint8(regConfig), me.configByte(), uint8(regMesasureCtrl), me.measurementByte()}
}

func (me Configuration) configByte() uint8 {

	return uint8(((standbyMask & uint8(me.StandbyTime)) << 5) | ((filteringMask & uint8(me.IIRFiltering)) << 2) | (spiMask & spiDisable))
}

func (me Configuration) measurementByte() uint8 {

	return uint8(((uint8(me.TemperatureOversampling) & oversamplingMask) << 5) | ((uint8(me.PressureOversampling) & oversamplingMask) << 2) | (uint8(me.Mode) & powerModeMask))
}

// ConfigDefault returns a default configuration for the BMP280 sensor. This configuration offers a decent compromise between power consumption, resolution, and responsiveness
func ConfigDefault() *Configuration {

	return &Configuration{
		StandbyTime:             Standby500u,
		PressureOversampling:    OversampleX4,
		TemperatureOversampling: OversampleX1,
		IIRFiltering:            Filtering16,
		Mode:                    ModeNormal,
	}
}

// ConfigLowPower returns a configuration for low resolution, low speed, ultra-low power consumption applications
func ConfigLowPower() *Configuration {

	return &Configuration{
		StandbyTime:             Standby500u,
		PressureOversampling:    OversampleX1,
		TemperatureOversampling: OversampleX1,
		IIRFiltering:            FilteringOff,
		Mode:                    ModeForced,
	}
}

// ConfigHighSpeed returns a configuration for low resolution, high speed (~125 Hz refresh rate), high power consumption applications.
func ConfigHighSpeed() *Configuration {

	return &Configuration{
		StandbyTime:             Standby500u,
		PressureOversampling:    OversampleX2,
		TemperatureOversampling: OversampleX1,
		IIRFiltering:            FilteringOff,
		Mode:                    ModeNormal,
	}
}
