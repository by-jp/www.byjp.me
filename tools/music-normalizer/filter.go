package main

import (
	"github.com/agnivade/levenshtein"
)

func filter[T interface{}](items []T, filters ...func(T) bool) []T {
	var allowed []T
	for _, item := range items {
		for _, filter := range filters {
			if filter(item) {
				allowed = append(allowed, item)
			}
		}
	}
	return allowed
}

func hasCorrectDetails(trackName, artistName string) func(mbRecording) bool {
	return func(mr mbRecording) bool {
		dt := levenshtein.ComputeDistance(trackName, mr.Title)
		da := levenshtein.ComputeDistance(artistName, mr.ArtistCredit[0].Name)

		return dt+da < 6
	}
}
