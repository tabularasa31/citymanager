package storage

import (
	"errors"
	"sync"

	pb "github.com/tabularasa31/citymanager/api/gen"
	"golang.org/x/net/context"
)

type Storage interface {
	AddCity(ctx context.Context, city *pb.City) error
	RemoveCity(ctx context.Context, name string) error
	GetCity(ctx context.Context, name string) (*pb.City, error)
	GetAllCities(ctx context.Context) ([]*pb.City, error)
}

type InMemoryStorage struct {
	cities map[string]*pb.City
	mutex  sync.RWMutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		cities: make(map[string]*pb.City),
	}
}

func (s *InMemoryStorage) AddCity(ctx context.Context, city *pb.City) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.cities[city.Name] = city
		return nil
	}
}

func (s *InMemoryStorage) RemoveCity(ctx context.Context, name string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		s.mutex.Lock()
		defer s.mutex.Unlock()
		if _, ok := s.cities[name]; !ok {
			return errors.New("city not found")
		}
		delete(s.cities, name)
		return nil
	}
}

func (s *InMemoryStorage) GetCity(ctx context.Context, name string) (*pb.City, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		s.mutex.RLock()
		defer s.mutex.RUnlock()
		city, ok := s.cities[name]
		if !ok {
			return nil, errors.New("city not found")
		}
		return city, nil
	}
}

func (s *InMemoryStorage) GetAllCities(ctx context.Context) ([]*pb.City, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		s.mutex.RLock()
		defer s.mutex.RUnlock()
		cities := make([]*pb.City, 0, len(s.cities))
		for _, city := range s.cities {
			cities = append(cities, city)
		}
		return cities, nil
	}
}
