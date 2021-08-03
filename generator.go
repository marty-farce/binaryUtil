package binaryUtil

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
)

func Encode(packet int) *bytes.Buffer {
	binPacket := new(bytes.Buffer)
	switch packet {
	case 1:
		if err := binary.Write(binPacket, binary.BigEndian, populatePayload().P1); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	case 2:
		if err := binary.Write(binPacket, binary.BigEndian, populatePayload().P2); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	case 3:
		if err := binary.Write(binPacket, binary.BigEndian, populatePayload().P3); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	case 4:
		if err := binary.Write(binPacket, binary.BigEndian, populatePayload().P4); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	case 5:
		if err := binary.Write(binPacket, binary.BigEndian, populatePayload().P5); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	case 6:
		if err := binary.Write(binPacket, binary.BigEndian, populatePayload().P6); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	case 7:
		if err := binary.Write(binPacket, binary.BigEndian, populatePayload().P7); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	case 8:
		if err := binary.Write(binPacket, binary.BigEndian, populatePayload().P8); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	default:
		return nil
	}

	return binPacket
}

func Decode(data *bytes.Buffer) payload {
	var dataOut payload
	if err := binary.Read(data, binary.BigEndian, &dataOut); err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	return dataOut
}

func randomInt8() uint8 {
	return uint8(rand.Intn(225))
}

func randomInt16() uint16 {
	return uint16(rand.Intn(65535))
}
func randomInt32() int32 {
	return int32(rand.Intn(90999999))
}
