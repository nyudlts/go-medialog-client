package medialog_client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (c *MedialogClient) GetEntryIDs() ([]uuid.UUID, error) {

	entryIDS := []uuid.UUID{}

	url := c.RootURL + "/entries?all_ids=true"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Medialog-Token", c.SessionToken)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get entries: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&entryIDS); err != nil {
		return nil, err
	}

	return entryIDS, nil
}
