package updater

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"
)

// Update updates the Resource Record from CloudFlare using the information from the parent DNSRecordRequest struct
//
// Arguments:
//     None
//
// Returns:
//     (*DNSRecordResponse): The response from CloudFlare if successful, nil otherwise
//     (error):              An error if one exists, nil otherwise
func (U *UpdateDNSRecordRequest) Update() (*DNSRecordResponse, error) {
	client := http.DefaultClient
	response := new(DNSRecordResponse)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	j, err := U.Marshal()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, U.URL, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", U.APIToken))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if err := response.Unmarshal(resp.Body); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cloudflare responded with %v", resp.StatusCode)
	}

	return response, nil
}
