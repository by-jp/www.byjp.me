package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

func searchMusicbrainz(trackName, artistName string) (musicbrainzData, error) {
	var mb musicbrainzData
	query := url.QueryEscape(fmt.Sprintf("title:\"%s\" AND artist:\"%s\"", trackName, artistName))
	searchURL := fmt.Sprintf("https://musicbrainz.org/ws/2/recording/?query=%s&method=advanced&limit=50&fmt=json", query)

	resp, err := http.Get(searchURL)
	if err != nil {
		return mb, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return mb, err
	}

	var res mbResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return mb, err
	}
	if len(res.Recordings) == 0 {
		return mb, fmt.Errorf("no results found on MusicBrainz")
	}

	recs := filter(res.Recordings,
		hasCorrectDetails(trackName, artistName),
	)

	sort.Sort(ByDesirability(recs))

	mb.ID = recs[0].ID
	mb.URL = "https://musicbrainz.org/recording/" + recs[0].ID
	mb.Title = recs[0].Title
	mb.Artist = recs[0].ArtistCredit[0].Name

	return mb, nil
}

func stringIs(str string, has ...string) bool {
	t := strings.ToLower(str)
	for _, ha := range has {
		if strings.Contains(t, ha) {
			return true
		}
	}
	return false
}

func parseReleaseDate(str string) time.Time {
	if t, err := time.Parse("2006-01-02", str); err == nil {
		return t
	}

	if t, err := time.Parse("2006", str); err == nil {
		// Go to the last day of the year, as we want more specific dates to take precedence
		return t.AddDate(1, 0, -1)
	}

	return time.Time{}
}
