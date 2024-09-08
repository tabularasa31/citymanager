package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	pb "github.com/tabularasa31/citymanager/api/gen"
	"github.com/tabularasa31/citymanager/internal/storage"
)

func TestInMemoryStorage(t *testing.T) {
	store := storage.NewInMemoryStorage()
	ctx := context.Background()

	t.Run("AddCity", func(t *testing.T) {
		city := &pb.City{Name: "TestCity", Latitude: 1.0, Longitude: 2.0}
		err := store.AddCity(ctx, city)
		assert.NoError(t, err)
	})

	t.Run("GetCity", func(t *testing.T) {
		city, err := store.GetCity(ctx, "TestCity")
		assert.NoError(t, err)
		assert.NotNil(t, city)
		assert.Equal(t, "TestCity", city.Name)
		assert.Equal(t, 1.0, city.Latitude)
		assert.Equal(t, 2.0, city.Longitude)
	})

	t.Run("GetAllCities", func(t *testing.T) {
		cities, err := store.GetAllCities(ctx)
		assert.NoError(t, err)
		assert.Len(t, cities, 1)
	})

	t.Run("RemoveCity", func(t *testing.T) {
		err := store.RemoveCity(ctx, "TestCity")
		assert.NoError(t, err)

		_, err = store.GetCity(ctx, "TestCity")
		assert.Error(t, err)
	})
}
