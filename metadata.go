package profile

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Metadata holds the metadata for the requested Profile.
type Metadata struct {
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	ExpiresAt      string `json:"expires_at"`
	FirstMessageID string `json:"first_message_id"`
	FirstSourceID  string `json:"first_source_id"`
	LastMessageID  string `json:"last_message_id"`
	LastSourceID   string `json:"last_source_id"`
}

// GetMetadata queries the Profile API for the given ID's metadata.
func (c *Client) GetMetadata(id, value string) (*Metadata, error) {
	url := baseURL + c.namespaceID + usersCollection + id + ":" + value + "/metadata"

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

	var metadata Metadata
	dec := json.NewDecoder(res.Body)

	err = dec.Decode(&metadata)
	if err != nil {
		return nil, err
	}

	fmt.Println(metadata)
	return &metadata, nil
}
