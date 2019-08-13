package profile

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Property contains the key-value pair comprising a Segment property.
type Property struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// GetProperties queries the Profile API for the given ID's properties.
func (c *Client) GetProperties(id, value string) error {
	url := baseURL + c.namespaceID + usersCollection + id + ":" + value + "/properties"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.secret, "")

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
