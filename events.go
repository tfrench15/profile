package profile

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Event represents an event associated with a user's Profile.
type Event map[string]interface{}

// Events contain a slice of Event.
type Events struct {
	Data   Event   `json:"data"`
	Cursor *Cursor `json:"cursor"`
}

// GetEvents queries the Profile API for the given ID's events.
func (c *Client) GetEvents(id, value string) (*Events, error) {
	url := baseURL + c.namespaceID + usersCollection + id + ":" + value + "/events"

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

func newEvents() *Events {
	m := make(Event)
	c := new(Cursor)

	return &Events{
		Data:   m,
		Cursor: c,
	}
}
