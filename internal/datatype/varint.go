package datatype

import (
	"bufio"
)

func ReadVarInt(buf *bufio.Reader) (int32, error) {
	var val int32
	var pos int32
	for {
		cur, err := buf.ReadByte()
		val |= (int32(cur) & 0x7F) << pos
		if err != nil {
			return 0, err
		}
		if cur&0x80 == 0 {
			break;
		}
		pos += 7
	}

	return val, nil
}

func WriteVarInt(buf *bufio.Writer, val int32) error {
	for {
		if val < 0x80 {
			err := buf.WriteByte(byte(val))
			if err != nil {
				return err
			}
			break
		}
		err := buf.WriteByte(byte(val | 0x80))
		val = int32(uint32(val) >> 7)
		if err != nil {
			return err
		}
	}
	return nil
}
	// for {
	// 	if val < 0x80 {
	// 		err := buf.WriteByte(byte(val))
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	// 	err := buf.WriteByte(byte(val | 0x80))
	// 	val = int32(uint32(val) >> 7)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
