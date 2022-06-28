package conquest_api

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/client"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type Config struct {
	Address            string `json:"address"`
	InsecureSkipVerify bool   `json:"insecure"`
	Debug              bool   `json:"debug"`
	AccessToken        string `json:"access_token"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, loadErr(filename, err)
	}
	var c Config
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, loadErr(filename, err)
	}
	return &c, nil
}

func loadErr(filename string, err error) error {
	return fmt.Errorf("could not read the Conquest API configuration from the file '%s', err:\n%s", filename, err.Error())
}

func useAccessToken(accessToken string) runtime.ClientAuthInfoWriterFunc {
	return func(r runtime.ClientRequest, _ strfmt.Registry) error {
		return r.SetHeaderParam("Authorization", "Bearer "+accessToken)
	}
}

func NewClient(config Config) (*client.ConquestAPIV4, error) {
	u, err := url.Parse(config.Address)
	if err != nil {
		return nil, fmt.Errorf("invalid Conquest API 'address'")
	}
	if u.Scheme != "https" {
		return nil, fmt.Errorf("invalid Conquest API 'address', must use https")
	}

	var transport *httptransport.Runtime

	// If you are using a self-signed dev certificate, skip the cert check
	if config.InsecureSkipVerify {
		transport = httptransport.NewWithClient(u.Host, u.Path, []string{"https"}, &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: config.InsecureSkipVerify,
				},
			},
		})
	} else {
		transport = httptransport.New(u.Host, u.Path, []string{"https"})
	}

	transport.DefaultAuthentication = useAccessToken(config.AccessToken)

	transport.SetDebug(config.Debug)

	return client.New(transport, nil), nil
}

func DefaultConfig() (*Config, error) {
	return LoadConfig("conquest_api.credentials.json")
}

func New() (*client.ConquestAPIV4, error) {
	cfg, err := DefaultConfig()
	if err != nil {
		return nil, err
	}
	return NewClient(*cfg)
}

var _ runtime.ClientResponseReaderFunc
