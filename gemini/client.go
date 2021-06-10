package gemini

import (
	"net/http"

	"github.com/sirupsen/logrus"
)


type Client struct {
	baseURLv1 string
	baseURLv2 string
	httpClient *http.Client
	logger *logrus.Logger
}

func New() Client {
	return Client{
		baseURLv1: "https://api.gemini.com/v1",
		baseURLv2: "https://api.gemini.com/v2",
		httpClient: http.DefaultClient,
		logger: logrus.New(),
	}
}

func (c *Client) RoundTrip(req *http.Request) (*http.Response, error) {
	r := req.Clone(req.Context())
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Accept-Charset", "UTF-8")

	return http.DefaultTransport.RoundTrip(r)
}
