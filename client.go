package profile

import "net/http"

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
func (c *Client) Query(request Request) {
	switch q := request.(type) {
	case *TraitsRequest:
		c.getTraits(q)
	case *EventRequest:
		c.getEvents(q)
	case *ExternalIDsRequest:
		c.getExternalIDs(q)
	case *MetadataRequest:
		c.getMetadata(q)
	case *LinksRequest:
		c.getLinks(q)
	}
}
