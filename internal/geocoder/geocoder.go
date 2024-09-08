package geocoder

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Geocoder interface {
	Geocode(ctx context.Context, cityName string) (float64, float64, error)
}

type OpenStreetMapGeocoder struct {
	client  *http.Client
	baseURL string
}

func NewOpenStreetMapGeocoder() *OpenStreetMapGeocoder {
	return &OpenStreetMapGeocoder{
		client:  &http.Client{},
		baseURL: "https://nominatim.openstreetmap.org",
	}
}

func (g *OpenStreetMapGeocoder) SetHTTPClient(client *http.Client) {
	g.client = client
}

func (g *OpenStreetMapGeocoder) SetBaseURL(baseURL string) {
	g.baseURL = baseURL
}

func (g *OpenStreetMapGeocoder) Geocode(ctx context.Context, cityName string) (float64, float64, error) {
	requestURL := fmt.Sprintf("%s/search?q=%s+city&format=json&featuretype=city&limit=1", g.baseURL, url.QueryEscape(cityName))

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := g.client.Do(req)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to send request: %w", err)
	}
	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			if err != nil {
				err = fmt.Errorf("%w; failed to close response body: %w", err, closeErr)
			} else {
				err = fmt.Errorf("failed to close response body: %w", closeErr)
			}
		}
	}()

	var result []struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, 0, fmt.Errorf("failed to decode response: %w", err)
	}

	if len(result) == 0 {
		return 0, 0, fmt.Errorf("no results found for %s", cityName)
	}

	lat, err := strconv.ParseFloat(result[0].Lat, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse latitude: %w", err)
	}

	lon, err := strconv.ParseFloat(result[0].Lon, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse longitude: %w", err)
	}

	return lat, lon, nil
}
