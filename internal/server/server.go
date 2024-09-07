package server

import (
	"context"
	"errors"
	"math"
	"sort"
	"time"

	pb "github.com/tabularasa31/citymanager/api/gen"
	"github.com/tabularasa31/citymanager/internal/geocoder"
	"github.com/tabularasa31/citymanager/internal/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CityManagerServer struct {
	pb.UnimplementedCityManagerServer
	store    storage.Storage
	geocoder geocoder.Geocoder
}

func NewCityManagerServer(store storage.Storage, geocoder geocoder.Geocoder) *CityManagerServer {
	return &CityManagerServer{store: store, geocoder: geocoder}
}

func (s *CityManagerServer) AddCity(ctx context.Context, req *pb.AddCityRequest) (*pb.AddCityResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "City name cannot be empty")
	}

	lat, lon, err := s.geocoder.Geocode(ctx, req.Name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to geocode city: %v", err)
	}

	city := &pb.City{
		Name:      req.Name,
		Latitude:  lat,
		Longitude: lon,
	}

	if err := s.store.AddCity(ctx, city); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to add city: %v", err)
	}

	return &pb.AddCityResponse{Success: true, Message: "City added successfully"}, nil
}

func (s *CityManagerServer) RemoveCity(ctx context.Context, req *pb.RemoveCityRequest) (*pb.RemoveCityResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "City name cannot be empty")
	}

	if err := s.store.RemoveCity(ctx, req.Name); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to remove city: %v", err)
	}

	return &pb.RemoveCityResponse{Success: true, Message: "City removed successfully"}, nil
}

func (s *CityManagerServer) GetCity(ctx context.Context, req *pb.GetCityRequest) (*pb.GetCityResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "City name cannot be empty")
	}

	city, err := s.store.GetCity(ctx, req.Name)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "City not found: %v", err)
	}

	return &pb.GetCityResponse{City: city}, nil
}

func (s *CityManagerServer) GetNearestCities(ctx context.Context, req *pb.GetNearestCitiesRequest) (*pb.GetNearestCitiesResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Валидация входных данных
	if err := validateCoordinates(req.Latitude, req.Longitude); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid coordinates: %v", err)
	}

	cities, err := s.store.GetAllCities(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Failed to get all cities: %v", err)
	}
	if len(cities) < 2 {
		return nil, status.Errorf(codes.FailedPrecondition, "Not enough cities in storage, at least 2 required")
	}

	// Сортировка городов по расстоянию
	sort.Slice(cities, func(i, j int) bool {
		dist1 := haversine(req.Latitude, req.Longitude, cities[i].Latitude, cities[i].Longitude)
		dist2 := haversine(req.Latitude, req.Longitude, cities[j].Latitude, cities[j].Longitude)
		return dist1 < dist2
	})

	// Возвращаем два ближайших города
	return &pb.GetNearestCitiesResponse{Cities: cities[:2]}, nil
}

// haversine вычисляет расстояние между двумя точками на сфере (Земле)
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Радиус Земли в километрах

	dLat := toRadians(lat2 - lat1)
	dLon := toRadians(lon2 - lon1)
	lat1 = toRadians(lat1)
	lat2 = toRadians(lat2)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)

	// Проверка на случай точек-антиподов (точки почти на противоположных сторонах Земли)
	if a > 1.0 {
		a = 1.0
	}

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

func toRadians(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func validateCoordinates(lat, lon float64) error {
	if lat < -90 || lat > 90 {
		return errors.New("latitude must be between -90 and 90")
	}
	if lon < -180 || lon > 180 {
		return errors.New("longitude must be between -180 and 180")
	}
	return nil
}
