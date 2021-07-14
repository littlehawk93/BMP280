package bmp280

type sensorData float64

func newSensorData(b []byte, offset int) sensorData {

	if offset+3 > len(b) {
		panic("Offset and length must fit within size of slice")
	}

	var intVal uint32
	intVal = uint32(b[offset])<<12 | (0xFF&uint32(b[offset+1]))<<4 | (0xF0&uint32(b[offset+2]))>>4

	return sensorData(intVal)
}
