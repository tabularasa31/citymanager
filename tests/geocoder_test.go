package test

import (
	"context"
	"github.com/tabularasa31/citymanager/internal/geocoder"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenStreetMapGeocoder(t *testing.T) {
	// Создаем тестовый HTTP сервер с более реалистичными данными
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"lat":"51.4893335","lon":"-0.14405508452768728"}]`))
	}))
	defer server.Close()

	// Создаем геокодер и устанавливаем тестовый клиент и URL
	geocoder := geocoder.NewOpenStreetMapGeocoder()
	geocoder.SetHTTPClient(server.Client())
	geocoder.SetBaseURL(server.URL)

	ctx := context.Background()
	lat, lon, err := geocoder.Geocode(ctx, "London")

	assert.NoError(t, err)
	// Используем более мягкую проверку с допустимым отклонением
	assert.InDelta(t, 51.4893335, lat, 0.01, "Latitude should be close to expected value")
	assert.InDelta(t, -0.14405508452768728, lon, 0.01, "Longitude should be close to expected value")
}

func TestOpenStreetMapGeocoderError(t *testing.T) {
	// Создаем тестовый HTTP сервер, который возвращает ошибку
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	// Создаем геокодер и устанавливаем тестовый клиент и URL
	geocoder := geocoder.NewOpenStreetMapGeocoder()
	geocoder.SetHTTPClient(server.Client())
	geocoder.SetBaseURL(server.URL)

	ctx := context.Background()
	_, _, err := geocoder.Geocode(ctx, "InvalidCity")

	assert.Error(t, err)
}
