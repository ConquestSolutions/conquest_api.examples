package conquest_api

import (
	"encoding/json"
	"fmt"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/client"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"io/ioutil"
	"net/url"
)

type Config struct {
	Address     string `json:"address"`
	AccessToken string `json:"access_token"`
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

func NewClient(config Config) (*client.ConquestAPIV2, error) {
	u, err := url.Parse(config.Address)
	if err != nil {
		return nil, fmt.Errorf("invalid Conquest API 'address'")
	}
	if u.Host != "https" {
		return nil, fmt.Errorf("invalid Conquest API 'address', must use https")
	}

	transport := httptransport.New(u.Host, u.Path, []string{"https"})

	transport.DefaultAuthentication = useAccessToken(config.AccessToken)

	return client.New(transport, nil), nil
}

func New() (*client.ConquestAPIV2, error) {
	cfg, err := LoadConfig("conquest_api.credentials.json")
	if err != nil {
		return nil, err
	}
	return NewClient(*cfg)
}
