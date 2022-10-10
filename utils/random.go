package utils

import (
	"github.com/google/uuid"
	"math/rand"
	pb "protobuf-demo/pb/proto"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_UNKNOWN
	}
}

func RandomBool() bool {
	return rand.Intn(2) == 1
}

func RandomCPUBrand() string {
	brands := []string{"INTEL", "APPLE", "AMD"}
	return brands[rand.Intn(len(brands))]
}

func RandomGPUBrand() string {
	brands := []string{"NVIDIA", "AMD"}
	return brands[rand.Intn(len(brands))]
}

func RandomCPUName(brand string) string {
	if brand == "INTEL" {
		return RandomStringFromList("i9", "i7", "i3", "evo")
	}

	return RandomStringFromList("Rezen 7", "Rezen 3")
}

func RandomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return RandomStringFromList("N1", "N2", "N3", "N4")
	}

	return RandomStringFromList("A 7", "A 3")
}

func RandomStringFromList(a ...string) string {
	l := len(a)
	if l == 0 {
		return ""
	} else {
		return a[rand.Intn(l)]
	}
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func RandomPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}

	return pb.Screen_LED
}

func RandomScreenResolution() *pb.Screen_Resolution {
	height := RandomInt(1080, 4320)
	width := height * 16 / 9

	return &pb.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}
}

func RandomUuid() string {
	return (uuid.New()).String()
}
