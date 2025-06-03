package medialog_client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (mlc *MedialogClient) GetEntryIDs() ([]uuid.UUID, error) {

	entryIDS := []uuid.UUID{}

	url := "/entries?all_ids=true"

	resp, err := mlc.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get entries: %s", resp.Status)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&entryIDS); err != nil {
		return nil, err
	}

	return entryIDS, nil
}
