package profile

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Events contain a slice of Event.
type Events struct {
	Data   interface{} `json:"data"`
	Cursor *Cursor     `json:"cursor"`
}

// EventRequest allows the client to pass in additional query parameters to customize
// its requst.
type EventRequest struct {
	// mandatory params
	id    string
	value string

	// optional params
	queryParams url.Values
}

// GetEvents queries the Profile API for the given ID's events.
func (c *Client) GetEvents(request *EventRequest) (*Events, error) {
	url := baseURL + c.namespaceID + usersCollection + request.id + ":" + request.value + "/events"
	if len(request.queryParams) > 0 {
		url = url + request.queryParams.Encode()
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

	return events, nil
}

// NewEventRequest returns a default EventRequest.
func NewEventRequest(id, value string) *EventRequest {
	return &EventRequest{
		id:          id,
		value:       value,
		queryParams: url.Values{},
	}
}

// SetInclude allows a client to specify certain events it would like to include in a request to the Profile API.
func (req *EventRequest) SetInclude(events ...string) error {
	if len(events) == 0 {
		return errors.New("0 events passed in, please specify events to include")
	}

	var include []string
	for _, event := range events {
		include = append(include, event)
	}

	req.queryParams.Set("inlcude", strings.Join(include, ","))
	return nil
}

// SetExclude allows a client to specify certain events it would like to exclude in a request to the Profile API.
func (req *EventRequest) SetExclude(events ...string) error {
	if len(events) == 0 {
		return errors.New("0 events passed in, please specify events to exclude")
	}

	var exclude []string
	for _, event := range events {
		exclude = append(exclude, event)
	}

	req.queryParams.Set("exclude", strings.Join(exclude, ","))
	return nil
}

// SetSort allows a client to specify whether the data returned from the Profile API should be sorted.
func (req *EventRequest) SetSort(sort string) error {
	if sort != "asc" || sort != "desc" {
		return fmt.Errorf("sort must be 'asc' or 'desc, got %s", sort)
	}

	req.queryParams.Set("sort", sort)

	return nil
}

// SetLimit sets the limit for number of Events that will be returned by the request to the Profile API.
func (req *EventRequest) SetLimit(limit int) error {
	if limit < 1 || limit > 100 {
		return fmt.Errorf("limit must be at least 1 and at most 100, got %d", limit)
	}

	req.queryParams.Set("limit", string(limit))
	return nil
}

// Validate ensures the request is valid, and satisfies the Request interface.
func (req *EventRequest) Validate() error {
	if len(req.id) == 0 {
		return errors.New("request must specify an ID to query by")
	}

	if len(req.value) == 0 {
		return errors.New("request must specify an ID value to query by")
	}

	return nil
}

func newEvents() *Events {
	m := make(map[string]interface{})
	c := new(Cursor)

	return &Events{
		Data:   m,
		Cursor: c,
	}
}

func (req *EventRequest) internal() {}
