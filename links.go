package profile

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Links represents a slice of Link.
type Links []*Link

// Link represents the accounts linked to a user.
type Link struct {
	ToCollection string        `json:"to_collection"`
	ExternalIDs  []*ExternalID `json:"external_ids"`
}

// LinkRequest comprises mandatory and required data for retrieving a user's Links from the Profile API.
type LinksRequest struct {
	id    string
	value string
}

// GetLinks queries the Profile API for the provided ID's links.
func (c *Client) GetLinks(request *LinksRequest) (*Links, error) {
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

// NewLinksRequest creates a new LinkRequest.
func NewLinksRequest(id, value string) *LinksRequest {
	return &LinksRequest{
		id:    id,
		value: value,
	}
}

// Validate ensures the request is valid and satisfies the Request interface.
func (req *LinksRequest) Validate() error {
	if len(req.id) == 0 {
		return errors.New("request must specify an ID to query by")
	}

	if len(req.value) == 0 {
		return errors.New("request must specify an ID value to query by")
	}

	return nil
}

func (req *LinksRequest) internal() {}
