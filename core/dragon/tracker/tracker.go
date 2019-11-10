package tracker

import (
	"encoding/json"
	"net/http"
	"time"
)

// resp header tracker key
const TrackKey = "dragonTrack"

type Tracker struct {
	SpanId    string      `json:"span_id"`
	Uri       string      `json:"uri"`
	Method    string      `json:"method"`
	ReqHeader http.Header `json:"req_header"`
	Body      string      `json:"body"`
	Resp      struct {
		Header http.Header `json:"header"`
		Data   string      `json:"data"`
	} `json:"resp"`
	Service struct {
		Req struct {
			Uri  string `json:"uri"`
			Body string `json:"body"`
		} `json:"req"`
		//Req *http.Request `json:"req"`
		Resp string `json:"resp"`
	} `json:"service"`
	StartTime time.Time   `json:"start_time"`
	CostTime  string      `json:"cost_time"`
	Error     interface{} `json:"error"`
}

func (tracker *Tracker) Marshal() string {
	m, _ := json.Marshal(tracker)
	return string(m)
}

func UnMarshal(s string) *Tracker {
	track := Tracker{}
	json.Unmarshal([]byte(s), &track)
	return &track
}
