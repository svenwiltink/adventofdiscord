package adventofdiscord

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

const (
	urlTemplate = `https://adventofcode.com/%d/leaderboard/private/view/%s.json`
)

type StatsCollector struct {
	client        http.Client
	leaderboardID string
	lastFetch     time.Time
	last          Stats
	TTL           time.Duration
}

func (s StatsCollector) FetchStats() (Stats, error) {
	if time.Since(s.lastFetch) < s.TTL {
		return s.last, nil
	}

	// todo cache url?
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(urlTemplate, time.Now().Year(), s.leaderboardID), nil)
	if err != nil {
		return Stats{}, fmt.Errorf("unable to create http request: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return Stats{}, fmt.Errorf("unable to fetch stats: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Stats{}, fmt.Errorf("http status not OK, was %d", resp.StatusCode)
	}

	dec := json.NewDecoder(resp.Body)
	stats := Stats{}
	err = dec.Decode(&stats)
	if err != nil {
		return Stats{}, fmt.Errorf("error unmarshaling json: %w", err)
	}

	s.lastFetch = time.Now()
	s.last = stats

	return stats, nil
}

func NewStatsCollector(leaderboardID string, sessionID string) *StatsCollector {
	u, err := url.Parse("https://adventofcode.com")
	if err != nil {
		panic(err)
	}

	jar, _ := cookiejar.New(nil)

	jar.SetCookies(u, []*http.Cookie{{
		Name:   "session",
		Value:  sessionID,
		Path:   "/",
		Domain: "adventofcode.com",
	}})

	client := http.Client{
		Jar: jar,
	}

	return &StatsCollector{
		client:        client,
		leaderboardID: leaderboardID,
		TTL:           15 * time.Minute,
	}
}
