package profile

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ExternalID represents an ExternalID in the Personas Identity Graph.
type ExternalID struct {
	SourceID       string `json:"source_id"`
	Collection     string `json:"collection"`
	ID             string `json:"id"`
	Type           string `json:"type"`
	CreatedAt      string `json:"created_at"`
	Encoding       string `json:"encoding"`
	FirstMessageID string `json:"first_message_id"`
}

// ExternalIDs comprise an array of returned ExternalID from a request to the Profile API.
type ExternalIDs struct {
	Data   []*ExternalID
	Cursor *Cursor
}

// GetExternalIDs queries the Profile API for the given ID's externalID's.
func (c *Client) GetExternalIDs(id, value string) (*ExternalIDs, error) {
	url := baseURL + c.namespaceID + usersCollection + id + ":" + value + "/external_ids"

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

	externalIDs := newExternalIDs()
	dec := json.NewDecoder(res.Body)

	err = dec.Decode(externalIDs)
	if err != nil {
		return nil, err
	}

	fmt.Println(externalIDs)
	return externalIDs, nil
}

func newExternalIDs() *ExternalIDs {
	c := &Cursor{}
	return &ExternalIDs{
		Data:   []*ExternalID{},
		Cursor: c,
	}
}
