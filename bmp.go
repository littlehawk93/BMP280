package bmp280

import "periph.io/x/periph/conn/i2c"

// BMP a BMP280 sensor
type BMP struct {
	device *i2c.Dev
}

// NewBMP280 create an initialize a new BMP280 sensor. Returns the newly created BMP280 sensor
func NewBMP280(dev *i2c.Dev) *BMP {
	return &BMP{
		device: dev,
	}
}

// NewBMP280WithConfig create an initialize a new BMP280 sensor and set up its initial configuration. Returns the newly created BMP280 sensor or any errors that occurred when configuring it.
func NewBMP280WithConfig(dev *i2c.Dev, cfg *Configuration) (*BMP, error) {

	bmp := NewBMP280(dev)

	err := bmp.SetConfiguration(cfg)

	if err != nil {
		return nil, err
	}
	return bmp, nil
}

// SetConfiguration set up the configuration for how this BMP280 sensor should operate. Returns any errors that occur when configuring the sensor
func (me *BMP) SetConfiguration(cfg *Configuration) error {

	_, err := me.device.Write(cfg.Bytes())
	return err
}
