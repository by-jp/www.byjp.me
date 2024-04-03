package main

import "time"

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
