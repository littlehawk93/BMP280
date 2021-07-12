package bmp280

type register uint8

const (
	regReset        register = 0xE0
	regStatus       register = 0xF3
	regMesasureCtrl register = 0xF4
	regConfig       register = 0xF5
	regPressMSB     register = 0xF7
	regPressLSB     register = 0xF8
	regPressXLSB    register = 0xF9
	regTempMSB      register = 0xFA
	regTempLSB      register = 0xFB
	regTempXLSB     register = 0xFC

	regCalibrationT1LSB register = 0x88
	regCalibrationT1MSB register = 0x89
	regCalibrationT2LSB register = 0x8A
	regCalibrationT2MSB register = 0x8B
	regCalibrationT3LSB register = 0x8C
	regCalibrationT3MSB register = 0x8D

	regCalibrationP1MSG register = 0x8E
)

var regsPressure = []register{
	regPressMSB,
	regPressLSB,
	regPressXLSB,
}

var regsTemperature = []register{
	regTempMSB,
	regTempLSB,
	regTempXLSB,
}

var regsTempCalibration = []register{
	regCalibrationT1LSB,
	regCalibrationT1MSB,
	regCalibrationT2LSB,
	regCalibrationT2MSB,
	regCalibrationT3LSB,
	regCalibrationT3MSB,
}
