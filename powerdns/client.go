package powerdns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/hashicorp/go-cleanhttp"
)

type Client struct {
	ServerUrl  string // Location of PowerDNS server to use
	ApiKey     string // REST API Static authentication key
	ApiVersion int    // API version to use
	Http       *http.Client
}

// NewClient returns a new PowerDNS client
func NewClient(serverUrl string, apiKey string) (*Client, error) {
	client := Client{
		ServerUrl: serverUrl,
		ApiKey:    apiKey,
		Http:      cleanhttp.DefaultClient(),
	}
	var err error
	client.ApiVersion, err = client.detectApiVersion()
	if err != nil {
		return nil, err
	}
	return &client, nil
}

// Creates a new request with necessary headers
func (c *Client) newRequest(method string, endpoint string, body []byte) (*http.Request, error) {

	var urlStr string
	if c.ApiVersion > 0 {
		urlStr = c.ServerUrl + "/api/v" + strconv.Itoa(c.ApiVersion) + endpoint
	} else {
		urlStr = c.ServerUrl + endpoint
	}
	url, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("Error during parsing request URL: %s", err)
	}

	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}

	req, err := http.NewRequest(method, url.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("Error during creation of request: %s", err)
	}

	req.Header.Add("X-API-Key", c.ApiKey)
	req.Header.Add("Accept", "application/json")

	if method != "GET" {
		req.Header.Add("Content-Type", "application/json")
	}

	return req, nil
}

type ZoneInfo struct {
	Id                 string              `json:"id"`
	Name               string              `json:"name"`
	URL                string              `json:"url"`
	Kind               string              `json:"kind"`
	DnsSec             bool                `json:"dnsssec"`
	Serial             int64               `json:"serial"`
	Records            []Record            `json:"records,omitempty"`
	ResourceRecordSets []ResourceRecordSet `json:"rrsets,omitempty"`
	Nameservers        []string            `json:"nameservers,omitempty"`
}

type Record struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Content  string `json:"content"`
	TTL      int    `json:"ttl"` // For API v0
	Disabled bool   `json:"disabled"`
}

type ResourceRecordSet struct {
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	ChangeType string   `json:"changetype"`
	TTL        int      `json:"ttl"` // For API v1
	Records    []Record `json:"records,omitempty"`
}

type zonePatchRequest struct {
	RecordSets []ResourceRecordSet `json:"rrsets"`
}

type errorResponse struct {
	ErrorMsg string `json:"error"`
}

const idSeparator string = ":::"

func (record *Record) Id() string {
	return record.Name + idSeparator + record.Type
}

func (rrSet *ResourceRecordSet) Id() string {
	return rrSet.Name + idSeparator + rrSet.Type
}

// Returns name and type of record or record set based on its ID
func parseId(recId string) (string, string, error) {
	s := strings.Split(recId, idSeparator)
	if len(s) == 2 {
		return s[0], s[1], nil
	}
	return "", "", fmt.Errorf("Unknown record ID format")
}

// Detects the API version in use on the server
// Uses int to represent the API version: 0 is the legacy AKA version 3.4 API
// Any other integer correlates with the same API version
func (client *Client) detectApiVersion() (int, error) {
	req, err := client.newRequest("GET", "/api/v1/servers", nil)
	if err != nil {
		return -1, err
	}
	resp, err := client.Http.Do(req)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return 1, nil
	}
	return 0, nil
}

// Returns all Zones of server, without records
func (client *Client) ListZones() ([]ZoneInfo, error) {
	req, err := client.newRequest("GET", "/servers/localhost/zones", nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var zoneInfos []ZoneInfo

	err = json.NewDecoder(resp.Body).Decode(&zoneInfos)
	if err != nil {
		return nil, err
	}

	return zoneInfos, nil
}

// GetZone gets a zone
func (c *Client) GetZone(name string) (ZoneInfo, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/servers/localhost/zones/%s", name), nil)
	if err != nil {
		return ZoneInfo{}, err
	}

	resp, err := c.Http.Do(req)
	if err != nil {
		return ZoneInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorResp := new(errorResponse)
		if err = json.NewDecoder(resp.Body).Decode(errorResp); err != nil {
			return ZoneInfo{}, fmt.Errorf("Error getting zone: %s", name)
		}
		return ZoneInfo{}, fmt.Errorf("Error getting zone: %s, reason: %q", name, errorResp.ErrorMsg)
	}

	var zoneInfo ZoneInfo
	err = json.NewDecoder(resp.Body).Decode(&zoneInfo)
	if err != nil {
		return ZoneInfo{}, err
	}

	return zoneInfo, nil
}

// ZoneExists checks if requested zone exists
func (c *Client) ZoneExists(name string) (bool, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/servers/localhost/zones/%s", name), nil)
	if err != nil {
		return false, err
	}

	resp, err := c.Http.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotFound {
		errorResp := new(errorResponse)
		if err = json.NewDecoder(resp.Body).Decode(errorResp); err != nil {
			return false, fmt.Errorf("Error getting zone: %s", name)
		}
		return false, fmt.Errorf("Error getting zone: %s, reason: %q", name, errorResp.ErrorMsg)
	}

	return resp.StatusCode == http.StatusOK, nil
}

func (c *Client) CreateZone(zoneInfo ZoneInfo) error {
	body, err := json.Marshal(zoneInfo)
	if err != nil {
		return err
	}

	req, err := c.newRequest("POST", "/servers/localhost/zones", body)
	if err != nil {
		return err
	}

	resp, err := c.Http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		errorResp := new(errorResponse)
		if err = json.NewDecoder(resp.Body).Decode(errorResp); err != nil {
			return fmt.Errorf("Error creating zone: %s", zoneInfo.Name)
		}
		return fmt.Errorf("Error creating zone: %s, reason: %q", zoneInfo.Name, errorResp.ErrorMsg)
	}

	return nil
}

