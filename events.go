package profile

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetEvents queries the Profile API for the given ID's events.
func (c *Client) GetEvents(id, value string) error {
	url := baseURL + c.namespaceID + usersCollection + id + ":" + value + "/events"

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
