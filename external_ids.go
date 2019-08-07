package profile

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetExternalIDs queries the Profile API for the given ID's externalID's.
func (c *Client) GetExternalIDs(id, value string) error {
	url := baseURL + c.namespaceID + usersCollection + id + ":" + value + "/external_ids"

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
