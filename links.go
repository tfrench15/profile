package profile

import (
	"encoding/json"
	"net/http"
)

// Links represents a slice of Link.
type Links []*Link

// Link represents the accounts linked to a user.
type Link struct {
	ToCollection string        `json:"to_collection"`
	ExternalIDs  []*ExternalID `json:"external_ids"`
}

// GetLinks queries the Profile API for the provided ID's links.
func (c *Client) GetLinks(id, value string) (*Links, error) {
	url := baseURL + c.namespaceID + usersCollection + id + ":" + value + "/links"

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

	links := new(Links)
	dec := json.NewDecoder(res.Body)

	err = dec.Decode(links)
	if err != nil {
		return nil, err
	}

	return links, nil
}
