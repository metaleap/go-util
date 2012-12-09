package geo

type Location struct {
	LonLat
	Altitude         float64
	AltitudeAbsolute bool
}

type LonLat struct {
	Longitude float64
	Latitude  float64
}
