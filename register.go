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