// DeleteZone deletes a zone
func (c *Client) DeleteZone(name string) error {
	req, err := c.newRequest("DELETE", fmt.Sprintf("/servers/localhost/zones/%s", name), nil)
	if err != nil {
		return err
	}

	resp, err := c.Http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		errorResp := new(errorResponse)
		if err = json.NewDecoder(resp.Body).Decode(errorResp); err != nil {
			return fmt.Errorf("Error deleting zone: %s", name)
		}
		return fmt.Errorf("Error deleting zone: %s, reason: %q", name, errorResp.ErrorMsg)
	}
	return nil
}

// Returns all records in Zone
func (client *Client) ListRecords(zone string) ([]Record, error) {
	req, err := client.newRequest("GET", fmt.Sprintf("/servers/localhost/zones/%s", zone), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	zoneInfo := new(ZoneInfo)
	err = json.NewDecoder(resp.Body).Decode(zoneInfo)
	if err != nil {
		return nil, err
	}

	records := zoneInfo.Records
	// Convert the API v1 response to v0 record structure
	for _, rrs := range zoneInfo.ResourceRecordSets {
		for _, record := range rrs.Records {
			records = append(records, Record{
				Name:    rrs.Name,
				Type:    rrs.Type,
				Content: record.Content,
				TTL:     rrs.TTL,
			})
		}
	}

	return records, nil
}

// Returns only records of specified name and type
func (client *Client) ListRecordsInRRSet(zone string, name string, tpe string) ([]Record, error) {
	allRecords, err := client.ListRecords(zone)
	if err != nil {
		return nil, err
	}

	records := make([]Record, 0, 10)
	for _, r := range allRecords {
		if r.Name == name && r.Type == tpe {
			records = append(records, r)
		}
	}

	return records, nil
}

func (client *Client) ListRecordsByID(zone string, recId string) ([]Record, error) {
	name, tpe, err := parseId(recId)
	if err != nil {
		return nil, err
	}
	return client.ListRecordsInRRSet(zone, name, tpe)
}

// Checks if requested record exists in Zone
func (client *Client) RecordExists(zone string, name string, tpe string) (bool, error) {
	allRecords, err := client.ListRecords(zone)
	if err != nil {
		return false, err
	}

	for _, record := range allRecords {
		if record.Name == name && record.Type == tpe {
			return true, nil
		}
	}
	return false, nil
}

// Checks if requested record exists in Zone by it's ID
func (client *Client) RecordExistsByID(zone string, recId string) (bool, error) {
	name, tpe, err := parseId(recId)
	if err != nil {
		return false, err
	}
	return client.RecordExists(zone, name, tpe)
}

// Creates new record set in Zone
func (client *Client) ReplaceRecordSet(zone string, rrSet ResourceRecordSet) (string, error) {
	rrSet.ChangeType = "REPLACE"

	reqBody, _ := json.Marshal(zonePatchRequest{
		RecordSets: []ResourceRecordSet{rrSet},
	})

	req, err := client.newRequest("PATCH", fmt.Sprintf("/servers/localhost/zones/%s", zone), reqBody)
	if err != nil {
		return "", err
	}

	resp, err := client.Http.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		errorResp := new(errorResponse)
		if err = json.NewDecoder(resp.Body).Decode(errorResp); err != nil {
			return "", fmt.Errorf("Error creating record set: %s", rrSet.Id())
		}
		return "", fmt.Errorf("Error creating record set: %s, reason: %q", rrSet.Id(), errorResp.ErrorMsg)
	}
	return rrSet.Id(), nil
}

// Deletes record set from Zone
func (client *Client) DeleteRecordSet(zone string, name string, tpe string) error {
	reqBody, _ := json.Marshal(zonePatchRequest{
		RecordSets: []ResourceRecordSet{
			{
				Name:       name,
				Type:       tpe,
				ChangeType: "DELETE",
			},
		},
	})

	req, err := client.newRequest("PATCH", fmt.Sprintf("/servers/localhost/zones/%s", zone), reqBody)
	if err != nil {
		return err
	}

	resp, err := client.Http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		errorResp := new(errorResponse)
		if err = json.NewDecoder(resp.Body).Decode(errorResp); err != nil {
			return fmt.Errorf("Error deleting record: %s %s", name, tpe)
		}
		return fmt.Errorf("Error deleting record: %s %s, reason: %q", name, tpe, errorResp.ErrorMsg)
	}
	return nil
}

// Deletes record from Zone by its ID
func (client *Client) DeleteRecordSetByID(zone string, recId string) error {
	name, tpe, err := parseId(recId)
	if err != nil {
		return err
	}
	return client.DeleteRecordSet(zone, name, tpe)
}
