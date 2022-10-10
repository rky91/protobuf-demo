package sample

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "protobuf-demo/pb/proto"
	"protobuf-demo/utils"
)

func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  utils.RandomLayout(),
		Backlit: utils.RandomBool(),
	}
	return keyboard
}
func NewCpu() *pb.CPU {
	brand := utils.RandomCPUBrand()
	noOfCores := uint32(utils.RandomInt(2, 16))
	noOfThreads := uint32(utils.RandomInt(int(noOfCores), 32))
	minGhz := utils.RandomFloat64(2.0, 3.5)
	maxGhz := utils.RandomFloat64(minGhz, 5.0)

	cpu := &pb.CPU{
		Brand:         brand,
		Name:          utils.RandomCPUName(brand),
		NumberCores:   noOfCores,
		NumberThreads: noOfThreads,
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}

	return cpu
}

func NewGpu() *pb.GPU {
	brand := utils.RandomGPUBrand()
	name := utils.RandomGPUName(brand)

	minGhz := utils.RandomFloat64(1.0, 3.0)
	maxGhz := utils.RandomFloat64(minGhz, 6.0)

	memory := &pb.Memory{
		Value: uint64(utils.RandomInt(4, 16)),
		Unit:  pb.Memory_GB,
	}

	gpu := &pb.GPU{
		Name:   name,
		Brand:  brand,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}
	return gpu
}

func NewRam() *pb.Memory {
	return &pb.Memory{
		Value: uint64(utils.RandomInt(2, 16)),
		Unit:  pb.Memory_GB,
	}
}

func NewSSD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(utils.RandomInt(128, 2000)),
			Unit:  pb.Memory_GB,
		},
	}

	return ssd
}

func NewHDD() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(utils.RandomInt(128, 2000)),
			Unit:  pb.Memory_GB,
		},
	}

	return hdd
}

func NewScreen() *pb.Screen {
	panel := utils.RandomPanel()
	resolution := utils.RandomScreenResolution()

	return &pb.Screen{
		SizeInch:   utils.RandomFloat32(13, 32),
		Resolution: resolution,
		Panel:      panel,
		MultiTouch: utils.RandomBool(),
	}
}

func NewLaptop() *pb.Laptop {
	return &pb.Laptop{
		Id:       utils.RandomUuid(),
		Name:     utils.RandomStringFromList("L1", "L2"),
		Brand:    utils.RandomStringFromList("APPLE", "LENOVO", "HP"),
		Cpu:      NewCpu(),
		Memory:   NewRam(),
		GPUs:     []*pb.GPU{NewGpu()},
		Storages: []*pb.Storage{NewHDD(), NewSSD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: utils.RandomFloat64(1.0, 3.0),
		},
		PriceUsd:    utils.RandomFloat64(200, 300),
		ReleaseYear: uint32(utils.RandomInt(2012, 2022)),
		UpdatedAt:   timestamppb.Now(),
	}
}
