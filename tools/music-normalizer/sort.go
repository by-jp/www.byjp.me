package main

import (
	"strings"
)

type mbResponse struct {
	Recordings []mbRecording
}

type mbRecording struct {
	ID             string
	Score          int
	Title          string
	Disambiguation string
	FirstRelease   string      `json:"first-release-date"`
	ArtistCredit   []mbArtist  `json:"artist-credit"`
	Releases       []mbRelease `json:"releases"`
}

type mbArtist struct {
	Name string `json:"name"`
}

type mbRelease struct {
	ID           string
	Score        int
	Title        string
	Country      string
	Status       string
	Quality      string
	ArtistCredit []mbArtist `json:"artist-credit"`
	Date         string     `json:"first-release-date"`
}

type ByDesirability []mbRecording

func (r ByDesirability) Len() int      { return len(r) }
func (r ByDesirability) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func (r ByDesirability) Less(i, j int) bool {
	return sortInTurn(r[i], r[j],
		// highScoreFirst,
		preferGoodRelease,
		preferNonLiveRecordings,
		earliestRelease,
		preferShorterName,
		deterministicallyByID,
	)
}

func sortInTurn[T interface{}](a, b T, comparators ...func(T, T) int) bool {
	for _, comparator := range comparators {
		switch comparator(a, b) {
		case -1:
			return true
		case 1:
			return false
		default:
			// Try the next comparator
		}
	}
	// Don't swap if they're identical
	return false
}

func ternary(first, last bool) int {
	if first {
		return -1
	} else if last {
		return 1
	} else {
		return 0
	}
}

func highScoreFirst(a, b mbRecording) int {
	as := a.Score / 20
	bs := b.Score / 20
	return ternary(as > bs, as < bs)
}

func preferNonLiveRecordings(a, b mbRecording) int {
	al := stringIs(a.Disambiguation, "live")
	bl := stringIs(b.Disambiguation, "live")
	return ternary(!al && bl, al && !bl)
}

func preferGoodRelease(a, b mbRecording) int {
	filter := func(r mbRelease) bool {
		if len(r.ArtistCredit) > 0 && stringIs(r.ArtistCredit[0].Name, "various") {
			return false
		}
		if !stringIs(r.Status, "official") {
			return false
		}
		return true
	}

	ar := filterReleases(a, filter)
	br := filterReleases(a, filter)
	return ternary(!ar && br, ar && !br)
}

func preferShorterName(a, b mbRecording) int {
	return ternary(len(a.Title) < len(b.Title), len(a.Title) > len(b.Title))
}

func filterReleases(r mbRecording, filter func(mbRelease) bool) bool {
	for _, rel := range r.Releases {
		if filter(rel) {
			return true
		}
	}
	return false
}

func earliestRelease(a, b mbRecording) int {
	ad := parseReleaseDate(a.FirstRelease)
	bd := parseReleaseDate(b.FirstRelease)
	return ternary(ad.Before(bd), ad.After(bd))
}

func deterministicallyByID(a, b mbRecording) int {
	return strings.Compare(a.ID, b.ID)
}
