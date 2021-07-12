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

	intVal, err := me.readRawSensorValue(me.config.TemperatureOversampling.GetNumBits(), regsTemperature...)

	if err != nil {
		return float32(math.NaN()), err
	}

	floatVal := float32(intVal)

	v1 := (floatVal / 16384.0 - 

	return float32(intVal), nil
}

// ReadPressure read the latest pressure reading from the BMP280. Returns the pressure in Pascals or any i2c errors that occurred
func (me *BMP) ReadPressure() (float32, error) {

	return 0.0, nil
}

func (me *BMP) readRawSensorValue(numBits int, registers ...register) (uint, error) {

	if numBits <= 0 {
		return 0, nil
	}

	w := make([]byte, len(registers))

	for i, reg := range registers {
		w[i] = byte(reg)
	}

	r := make([]byte, len(registers))

	if err := me.device.Bus.Tx(me.device.Addr, w, r); err != nil {
		return 0, err
	}

	var num uint = 0

	for i := 0; i < len(r); i++ {
		if numBits-(i*8) < 8 {
			num <<= (numBits - (i * 8))
			num |= uint((0xFF&int(r[i]))>>8 - (numBits - (i * 8)))
		} else {
			num <<= 8
			num |= uint(0xFF & r[i])
		}
	}

	if numBits < 20 {
		num <<= 20 - numBits
	}

	return num, nil
}
