package profile

import "net/http"

const (
	baseURL         = "https://profiles.segment.com/v1/spaces/"
	usersCollection = "/collections/users/profiles/"
)

// Client represents a client to the Personas API.
type Client struct {
	client *http.Client

	namespaceID string
	secret      string // TODO: make this Environment variable
}

// New returns a new Client.
func New(namespaceID string) *Client {
	secret := "TODO"

	return &Client{
		client:      &http.Client{},
		namespaceID: namespaceID,
		secret:      secret,
	}
}