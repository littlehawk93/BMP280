package bmp280

type calibrationData float64

func (me calibrationData) f() float64 {
	return float64(me)
}

func newCalibrationData(b []byte, offset int, unsigned bool) calibrationData {

	if offset+2 > len(b) {
		panic("Offset and length must fit within size of slice")
	}

	if unsigned {
		var intVal uint16
		intVal = uint16(b[offset+1])<<8 | uint16(b[offset])&0xFF
		return calibrationData(intVal)
	}

	var intVal int16
	intVal = int16(b[offset+1])<<8 | int16(b[offset])&0xFF
	return calibrationData(intVal)
}
