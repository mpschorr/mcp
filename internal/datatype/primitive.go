package datatype

import "bufio"

// Boolean

func ReadBoolean(buf *bufio.Reader) bool {
	val, _ := buf.ReadByte()
	return val == 1
}

func WriteBoolean(buf *bufio.Writer, val bool) {
	if val {
		buf.WriteByte(1)
	} else {
		buf.WriteByte(0)
	}
}


// Byte

// TODO ummm idk this might do some weird thing with the sign bit
func ReadByte(buf *bufio.Reader) int8 {
	val, _ := buf.ReadByte()
	return int8(val)
}

func WriteByte(buf *bufio.Writer, val int8) {
	buf.WriteByte(byte(val))
}

func ReadUnsignedByte(buf *bufio.Reader) byte {
	val, _ := buf.ReadByte()
	return val
}

func WriteUnsignedByte(buf *bufio.Writer, val byte) {
	buf.WriteByte(val)
}

// Short

func ReadShort(buf *bufio.Reader) int16 {
	var val int16
	for i := 0; i < 2; i++ {
		read, _ := buf.ReadByte()
		val = val << 8 | int16(read)
	}
	return val
}

// TODO maybe broken? ai gen
func WriteShort(buf *bufio.Writer, val int16) {
	for i := 0; i < 2; i++ {
		buf.WriteByte(byte(val >> (i * 8)))
	}
}

func ReadUnsignedShort(buf *bufio.Reader) uint16 {
	var val uint16
	for i := 0; i < 2; i++ {
		read, _ := buf.ReadByte()
		val = val << 8 | uint16(read)
	}
	return val
}

// TODO maybe broken? ai gen
func WriteUnsignedShort(buf *bufio.Writer, val uint16) {
	for i := 0; i < 2; i++ {
		buf.WriteByte(byte(val >> (i * 8)))
	}
}

// Int

func ReadInt(buf *bufio.Reader) int32 {
	var val int32
	for i := 0; i < 4; i++ {
		read, _ := buf.ReadByte()
		val = val << 8 | int32(read)
	}
	return val
}

func WriteInt(buf *bufio.Writer, val int32) {
	for i := 0; i < 4; i++ {
		buf.WriteByte(byte(val >> (i * 8)))
	}
}

// Long

func ReadLong(buf *bufio.Reader) int64 {
	var val int64
	for i := 0; i < 8; i++ {
		read, _ := buf.ReadByte()
		val = val << 8 | int64(read)
	}
	return val
}

func WriteLong(buf *bufio.Writer, val int64) {
	for i := 0; i < 8; i++ {
		buf.WriteByte(byte(val >> (i * 8)))
	}
}