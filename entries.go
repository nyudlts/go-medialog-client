package medialog_client

import (
	"fmt"
	"io"
	"net/http"
)

func (c *MedialogClient) GetEntries() error {

	url := c.RootURL + "/entries"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-Medialog-Token", c.SessionToken)

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get entries: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Response Body:", string(body))
	return nil
}
