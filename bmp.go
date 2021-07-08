package bmp280

import (
	"math"

	"periph.io/x/periph/conn/i2c"
)

// BMP a BMP280 sensor
type BMP struct {
	device          *i2c.Dev
	config          *Configuration
	lastTempReading float32
}

// NewBMP280 create an initialize a new BMP280 sensor. Returns the newly created BMP280 sensor
func NewBMP280(dev *i2c.Dev) (*BMP, error) {
	bmp := &BMP{
		device:          dev,
		lastTempReading: float32(math.NaN()),
	}

	err := bmp.SetConfiguration(ConfigDefault())
	return bmp, err
}

// NewBMP280WithConfig create an initialize a new BMP280 sensor and set up its initial configuration. Returns the newly created BMP280 sensor or any errors that occurred when configuring it.
func NewBMP280WithConfig(dev *i2c.Dev, cfg *Configuration) (*BMP, error) {

	bmp := &BMP{
		device:          dev,
		config:          cfg,
		lastTempReading: float32(math.NaN()),
	}

	err := bmp.SetConfiguration(cfg)
	return bmp, err
}

// Configuration get the current configuration for this BMP280 sensor
func (me *BMP) Configuration() *Configuration {

	return me.config
}

// SetConfiguration set up the configuration for how this BMP280 sensor should operate. Returns any errors that occur when configuring the sensor
func (me *BMP) SetConfiguration(cfg *Configuration) error {

	me.config = cfg
	_, err := me.device.Write(cfg.bytes())
	return err
}

// Reset perform a full reset on the BMP280 device as if it had first powered on
func (me *BMP) Reset() error {

	_, err := me.device.Write([]byte{uint8(regReset), resetCode})
	return err
}

// ReadTemperature read the latest temperature reading from the BMP280. Returns the temperature in Celsius or any i2c errors that occurred
func (me *BMP) ReadTemperature() (float32, error) {

	return 0.0, nil
}

// ReadPressure read the latest pressure reading from the BMP280. Returns the pressure in Pascals or any i2c errors that occurred
func (me *BMP) ReadPressure() (float32, error) {

	return 0.0, nil
}

func readRaw(registers []byte, numBits int) (uint, error) {

	return 0, nil
}
