package profile

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Traits holds the traits for the requested Profile.
type Traits struct {
	Traits map[string]interface{} `json:"traits"`
	Cursor *Cursor                `json:"cursor"`
}

// TraitsRequest comprises the data necessary for retrieving a profile's traits from the Profile API.
type TraitsRequest struct {
	// mandatory fields
	id    string
	value string

	// optional params
	queryParams url.Values
}

// getTraits queries the Profile API for the provided ID's traits.
func (c *Client) getTraits(request *TraitsRequest) (*Traits, error) {
	url := baseURL + c.namespaceID + usersCollection + request.id + ":" + request.value + "/traits"
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

	traits := newTraits()
	dec := json.NewDecoder(res.Body)

	err = dec.Decode(traits)
	if err != nil {
		return nil, err
	}

	fmt.Println(traits)
	return traits, nil
}

// Marshal marshals Traits into JSON and satisfies the Response interface.
func (t *Traits) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

// NewTraitsRequest returns a new configuration struct for further customizing Traits requests.
func NewTraitsRequest(id, value string) *TraitsRequest {
	return &TraitsRequest{
		id:          id,
		value:       value,
		queryParams: url.Values{},
	}
}

// SetVerbose sets the TraitRequest's verbose query parameter to 'true'.
func (req *TraitsRequest) SetVerbose() {
	req.queryParams.Set("verbose", "true")
}

// SetLimit sets the TraitRequest's limit query paramaeter.
func (req *TraitsRequest) SetLimit(limit int) error {
	if limit < 1 || limit > 100 {
		return fmt.Errorf("limit must be at least 1 and at most 100, got %d", limit)
	}

	req.queryParams.Set("limit", string(limit))
	return nil
}

// SetInclude sets the TraitRequest's include query parameter.
func (req *TraitsRequest) SetInclude(traits ...string) error {
	if len(traits) == 0 {
		return errors.New("cannot pass in 0 arguments to SetInclude")
	}

	var include []string
	for _, trait := range traits {
		include = append(include, trait)
	}

	req.queryParams.Set("include", strings.Join(include, ","))
	return nil
}

// Validate ensures the request is valid and satisfies the Request interface.
func (req *TraitsRequest) Validate() error {
	if len(req.id) == 0 {
		return errors.New("request must specify an ID to query by")
	}

	if len(req.value) == 0 {
		return errors.New("request must specify an ID value to query by")
	}

	return nil
}

func newTraits() *Traits {
	c := &Cursor{}
	m := make(map[string]interface{})

	return &Traits{
		Traits: m,
		Cursor: c,
	}
}

func (req *TraitsRequest) internal() {}
func (t *Traits) internal()          {}
