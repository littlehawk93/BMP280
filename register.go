package bmp280

type register uint8

const (
	regReset        register = 0xE0
	regStatus       register = 0xF3
	regMesasureCtrl register = 0xF4
	regConfig       register = 0xF5

	regPressMSB  register = 0xF7
	regPressLSB  register = 0xF8
	regPressXLSB register = 0xF9

	regPressure register = regPressMSB

	regTempMSB  register = 0xFA
	regTempLSB  register = 0xFB
	regTempXLSB register = 0xFC

	regTemperature register = regTempMSB

	regCalibrationT1LSB register = 0x88
	regCalibrationT1MSB register = 0x89
	regCalibrationT2LSB register = 0x8A
	regCalibrationT2MSB register = 0x8B
	regCalibrationT3LSB register = 0x8C
	regCalibrationT3MSB register = 0x8D

	regCalibrationT1 register = regCalibrationT1LSB
	regCalibrationT2 register = regCalibrationT2LSB
	regCalibrationT3 register = regCalibrationT3LSB

	regCalibrationP1LSB register = 0x8E
	regCalibrationP1MSB register = 0x8F
	regCalibrationP2LSB register = 0x90
	regCalibrationP2MSB register = 0x91
	regCalibrationP3LSB register = 0x92
	regCalibrationP3MSB register = 0x93
	regCalibrationP4LSB register = 0x94
	regCalibrationP4MSB register = 0x95
	regCalibrationP5LSB register = 0x96
	regCalibrationP5MSB register = 0x97
	regCalibrationP6LSB register = 0x98
	regCalibrationP6MSB register = 0x99
	regCalibrationP7LSB register = 0x9A
	regCalibrationP7MSB register = 0x9B
	regCalibrationP8LSB register = 0x9C
	regCalibrationP8MSB register = 0x9D
	regCalibrationP9LSB register = 0x9E
	regCalibrationP9MSB register = 0x9F

	regCalibrationP1 register = regCalibrationP1LSB
	regCalibrationP2 register = regCalibrationP2LSB
	regCalibrationP3 register = regCalibrationP3LSB
	regCalibrationP4 register = regCalibrationP4LSB
	regCalibrationP5 register = regCalibrationP5LSB
	regCalibrationP6 register = regCalibrationP6LSB
	regCalibrationP7 register = regCalibrationP7LSB
	regCalibrationP8 register = regCalibrationP8LSB
	regCalibrationP9 register = regCalibrationP9LSB
)

var regsTempCalibration = []register{
	regCalibrationT1,
	regCalibrationT2,
	regCalibrationT3,
}

var regsPressureCalibration = []register{
	regCalibrationP1,
	regCalibrationP2,
	regCalibrationP3,
	regCalibrationP4,
	regCalibrationP5,
	regCalibrationP6,
	regCalibrationP7,
	regCalibrationP8,
	regCalibrationP9,
}

func (me register) incrementRegisters(b []byte, offset, length int) {

	if offset+length > len(b) {
		panic("Offset and length must fit within size of slice")
	}

	for i := 0; i < length; i++ {
		b[offset+i] = byte(me) + byte(i)
	}
}
