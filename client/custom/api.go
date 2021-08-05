// Package bitlinks contains the methods to interact with the Bitlinks in Bitly
package custom

import "encoding/json"

type CustomBitlinkRequest struct {
	CustomBitlink string `json:"custom_bitlink"`
	BitlinkID     string `json:"bitlink_id"`
}

type UpdateCustomBitlinkRequest struct {
	BitlinkID string `json:"bitlink_id"`
}

type CustomBitlink struct {
	CustomBitlink  string           `json:"custom_bitlink"`
	Bitlink        Bitlink          `json:"bitlink"`
	BitlinkHistory []BitlinkHistory `json:"bitlink_history"`
}

type Deeplinks struct {
	GUID        string `json:"guid"`
	Bitlink     string `json:"bitlink"`
	AppURIPath  string `json:"app_uri_path"`
	InstallURL  string `json:"install_url"`
	AppGUID     string `json:"app_guid"`
	Os          string `json:"os"`
	InstallType string `json:"install_type"`
	Created     string `json:"created"`
	Modified    string `json:"modified"`
	BrandGUID   string `json:"brand_guid"`
}
type Bitlink struct {
	Link           string      `json:"link"`
	ID             string      `json:"id"`
	LongURL        string      `json:"long_url"`
	Title          string      `json:"title"`
	Archived       bool        `json:"archived"`
	CreatedAt      string      `json:"created_at"`
	CreatedBy      string      `json:"created_by"`
	ClientID       string      `json:"client_id"`
	CustomBitlinks []string    `json:"custom_bitlinks"`
	Tags           []string    `json:"tags"`
	LaunchpadIds   []string    `json:"launchpad_ids"`
	Deeplinks      []Deeplinks `json:"deeplinks"`
}
type BitlinkHistory struct {
	UUID         string `json:"uuid"`
	GroupGUID    string `json:"group_guid"`
	Keyword      string `json:"keyword"`
	Bsd          string `json:"bsd"`
	Hash         string `json:"hash"`
	Login        string `json:"login"`
	LongURL      string `json:"long_url"`
	Created      string `json:"created"`
	FirstCreated string `json:"first_created"`
	Deactivated  string `json:"deactivated"`
	IsActive     string `json:"is_active"`
}

func (r *CustomBitlinkRequest) marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *UpdateCustomBitlinkRequest) marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *CustomBitlink) marshal() ([]byte, error) {
	return json.Marshal(r)
}

func unmarshalCustomBitlink(data []byte) (CustomBitlink, error) {
	var r CustomBitlink
	err := json.Unmarshal(data, &r)
	return r, err
}
