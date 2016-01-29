package gountries

// BaseLang is a basic structure for common language formatting in the JSON file
type BaseLang struct {
	Common   string `json:"common"`
	Official string `json:"official"`
}

// Geo contains geographical information
type Geo struct {
	Region    string `json:"region"`
	SubRegion string `json:"subregion"`
	Continent string // Yaml
	Capital   string `json:"capital"`

	Area float64
}

// Coordinates contains the coordinates for both Country and SubDivision
type Coordinates struct {
	LongitudeString string `json:"longitude"`
	LatitudeString  string `json:"latitude"`

	MinLongitude float64
	MinLatitude  float64
	MaxLongitude float64
	MaxLatitude  float64
	Latitude     float64
	Longitude    float64
}

// Measurer provides coordinates for measurements
type Measurer interface {
	Coordinates() (minLong, minLat, maxLong, maxLat float64)
}