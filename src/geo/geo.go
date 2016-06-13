package geo

import (
	"github.com/abh/geoip"
)

type Loc struct {
	Iso2Code    string  `json:"iso2_code"`
	Iso3Code    string  `json:"iso3_code"`
	CountryName string  `json:"country_name"`
	Region      string  `json:"region"`
	City        string  `json:"city"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
}

type Geo struct {
	gi *geoip.GeoIP
}

func New() (*Geo, error) {
	var err error
	g := &Geo{}
	g.gi, err = geoip.Open() // Will use system installed database(s)
	return g, err
}

func (g *Geo) GetLoc(ipAddress string) *Loc {
	if record := g.gi.GetRecord(ipAddress); record != nil {
		return &Loc{
			Iso2Code:    record.CountryCode,
			Iso3Code:    record.CountryCode3,
			CountryName: record.CountryName,
			Region:      record.Region,
			City:        record.City,
			Latitude:    record.Latitude,
			Longitude:   record.Longitude,
		}
	}
	return nil
}
