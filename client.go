package gitea

import "net/http"
import "time"
import "github.com/dghubble/sling"

// Client represents an HTTP client to communicate with the Gitea server's API with.
type Client struct {
	baseURL string
	token   string
	sling   *sling.Sling
	client  *http.Client

	Admin *AdminService
	Users *UserService
}

func NewClient(token string, baseURL string) *Client {
	s := sling.New().Base(baseURL)
	return &Client{
		baseURL: baseURL,
		token:   token,
		sling:   s,
		client: &http.Client{
			Timeout: time.Second * 10,
		},

		Users: NewUserService(s.New()),
	}
}
