package binaryUtil

type packet1 struct {
	BatteryLife     uint8
	FirmwareVersion uint8
	Latitude        int32
	Longitude       int32
	Reserved        uint8
}

type packet2 struct {
	Lux      [2]uint16
	Temp     uint16
	Humidity uint16
}

type packet3 struct {
	SoilTemp6          uint16
	SoilCarbonDioxide6 uint16
	SoilEC6            uint16
	SoilMoisture6      uint16
}

type packet4 struct {
	SoilPH6 uint16
	SoilN6  uint16
	SoilP6  uint16
	SoilK6  uint16
}

type packet5 struct {
	SoilTemp18          uint16
	SoilCarbonDioxide18 uint16
	SoilEC18            uint16
	SoilMoisture18      uint16
}

type packet6 struct {
	SoilPH18 uint16
	SoilN18  uint16
	SoilP18  uint16
	SoilK18  uint16
}

type packet7 struct {
	SoilTemp36          uint16
	SoilCarbonDioxide36 uint16
	SoilEC36            uint16
	SoilMoisture36      uint16
}

type packet8 struct {
	SoilPH36 uint16
	SoilN36  uint16
	SoilP36  uint16
	SoilK36  uint16
}

type payload struct {
	P1 packet1
	P2 packet2
	P3 packet3
	P4 packet4
	P5 packet5
	P6 packet6
	P7 packet7
	P8 packet8
}

func populatePayload() payload {

	p1 := packet1{
		BatteryLife:     randomInt8(),
		FirmwareVersion: 1,
		Latitude:        randomInt32(),
		Longitude:       randomInt32(),
		Reserved:        randomInt8(),
	}

	p2 := packet2{
		Lux:      [2]uint16{randomInt16(), randomInt16()},
		Temp:     randomInt16(),
		Humidity: randomInt16(),
	}

	p3 := packet3{
		SoilTemp6:          randomInt16(),
		SoilCarbonDioxide6: randomInt16(),
		SoilEC6:            randomInt16(),
		SoilMoisture6:      randomInt16(),
	}

	p4 := packet4{
		SoilPH6: randomInt16(),
		SoilN6:  randomInt16(),
		SoilP6:  randomInt16(),
		SoilK6:  randomInt16(),
	}

	p5 := packet5{
		SoilTemp18:          randomInt16(),
		SoilCarbonDioxide18: randomInt16(),
		SoilEC18:            randomInt16(),
		SoilMoisture18:      randomInt16(),
	}

	p6 := packet6{
		SoilPH18: randomInt16(),
		SoilN18:  randomInt16(),
		SoilP18:  randomInt16(),
		SoilK18:  randomInt16(),
	}

	p7 := packet7{
		SoilTemp36:          randomInt16(),
		SoilCarbonDioxide36: randomInt16(),
		SoilEC36:            randomInt16(),
		SoilMoisture36:      randomInt16(),
	}

	p8 := packet8{
		SoilPH36: randomInt16(),
		SoilN36:  randomInt16(),
		SoilP36:  randomInt16(),
		SoilK36:  randomInt16(),
	}

	populated := payload{
		P1: p1,
		P2: p2,
		P3: p3,
		P4: p4,
		P5: p5,
		P6: p6,
		P7: p7,
		P8: p8,
	}

	return populated
}
