package profile

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
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

// MetadataRequest comprises the data necessary to retrieve a profile's Metadata from the Profile API.
type MetadataRequest struct {
	// mandatory fields
	id    string
	value string

	// optional query params
	queryParams url.Values
}

// getMetadata queries the Profile API for the given ID's metadata.
func (c *Client) getMetadata(request *MetadataRequest) (*Metadata, error) {
	url := baseURL + c.namespaceID + usersCollection + request.id + ":" + request.value + "/metadata"
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

	var metadata Metadata
	dec := json.NewDecoder(res.Body)

	err = dec.Decode(&metadata)
	if err != nil {
		return nil, err
	}

	fmt.Println(metadata)
	return &metadata, nil
}

// Marshal marshals the Metadata into JSON and satisfies the Response interface.
func (m *Metadata) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// NewMetadataRequest constructs a new MetadataRequest with the given ID and value.
func NewMetadataRequest(id, value string) *MetadataRequest {
	return &MetadataRequest{
		id:          id,
		value:       value,
		queryParams: url.Values{},
	}
}

// SetVerbose sets the MetadataRequest's verbose query parameter to 'true'.
func (req *MetadataRequest) SetVerbose() {
	req.queryParams.Set("verbose", "true")
}

// Validate ensures the MetadataRequest is valid and satisfies the Request interface.
func (req *MetadataRequest) Validate() error {
	if len(req.id) == 0 {
		return errors.New("request must specify an ID to query by")
	}

	if len(req.value) == 0 {
		return errors.New("request must specify an ID value to query by")
	}

	return nil
}

func (req *MetadataRequest) internal() {}
func (m *Metadata) internal()          {}
