# CityManager

CityManager is a gRPC-based service for managing city data and performing geographical calculations.

## Features

- Add and remove cities from storage
- Retrieve city information
- Find nearest cities based on coordinates
- Calculate distances between cities using the haversine formula

## GRPC API Methods

| Method | Description |
|--------|-------------|
| `AddCity` | Adds a new city to the storage |
| `RemoveCity` | Removes a city from the storage |
| `GetCity` | Retrieves information about a specific city |
| `GetNearestCities` | Finds the two nearest cities to given coordinates |

## Distance Calculation

The service uses the haversine formula to calculate the distance between two points on a sphere (Earth). This provides a good approximation for the distance between cities.

### Haversine Formula

```
d = 2R * arcsin(√(hav(φ2 - φ1) + cos(φ1)cos(φ2)hav(λ2 - λ1)))
```

Where:
- d: distance between two points on the surface of the sphere
- R: radius of the sphere (for Earth, approximately 6371 km)
- φ1, φ2: latitudes of the two points in radians
- λ1, λ2: longitudes of the two points in radians
- hav: haversine function, where hav(θ) = sin²(θ/2)

## Getting Started

### Prerequisites

- Go (version 1.23 or higher)
- Make

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/citymanager.git
   ```
2. Navigate to the project directory:
   ```
   cd citymanager
   ```

### Running the Service

To start the service, run the following command in the project directory:

```
make up
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the [MIT License](LICENSE).
