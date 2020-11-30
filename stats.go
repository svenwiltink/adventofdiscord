package adventofdiscord

import (
	"strconv"
	"strings"
	"time"
)

type CompletionTime time.Time

func (c CompletionTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(c).Unix(), 10)), nil
}

func (c *CompletionTime) UnmarshalJSON(bytes []byte) error {
	r := strings.Replace(string(bytes), `"`, ``, -1)

	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(c) = time.Unix(q, 0)
	return nil
}

type CompletionLevel struct {
	Time CompletionTime `json:"get_star_ts"`
}

type CompletionLevels struct {
	PartOne CompletionLevel `json:"1"`
	PartTwo CompletionLevel `json:"2"`
}

type Member struct {
	Name               string
	Stars              int
	GlobalScore        int                         `json:"global_score"`
	LocalScore         int                         `json:"local_score"`
	LastStarTime       CompletionTime              `json:"last_star_ts"`
	CompletionDayLevel map[string]CompletionLevels `json:"completion_day_level"`
}

type Stats struct {
	Event   string
	OwnerID string `json:"owner_id"`
	Members map[string]Member
}
