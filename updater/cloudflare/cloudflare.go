package updater

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"time"
)

// CloudFlareRequest is a general interface for all Requests to CloudFlare
type CloudFlareRequest interface {
	Marshal() ([]byte, error)
}

// CloudFlareResponse is a general interface for all Responses from CloudFlare
type CloudFlareResponse interface {
	Unmarshal(body io.ReadCloser) error
}

// DNSRecordRequest is a struct that holds all the data required to marshal info a CloudFlare DNS Record Request
type DNSRecordRequest struct {
	URL        string
	APIToken   string
	ZoneID     string
	RecordName string
}

// Marshal marshals the contents of the parent DNSRecordRequest struct
//
// Arguments:
//     None
//
// Returns:
//     ([]byte): The marshaled bytes if successful, nil otherwise
//     (error):  An error if one exists, nil otherwise
func (D *DNSRecordRequest) Marshal() ([]byte, error) {
	return json.Marshal(D)
}

// DNSRecordResponse is a struct that holds all the data associated with a CloudFlare DNS Record Response
type DNSRecordResponse struct {
	Result []struct {
		ID         string    `json:"id"`
		Type       string    `json:"type"`
		Name       string    `json:"name"`
		Content    string    `json:"content"`
		Proxiable  bool      `json:"proxiable"`
		Proxied    bool      `json:"proxied"`
		TTL        int       `json:"ttl"`
		Locked     bool      `json:"locked"`
		ZoneID     string    `json:"zone_id"`
		ZoneName   string    `json:"zone_name"`
		ModifiedOn time.Time `json:"modified_on"`
		CreatedOn  time.Time `json:"created_on"`
		Meta       struct {
			AutoAdded           bool `json:"auto_added"`
			ManagedByApps       bool `json:"managed_by_apps"`
			ManagedByArgoTunnel bool `json:"managed_by_argo_tunnel"`
		} `json:"meta"`
	} `json:"result"`
	ResultInfo struct {
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		TotalPages int `json:"total_pages"`
		Count      int `json:"count"`
		TotalCount int `json:"total_count"`
	} `json:"result_info"`
	Success bool `json:"success"`
	Errors  []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
	Messages []string `json:"messages"`
}

// Unmarshal unmarshals the contents of the parent DNSRecordResponse struct
//
// Arguments:
//     body (io.ReadCloser): The http response body
//
// Returns:
//     (error):  An error if one exists, nil otherwise
func (D *DNSRecordResponse) Unmarshal(body io.ReadCloser) error {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, D)
}

// UpdateDNSRecordRequest is a struct that contains the data to make a CloudFlare DNS Records Update Request
type UpdateDNSRecordRequest struct {
	URL        string `json:"-"`
	APIToken   string `json:"-"`
	ZoneID     string `json:"-"`
	RecordName string `json:"-"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	Content    string `json:"content"`
	TTL        int    `json:"ttl"`
	Proxied    bool   `json:"proxied"`
}

// Marshal marshals the contents of the parent DNSRecordRequest struct
//
// Arguments:
//     None
//
// Returns:
//     ([]byte): The marshaled bytes if successful, nil otherwise
//     (error):  An error if one exists, nil otherwise
func (U *UpdateDNSRecordRequest) Marshal() ([]byte, error) {
	return json.Marshal(U)
}
