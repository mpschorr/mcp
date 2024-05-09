package datatype

import (
	"bufio"
)

func ReadString(buf *bufio.Reader) string {
	length, _ := ReadVarInt(buf)
	str := make([]byte, length)
	for i := 0; i < int(length); i++ {
		str[i], _ = buf.ReadByte()
	}

	return string(str)
}

func WriteString(buf *bufio.Writer, str string) {
	WriteVarInt(buf, int32(len(str)))
	buf.WriteString(str)
}