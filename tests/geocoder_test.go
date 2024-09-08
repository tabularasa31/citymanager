package test

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tabularasa31/citymanager/internal/geocoder"
)

func TestOpenStreetMapGeocoder(t *testing.T) {
	// Создаем тестовый HTTP сервер с более реалистичными данными
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`[{"lat":"51.4893335","lon":"-0.14405508452768728"}]`))
		if err != nil {
			log.Printf("Error writing response: %v", err)
		}
	}))
	defer server.Close()

	// Создаем геокодер и устанавливаем тестовый клиент и URL
	geoClient := geocoder.NewOpenStreetMapGeocoder()
	geoClient.SetHTTPClient(server.Client())
	geoClient.SetBaseURL(server.URL)

	ctx := context.Background()
	lat, lon, err := geoClient.Geocode(ctx, "London")

	assert.NoError(t, err)
	// Используем более мягкую проверку с допустимым отклонением
	assert.InDelta(t, 51.4893335, lat, 0.01, "Latitude should be close to expected value")
	assert.InDelta(t, -0.14405508452768728, lon, 0.01, "Longitude should be close to expected value")
}

func TestOpenStreetMapGeocoderError(t *testing.T) {
	// Создаем тестовый HTTP сервер, который возвращает ошибку
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
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
