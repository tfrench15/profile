package profile

import (
	"encoding/json"
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

// TraitsRequestConfig allows the client to pass in additional query parameters and customize
// its request.
type TraitsRequestConfig struct {
	Include []string
	Verbose bool
	Limit   int
}

// GetTraits queries the Profile API for the provided ID's traits.
func (c *Client) GetTraits(id, value string, config *TraitsRequestConfig) (*Traits, error) {
	url := baseURL + c.namespaceID + usersCollection + id + ":" + value + "/traits"

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

// NewTraitsConfig returns a new configuration struct for further customizing Traits requests.
func NewTraitsConfig(traits []string, verbose bool, limit int) (*TraitsRequestConfig, error) {
	if limit < 1 || limit > 100 {
		return nil, fmt.Errorf("limit %d out of accepted range, must be between 1 and 100", limit)
	}

	config := &TraitsRequestConfig{
		Include: traits,
		Verbose: verbose,
		Limit:   limit,
	}

	return config, nil
}

// Encode returns the URL encoding of the query parameters contained in the TraitsRequestConfig.
func (config *TraitsRequestConfig) Encode() string {
	var verbose string
	if config.Verbose {
		verbose = "true"
	}
	verbose = "false"

	params := url.Values{}

	params.Set("verbose", verbose)
	params.Set("limit", string(config.Limit))
	params.Set("include", strings.Join(config.Include, ","))

	return params.Encode()
}

func newTraits() *Traits {
	c := &Cursor{}
	m := make(map[string]interface{})

	return &Traits{
		Traits: m,
		Cursor: c,
	}
}
