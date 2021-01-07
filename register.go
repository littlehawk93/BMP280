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
)

var regsPressure = []byte{uint8(regPressMSB), uint8(regPressLSB), uint8(regPressXLSB)}

var regsTemperature = []byte{uint8(regTempMSB), uint8(regTempLSB), uint8(regTempXLSB)}
