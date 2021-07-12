package bmp280

type calibrationData uint16

func readBytes(b []byte, offset int) calibrationData {

	if offset+1 < len(b)-1 {
		panic("Not enough bytes remaining in slice to parse calibration data")
	}

	var d calibrationData = 0

	d |= calibrationData(0xFF & b[offset])
	d <<= 8
	d |= calibrationData(0xFF & b[offset+1])
	return d
}
