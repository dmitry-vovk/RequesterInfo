package geo

import (
	"testing"
	"log"
)

func TestGeo_GetLoc(t *testing.T) {
	g, err := New()
	if err != nil {
		t.Errorf("Error creating Geo: %s", err)
	}
	loc := g.GetLoc("127.0.0.1")
	log.Printf("Got location: %v", loc)
}
