package updater

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Get fetches the IP Address information from https://ipinfo.io using the information from the parent IPInfoRequest struct
//
// Arguments:
//     None
//
// Returns:
//     (*DNSRecordResponse): The response from IPInfo if successful, nil otherwise
//     (error):              An error if one exists, nil otherwise
func (I *IPInfoRequest) Get() (*IPInfoResponse, error) {
	response := new(IPInfoResponse)
	client := http.DefaultClient

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, I.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating IPInfoRequest http request, %w", err)
	}

	req.Header.Set("Accept", "application/json")

	// Allow no API Token requests
	if I.APIToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", I.APIToken))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending IPInfoRequest http request, %w", err)
	}

	if err := response.Unmarshal(resp.Body); err != nil {
		return nil, fmt.Errorf("error unmarshaling IPInfoResponse, %w", err)
	}

	return response, nil
}
