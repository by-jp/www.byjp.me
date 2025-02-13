package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Spotify API credentials (set as environment variables)
var spotifyClientID = os.Getenv("SPOTIFY_CLIENT_ID")
var spotifyClientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")

// Structs to parse Spotify API responses
type SpotifyAuthResponse struct {
	AccessToken string `json:"access_token"`
}

type SpotifyTrackResponse struct {
	Name    string `json:"name"`
	Artists []struct {
		Name string `json:"name"`
	} `json:"artists"`
}

func getSpotifyAccessToken() (string, error) {
	authURL := "https://accounts.spotify.com/api/token"
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, _ := http.NewRequest("POST", authURL, strings.NewReader(data.Encode()))
	req.SetBasicAuth(spotifyClientID, spotifyClientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var authResp SpotifyAuthResponse
	json.Unmarshal(body, &authResp)

	return authResp.AccessToken, nil
}

func getSpotifyTrackMetadata(trackID string) (string, string, error) {
	token, err := getSpotifyAccessToken()
	if err != nil {
		return "", "", err
	}

	url := fmt.Sprintf("https://api.spotify.com/v1/tracks/%s", trackID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", "", err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var trackResp SpotifyTrackResponse
	json.Unmarshal(body, &trackResp)

	if len(trackResp.Artists) == 0 {
		return "", "", fmt.Errorf("no artist found for the track")
	}

	return trackResp.Name, trackResp.Artists[0].Name, nil
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type musicbrainzData struct {
	ID       string   `json:"-"`
	Composer string   `json:"composer,omitempty"`
	Artist   string   `json:"artist,omitempty"`
	Title    string   `json:"title"`
	URL      string   `json:"musicbrainz"`
	Links    []string `json:"links"`
}

func fromSpotifyTrack(trackID string) (musicbrainzData, error) {
	trackName, artistName, err := getSpotifyTrackMetadata(trackID)
	if err != nil {
		return musicbrainzData{}, err
	}

	fmt.Printf("Searching for %s - %s\n", trackName, artistName)
	mb, err := searchMusicbrainz(trackName, artistName)
	if err != nil {
		return mb, err
	}

	mb.Links = []string{"https://open.spotify.com/track/" + trackID}

	return mb, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <streaming-music-url>\n", os.Args[0])
		os.Exit(1)
	}

	u, err := url.Parse(os.Args[1])
	check(err)

	var mb musicbrainzData
	switch {
	case u.Host == "open.spotify.com" && strings.HasPrefix(u.Path, "/track/"):
		mb, err = fromSpotifyTrack(u.Path[7:])
	default:
		err = fmt.Errorf("unknown URL: %s", u)
	}
	check(err)

	f, err := os.OpenFile("./data/music/musicbrainz/"+mb.ID+".json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	check(err)

	check(json.NewEncoder(f).Encode(mb))
}
