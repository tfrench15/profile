package profile

import (
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
func (c *Client) Query(request Request) (Response, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	switch q := request.(type) {
	case *TraitsRequest:
		return c.getTraits(q)

	case *EventRequest:
		return c.getEvents(q)

	case *ExternalIDsRequest:
		return c.getExternalIDs(q)

	case *MetadataRequest:
		return c.getMetadata(q)

	case *LinksRequest:
		return c.getLinks(q)

	default:
		return nil, errors.New("could not execute Query")
	}
}
