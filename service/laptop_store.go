package service

import (
	"errors"
	"github.com/jinzhu/copier"
	pb "protobuf-demo/pb/proto"
	"sync"
)

// LaptopStore is an interface to store laptop
type LaptopStore interface {
	// Save saves the laptop to the store
	Save(laptop *pb.Laptop) error
	// Find finds a laptop by ID
	Find(id string) (*pb.Laptop, error)
}

// InMemoryLaptopStore stores laptop in memory
type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

// NewInMemoryLaptopStore returns a new InMemoryLaptopStore
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

// ErrAlreadyExists is returned when a record with the same ID already exists in the store
var ErrAlreadyExists = errors.New("record already exists")

// Save saves the laptop to the store
func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	other := &pb.Laptop{}
	err := copier.CopyWithOption(other, laptop, copier.Option{
		DeepCopy: true,
	})
	if err != nil {
		return err
	}

	store.data[other.Id] = other
	return nil
}

// Find finds a laptop by ID
func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	other := &pb.Laptop{}
	err := copier.CopyWithOption(other, laptop, copier.Option{
		DeepCopy: true,
	})
	return other, err
}
