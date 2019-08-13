package profile

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Property contains the key-value pair comprising a Segment property.
type Property struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Properties contains a slice of Property.
type Properties struct {
	Data   []*Property `json:"data"`
	Cursor *Cursor     `json:"cursor"`
}

// GetProperties queries the Profile API for the given ID's properties.
func (c *Client) GetProperties(id, value string) (*Properties, error) {
	url := baseURL + c.namespaceID + usersCollection + id + ":" + value + "/properties"

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

	properties := newProperties()
	dec := json.NewDecoder(res.Body)

	err = dec.Decode(properties)
	if err != nil {
		return nil, err
	}

	fmt.Println(properties)
	return properties, nil
}

func newProperties() *Properties {
	c := &Cursor{}
	return &Properties{
		Data:   []*Property{},
		Cursor: c,
	}
}
