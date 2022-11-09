package model

import (
	"encoding/json"
	"time"
)

type MetaData interface{}

type ResultZincSearch struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index     string    `json:"_index"`
			Type      string    `json:"_type"`
			ID        string    `json:"_id"`
			Score     float64   `json:"_score"`
			Timestamp time.Time `json:"@timestamp"`
			Source    struct {
				Bcc     string `json:"Bcc"`
				Cc      string `json:"Cc"`
				Content string `json:"Content"`
				DateMSG string `json:"DateMSG"`
				From    string `json:"From"`
				IDMSG   string `json:"IdMSG"`
				Path    string `json:"Path"`
				Subject string `json:"Subject"`
				To      string `json:"To"`
				XFrom   string `json:"X-From"`
				XTo     string `json:"X-To"`
				XBcc    string `json:"X-bcc"`
				XCc     string `json:"X-cc"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func (s *ResultZincSearch.Hits) ToJson() ([]byte, error) {
	return json.Marshal(s)
}

func (s *ResultZincSearch.Hits.Hits ) GetFrom() string {
	return s.From
}