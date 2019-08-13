package profile

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Traits holds the traits for the requested Profile.
type Traits struct {
	Traits map[string]interface{} `json:"traits"`
	Cursor *Cursor                `json:"cursor"`
}

// GetTraits queries the Profile API for the provided ID's traits.
func (c *Client) GetTraits(id, value string) (*Traits, error) {
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

	err = dec.Decode(traits.Cursor)
	if err != nil {
		return nil, err
	}

	err = dec.Decode(traits.Traits)
	if err != nil {
		return nil, err
	}

	fmt.Println(traits)
	return traits, nil
}

func newTraits() *Traits {
	c := &Cursor{}
	m := make(map[string]interface{})

	return &Traits{
		Traits: m,
		Cursor: c,
	}
}
