package updater

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Get fetches the Resource Record from CloudFlare using the information from the parent DNSRecordRequest struct
//
// Arguments:
//     None
//
// Returns:
//     (*DNSRecordResponse): The response from CloudFlare if successful, nil otherwise
//     (error):              An error if one exists, nil otherwise
func (D *DNSRecordRequest) Get() (*DNSRecordResponse, error) {
	client := http.DefaultClient
	response := new(DNSRecordResponse)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, D.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating DNSRecordRequest http request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", D.APIToken))

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending DNSRecordRequest http request: %w", err)
	}

	if err := response.Unmarshal(resp.Body); err != nil {
		return nil, fmt.Errorf("error unmarshaling DNSRecordResponse: %w", err)
	}

	if !response.Success {
		return nil, fmt.Errorf("errors: %v", response.Errors)
	}

	return response, nil
}
