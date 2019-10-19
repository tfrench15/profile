package profile

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	baseURL         = "https://profiles.segment.com/v1/spaces/"
	usersCollection = "/collections/users/profiles/"
)

// Client represents a client to the Personas API.
type Client struct {
	client *http.Client

	// identifiers
	namespaceID string
	secret      string // TODO: make this Environment variable
}

// New returns a new Client.
func New(namespaceID, secret string) *Client {
	return &Client{
		client:      &http.Client{},
		namespaceID: namespaceID,
		secret:      secret,
	}
}

// Query executes a request against the Profile API.
func (c *Client) Query(request Request) ([]byte, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	switch q := request.(type) {
	case *TraitsRequest:
		traits, err := c.getTraits(q)
		if err != nil {
			return nil, err
		}

		return json.Marshal(traits)

	case *EventRequest:
		events, err := c.getEvents(q)
		if err != nil {
			return nil, err
		}

		return json.Marshal(events)

	case *ExternalIDsRequest:
		externalIDs, err := c.getExternalIDs(q)
		if err != nil {
			return nil, err
		}

		return json.Marshal(externalIDs)

	case *MetadataRequest:
		metadata, err := c.getMetadata(q)
		if err != nil {
			return nil, err
		}

		return json.Marshal(metadata)

	case *LinksRequest:
		links, err := c.getLinks(q)
		if err != nil {
			return nil, err
		}

		return json.Marshal(links)

	default:
		return nil, errors.New("could not execute Query")
	}
}
