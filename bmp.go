package bmp280

import (
	"fmt"
	"math"

	"periph.io/x/periph/conn/i2c"
)

// BMP a BMP280 sensor
type BMP struct {
	device                     *i2c.Dev
	config                     *Configuration
	lastTemperatureMeasurement float64
	tempCalibration            []calibrationData
	pressCalibration           []calibrationData
}

// NewBMP280 create an initialize a new BMP280 sensor. Returns the newly created BMP280 sensor
func NewBMP280(dev *i2c.Dev) (*BMP, error) {
	bmp := &BMP{
		device:                     dev,
		tempCalibration:            make([]calibrationData, 3),
		pressCalibration:           make([]calibrationData, 9),
		lastTemperatureMeasurement: math.NaN(),
	}

	err := bmp.Init()
	return bmp, err
}

// NewBMP280WithConfig create an initialize a new BMP280 sensor and set up its initial configuration. Returns the newly created BMP280 sensor or any errors that occurred when configuring it.
func NewBMP280WithConfig(dev *i2c.Dev, cfg *Configuration) (*BMP, error) {

	bmp := &BMP{
		device:                     dev,
		config:                     cfg,
		tempCalibration:            make([]calibrationData, 3),
		pressCalibration:           make([]calibrationData, 9),
		lastTemperatureMeasurement: math.NaN(),
	}

	if err := bmp.SetConfiguration(cfg); err != nil {
		return bmp, err
	}

	err := bmp.Init()
	return bmp, err
}

// Init initializes the BMP280 sensor with calibration data saved in the BMP280's registers
func (me *BMP) Init() error {

	d, err := me.readCalibrationValues(regsTempCalibration...)

	if err != nil {
		return fmt.Errorf("Error reading temperature calibration values: %s", err.Error())
	}

	me.tempCalibration = d

	if d, err = me.readCalibrationValues(regsPressureCalibration...); err != nil {
		return fmt.Errorf("Error reading pressure calibration values: %s", err.Error())
	}

	me.pressCalibration = d
	return nil
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

// ReadAll reads both temperature and pressure sensor data at once. Returns temperature value, pressure value, or any errors that occurred
func (me *BMP) ReadAll() (float64, float64, error) {

	intVals, err := me.readSensorValues(regTemperature, regPressure)

	if err != nil {
		return math.NaN(), math.NaN(), err
	}

	temp := me.computeTemperature(intVals[0])
	press := me.computePressure(intVals[1])
	return temp, press, nil
}

// ReadTemperature read the latest temperature reading from the BMP280. Returns the temperature in Celsius or any i2c errors that occurred
func (me *BMP) ReadTemperature() (float64, error) {

	intVals, err := me.readSensorValues(regTemperature)

	if err != nil {
		return math.NaN(), err
	}

	return me.computeTemperature(intVals[0]), nil
}

// ReadPressure read the latest pressure reading from the BMP280. Returns the pressure in Pascals or any i2c errors that occurred
func (me *BMP) ReadPressure() (float64, error) {

	if math.IsNaN(float64(me.lastTemperatureMeasurement)) {
		if _, err := me.ReadTemperature(); err != nil {
			return float64(math.NaN()), err
		}
	}

	vals, err := me.readSensorValues(regPressure)

	if err != nil {
		return math.NaN(), err
	}
	return me.computePressure(vals[0]), nil
}

func (me *BMP) computeTemperature(intVal sensorData) float64 {

	floatVal := float64(intVal)

	var1 := (floatVal/16384.0 - me.pressCalibration[0].f()/1024.0) * me.pressCalibration[1].f()
	var2 := ((floatVal/131072.0 - me.pressCalibration[0].f()/8192.0) * (floatVal/131072.0 - me.pressCalibration[0].f()/8192.0)) * me.pressCalibration[2].f()

	me.lastTemperatureMeasurement = var1 + var2
	return me.lastTemperatureMeasurement / 5120.0
}

func (me *BMP) computePressure(intVal sensorData) float64 {

	floatVal := float64(intVal)

	var1 := me.lastTemperatureMeasurement/2.0 - 64000.0
	var2 := var1 * var1 * me.pressCalibration[5].f() / 32768.0
	var2 = var2 + var1*me.pressCalibration[4].f()*2.0
	var2 = var2/4.0 + me.pressCalibration[3].f()*65536.0
	var1 = (me.pressCalibration[2].f()*var1*var1/524288.0 + me.pressCalibration[1].f()*var1) / 524288.0
	var1 = (1.0 + var1/32768.0) * me.pressCalibration[0].f()

	if var1 == 0.0 {
		return 0.0
	}

	p := 1048576.0 - floatVal
	p = (p - (var2 / 4096.0)) * 6250.0 / var1
	var1 = me.pressCalibration[8].f() * p * p / 2147483648.0
	var2 = p * me.pressCalibration[7].f() / 32768.0
	p = p + (var1+var2+me.pressCalibration[6].f())/16.0
	return p
}

func (me *BMP) readCalibrationValues(registers ...register) ([]calibrationData, error) {

	b, err := me.readRegisters(2, registers...)

	if err != nil {
		return nil, err
	}

	vals := make([]calibrationData, len(registers))

	for i := 0; i < len(b); i += 2 {
		vals[i/2] = newCalibrationData(b, i, i == 0)
	}

	return vals, nil
}

func (me *BMP) readSensorValues(registers ...register) ([]sensorData, error) {

	b, err := me.readRegisters(3, registers...)

	if err != nil {
		return nil, err
	}

	vals := make([]sensorData, len(registers))

	for i := 0; i < len(b); i += 3 {
		vals[i/3] = newSensorData(b, i)
	}
	return vals, nil
}

func (me *BMP) readRegisters(size int, registers ...register) ([]byte, error) {

	w := make([]byte, len(registers)*size)
	r := make([]byte, len(w))

	if err := me.device.Bus.Tx(me.device.Addr, w, r); err != nil {
		return nil, err
	}
	return r, nil
}

func (me *BMP) convertUInt16(b []byte, offset int) uint16 {

	if offset+2 > len(b) {
		panic("Offset and length must fit within size of slice")
	}

	return uint16(b[offset+1])<<8 | (0xFF & uint16(b[offset]))
}

func (me *BMP) convertInt16(b []byte, offset int) int16 {

	if offset+2 > len(b) {
		panic("Offset and length must fit within size of slice")
	}

	return int16(b[offset+1])<<8 | (0xFF & int16(b[offset]))
}
