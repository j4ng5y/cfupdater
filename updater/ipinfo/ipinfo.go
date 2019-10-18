package updater

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

// IPInfoRequest is a struct that contains the data required to marshal a request to https://ipinfo.io
type IPInfoRequest struct {
	URL      string
	APIToken string
}

// IPInfoResponse is a struct that contains the data required to unmarshal a response from https://ipinfo.io
type IPInfoResponse struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

// Unmarshal unmarshals the contents of the parent IPInfoResponse struct
//
// Arguments:
//     data ([]byte): The bytes to unmarshal
//
// Returns:
//     (error):  An error if one exists, nil otherwise
func (I *IPInfoResponse) Unmarshal(body io.ReadCloser) error {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return fmt.Errorf("error reading http body, %w", err)
	}

	if err := json.Unmarshal(data, I); err != nil {
		return fmt.Errorf("error unmarshaling http body to IPInfoResponse, %w", err)
	}
	return nil
}
