package medialog_client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Token struct {
	Token string `json:"token"`
}

type MedialogClient struct {
	SessionToken string
	RootURL      string
	Client       *http.Client
}

type MedialogCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
	URL      string `json:"url"`
}

const timeout = 20

func GetClient(config string, environment string) (MedialogClient, error) {
	b, err := os.ReadFile(config)
	if err != nil {
		return MedialogClient{}, err
	}

	mlCreds, err := getCreds(environment, b)
	if err != nil {
		return MedialogClient{}, err
	}

	mlClient := MedialogClient{}
	url := fmt.Sprintf("%s/users/%s/login?password=%s", mlCreds.URL, mlCreds.Username, mlCreds.Password)

	mlClient.RootURL = mlCreds.URL

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    time.Duration(timeout) * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Transport: tr,
	}

	mlClient.Client = client

	resp, err := http.Post(url, "", nil)
	if err != nil {
		return MedialogClient{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return MedialogClient{}, err
	}

	token := Token{}
	if err := json.Unmarshal(body, &token); err != nil {
		return MedialogClient{}, err
	}

	mlClient.SessionToken = token.Token

	return mlClient, nil
}

func getCreds(environment string, configBytes []byte) (MedialogCreds, error) {
	credsMap := map[string]MedialogCreds{}

	err := yaml.Unmarshal(configBytes, &credsMap)
	if err != nil {
		return MedialogCreds{}, err
	}

	for k, v := range credsMap {
		if environment == k {
			return v, nil
		}
	}

	return MedialogCreds{}, fmt.Errorf("Credentials file did not contain %s\n", environment)
}

func (mlc *MedialogClient) Get(url string) (*http.Response, error) {

	req, err := http.NewRequest("GET", mlc.RootURL+url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Medialog-Token", mlc.SessionToken)

	resp, err := mlc.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
