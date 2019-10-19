package profile

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
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

// ExternalIDsRequest comprises the data necessary for retrieving a profile's external ID's from the Profile API.
type ExternalIDsRequest struct {
	// mandatory fields
	id    string
	value string

	// optional params
	queryParmas url.Values
}

// GetExternalIDs queries the Profile API for the given ID's externalID's.
func (c *Client) GetExternalIDs(request *ExternalIDsRequest) (*ExternalIDs, error) {
	url := baseURL + c.namespaceID + usersCollection + request.id + ":" + request.value + "/external_ids"
	if len(request.queryParmas) > 0 {
		url = url + request.queryParmas.Encode()
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

	externalIDs := newExternalIDs()
	dec := json.NewDecoder(res.Body)

	err = dec.Decode(externalIDs)
	if err != nil {
		return nil, err
	}

	fmt.Println(externalIDs)
	return externalIDs, nil
}

// NewExternalIDsRequest constructs a new ExternalIDsRequest with the given ID and value.
func NewExternalIDsRequest(id, value string) *ExternalIDsRequest {
	return &ExternalIDsRequest{
		id:          id,
		value:       value,
		queryParmas: url.Values{},
	}
}

// SetVerbose sets the ExternalIDsRequest's verbose query parameter to 'true'.
func (req *ExternalIDsRequest) SetVerbose() {
	req.queryParmas.Set("verbose", "true")
}

// SetInclude sets the type of ExternalIDs to include in the response from the Profile API.
func (req *ExternalIDsRequest) SetInclude(ids ...string) error {
	if len(ids) == 0 {
		return errors.New("must have at least 1 argument to req.SetInclude()")
	}

	var include []string
	for _, id := range ids {
		include = append(include, id)
	}

	req.queryParmas.Set("include", strings.Join(include, ","))
	return nil
}

// SetLimit sets the limit for the number of ExternalIDs that will be returned by the Profile API.
func (req *ExternalIDsRequest) SetLimit(limit int) error {
	if limit < 1 || limit > 100 {
		return errors.New("limit must be between 1 and 100, inclusive")
	}

	req.queryParmas.Set("limit", string(limit))
	return nil
}

// Validate ensures the request is valid and satisfies the Request interface.
func (req *ExternalIDsRequest) Validate() error {
	if len(req.id) == 0 {
		return errors.New("request must specify an ID to query by")
	}

	if len(req.value) == 0 {
		return errors.New("request must specify an ID value to query by")
	}

	return nil
}

func newExternalIDs() *ExternalIDs {
	c := &Cursor{}
	return &ExternalIDs{
		Data:   []*ExternalID{},
		Cursor: c,
	}
}
