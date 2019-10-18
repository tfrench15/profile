package profile

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Event represents an event associated with a user's Profile.
type Event map[string]interface{}

// Events contain a slice of Event.
type Events struct {
	Data   interface{} `json:"data"`
	Cursor *Cursor     `json:"cursor"`
}

// EventsRequestConfig allows the client to pass in additional query parameters to customize
// its requst.
type EventsRequestConfig struct {
	Include []string
	Exclude []string
	Start   time.Time
	End     time.Time
	Sort    string
	Limit   int
}

// GetEvents queries the Profile API for the given ID's events.
func (c *Client) GetEvents(id, value string, config *EventsRequestConfig) (*Events, error) {
	url := baseURL + c.namespaceID + usersCollection + id + ":" + value + "/events"
	if config != nil {
		url = url + config.Encode()
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.secret, "")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	events := newEvents()
	dec := json.NewDecoder(res.Body)

	err = dec.Decode(events)
	if err != nil {
		return nil, err
	}

	fmt.Println(events)
	return events, nil
}

// NewEventsRequestConfig returns a new configuration struct for further customizing Events requests.
func NewEventsRequestConfig(include, exclude []string, sort string, limit int, start, end time.Time) (*EventsRequestConfig, error) {
	if sort != "asc" || sort != "desc" {
		return nil, fmt.Errorf("invalid sort param: %s, must be 'asc' or 'desc'", sort)
	}

	config := &EventsRequestConfig{
		Include: include,
		Exclude: exclude,
		Sort:    sort,
		Limit:   limit,
		Start:   start,
		End:     end,
	}

	return config, nil
}

// Encode returns the URL encoding of the query parameters contained in the EventsRequestConfig.
func (config *EventsRequestConfig) Encode() string {
	params := url.Values{}

	params.Set("include", strings.Join(config.Include, ","))
	params.Set("exclude", strings.Join(config.Exclude, ","))
	params.Set("sort", config.Sort)
	params.Set("limit", string(config.Limit))

	// TODO: set Start and End params.

	return params.Encode()
}

func newEvents() *Events {
	m := make(Event)
	c := new(Cursor)

	return &Events{
		Data:   m,
		Cursor: c,
	}
}
