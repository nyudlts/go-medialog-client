package medialog_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/nyudlts/go-medialog/models"
)

func (mlc *MedialogClient) GetEntry(entryID uuid.UUID) (*models.Entry, error) {

	url := fmt.Sprintf("/entries/%s", entryID.String())

	resp, err := mlc.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get entry %s: %s", entryID.String(), resp.Status)
	}

	e := &models.Entry{}

	if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
		return nil, fmt.Errorf("failed to decode entry %s: %w", entryID.String(), err)
	}

	return e, nil
}

func (mlc *MedialogClient) UpdateEntry(entryID uuid.UUID, entry *models.Entry) (string, error) {
	url := fmt.Sprintf("/entries/%s/update", entryID.String())

	reqBody, err := json.Marshal(entry)
	if err != nil {
		return "", fmt.Errorf("failed to marshal entry %s: %w", entryID.String(), err)
	}

	resp, err := mlc.Post(url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to update entry %s: %s", entryID.String(), resp.Status)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body for entry %s: %w", entryID.String(), err)
	}

	return string(respBody), nil
}

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
