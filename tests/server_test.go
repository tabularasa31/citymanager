package test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	pb "github.com/tabularasa31/citymanager/api/gen"
	"github.com/tabularasa31/citymanager/internal/server"
)

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) AddCity(ctx context.Context, city *pb.City) error {
	args := m.Called(ctx, city)
	return args.Error(0)
}

func (m *MockStorage) RemoveCity(ctx context.Context, name string) error {
	args := m.Called(ctx, name)
	return args.Error(0)
}

func (m *MockStorage) GetCity(ctx context.Context, name string) (*pb.City, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(*pb.City), args.Error(1)
}

func (m *MockStorage) GetAllCities(ctx context.Context) ([]*pb.City, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*pb.City), args.Error(1)
}

type MockGeocoder struct {
	mock.Mock
}

func (m *MockGeocoder) Geocode(ctx context.Context, cityName string) (float64, float64, error) {
	args := m.Called(ctx, cityName)
	return args.Get(0).(float64), args.Get(1).(float64), args.Error(2)
}

func TestAddCity(t *testing.T) {
	mockStorage := new(MockStorage)
	mockGeocoder := new(MockGeocoder)
	s := server.NewCityManagerServer(mockStorage, mockGeocoder)

	ctx := context.Background()
	req := &pb.AddCityRequest{Name: "TestCity"}

	// Используем mock.MatchedBy для проверки типа контекста
	ctxMatcher := mock.MatchedBy(func(ctx context.Context) bool {
		_, ok := ctx.Deadline()
		return ok // Проверяем, что у контекста установлен дедлайн
	})

	// Используем mock.MatchedBy для проверки содержимого City
	cityMatcher := mock.MatchedBy(func(city *pb.City) bool {
		return city.Name == "TestCity" && city.Latitude == 1.0 && city.Longitude == 2.0
	})

	mockGeocoder.On("Geocode", ctxMatcher, "TestCity").Return(1.0, 2.0, nil)
	mockStorage.On("AddCity", ctxMatcher, cityMatcher).Return(nil)

	resp, err := s.AddCity(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Success)
	mockGeocoder.AssertExpectations(t)
	mockStorage.AssertExpectations(t)
}

func TestRemoveCity(t *testing.T) {
	mockStorage := new(MockStorage)
	mockGeocoder := new(MockGeocoder)
	s := server.NewCityManagerServer(mockStorage, mockGeocoder)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.RemoveCityRequest{Name: "TestCity"}

	mockStorage.On("RemoveCity", mock.Anything, "TestCity").Return(nil)

	resp, err := s.RemoveCity(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Success)
	mockStorage.AssertExpectations(t)
}

func TestGetCity(t *testing.T) {
	mockStorage := new(MockStorage)
	mockGeocoder := new(MockGeocoder)
	s := server.NewCityManagerServer(mockStorage, mockGeocoder)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.GetCityRequest{Name: "TestCity"}

	mockCity := &pb.City{Name: "TestCity", Latitude: 1.0, Longitude: 2.0}
	mockStorage.On("GetCity", mock.Anything, "TestCity").Return(mockCity, nil)

	resp, err := s.GetCity(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, mockCity, resp.City)
	mockStorage.AssertExpectations(t)
}

func TestGetNearestCities(t *testing.T) {
	mockStorage := new(MockStorage)
	mockGeocoder := new(MockGeocoder)
	s := server.NewCityManagerServer(mockStorage, mockGeocoder)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.GetNearestCitiesRequest{Latitude: 0.0, Longitude: 0.0}

	mockCities := []*pb.City{
		{Name: "City1", Latitude: 1.0, Longitude: 1.0},
		{Name: "City2", Latitude: 2.0, Longitude: 2.0},
		{Name: "City3", Latitude: 3.0, Longitude: 3.0},
	}
	mockStorage.On("GetAllCities", mock.Anything).Return(mockCities, nil)

	resp, err := s.GetNearestCities(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Cities, 2)
	assert.Equal(t, "City1", resp.Cities[0].Name)
	assert.Equal(t, "City2", resp.Cities[1].Name)
	mockStorage.AssertExpectations(t)
}
