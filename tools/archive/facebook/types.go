package main

import (
	"time"

	"github.com/by-jp/www.byjp.me/tools/shared"
)

type PostCheckinPhotoOrVideo struct {
	Title       string    `json:"title"`
	Timestamp   int       `json:"timestamp"`
	Data        DataSlice `json:"data"`
	Attachments []struct {
		Data DataSlice `json:"data"`
	} `json:"attachments"`
	Tags []Tag `json:"tags"`
}

type Tag struct {
	Name string `json:"name"`
}

type DataSlice []map[string]interface{}

func (ds DataSlice) GetString(key string) string {
	for _, obj := range ds {
		iface, ok := obj[key]
		if !ok {
			continue
		}

		str, ok := iface.(string)
		if !ok {
			continue
		}

		return str
	}

	return ""
}

func (ds DataSlice) GetTime(key string) time.Time {
	for _, obj := range ds {
		iface, ok := obj[key]
		if !ok {
			continue
		}

		i, ok := iface.(int)
		if !ok {
			continue
		}

		return time.Unix(int64(i), 0).UTC()
	}

	return time.Unix(0, 0)
}

func (ds DataSlice) GetSubstrings(key string, sub string) []string {
	var substrs []string
	for _, obj := range ds {
		iface, ok := obj[key]
		if !ok {
			continue
		}

		m, ok := iface.(map[string]interface{})
		if !ok {
			continue
		}

		si, ok := m[sub]
		if !ok {
			continue
		}

		s, ok := si.(string)
		if !ok {
			continue
		}

		substrs = append(substrs, s)
	}

	return substrs
}

func (ds DataSlice) GetLocation() (shared.Location, bool) {
	for _, obj := range ds {
		iface, ok := obj["place"]
		if !ok {
			continue
		}

		p, ok := iface.(map[string]interface{})
		if !ok {
			continue
		}

		nameI, nameOK := p["name"]
		if !nameOK {
			continue
		}

		name, ok := nameI.(string)
		if !ok {
			continue
		}

		coordsI, coordsOK := p["coordinate"]
		if !coordsOK {
			continue
		}

		coordsM, ok := coordsI.(map[string]interface{})
		if !ok {
			continue
		}

		lat, ok := coordsM["latitude"]
		if !ok {
			continue
		}

		lng, ok := coordsM["longitude"]
		if !ok {
			continue
		}

		return shared.Location{
			Name:      name,
			Latitude:  lat.(float64),
			Longitude: lng.(float64),
		}, true
	}

	return shared.Location{}, false
}
