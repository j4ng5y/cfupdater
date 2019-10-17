package updater

import (
	"context"
	"net/http"
	"time"
)

// Get fetches the IP Address information from https://ipinfo.io using the information from the parent IPInfoRequest struct
//
// Arguments:
//     None
//
// Returns:
//     (*DNSRecordResponse): The response from IPInfo if sucessful, nil otherwise
//     (error):              An error if one exists, nil otherwise
func (I *IPInfoRequest) Get() (*IPInfoResponse, error) {
	response := new(IPInfoResponse)
	client := http.DefaultClient

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, I.URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if err := response.Unmarshal(resp.Body); err != nil {
		return nil, err
	}

	return response, nil
}
