package ingest

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"time"
)

type Weather struct {
	METAR string
	TAF   string
	Time  time.Time
}

type RSS struct {
	Channel struct {
		Items []struct {
			Title       string `xml:"title"`
			Description string `xml:"description"`
		} `xml:"item"`
	} `xml:"channel"`
}

func FetchMETAR(url string) (Weather, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Weather{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return Weather{}, errors.New("bad status")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Weather{}, err
	}
	// BoM JSON: {"metar":"","taf":""}
	type bom struct {
		METAR string `json:"metar"`
		TAF   string `json:"taf"`
	}
	var b bom
	if err := json.Unmarshal(body, &b); err != nil {
		return Weather{}, err
	}
	return Weather{b.METAR, b.TAF, time.Now()}, nil
}

func FetchNOTAM(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("bad status")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var rss RSS
	if err := xml.Unmarshal(body, &rss); err != nil {
		return nil, err
	}
	var out []string
	for _, item := range rss.Channel.Items {
		out = append(out, item.Title+" "+item.Description)
	}
	return out, nil
}
