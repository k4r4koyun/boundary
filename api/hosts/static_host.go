// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2020-05-01 09:51:58.7285456 -0400 EDT m=+0.014777201
package hosts

import (
	"encoding/json"
	"time"

	"github.com/fatih/structs"
	"github.com/hashicorp/watchtower/api/internal/strutil"
)

type StaticHost struct {
	defaultFields []string

	// Canonical path of the resource from the API's base URI
	// Output only.
	Path *string `json:"path,omitempty"`
	// Friendly name, if set
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The time this host was created
	// Ouput only.
	CreatedTime time.Time `json:"created_time,omitempty"`
	// The time this host was last updated
	// Output only.
	UpdatedTime time.Time `json:"updated_time,omitempty"`
	// Whether the host is disabled
	Disabled *bool `json:"disabled,omitempty"`
	// The address (DNS or IP name) used to reach the host
	Address *string `json:"address,omitempty"`
}

func (s *StaticHost) SetDefault(key string) {
	s.defaultFields = strutil.AppendIfMissing(s.defaultFields, key)
}

func (s *StaticHost) UnsetDefault(key string) {
	s.defaultFields = strutil.StrListDelete(s.defaultFields, key)
}

func (s StaticHost) MarshalJSON() ([]byte, error) {
	m := structs.Map(s)
	if m == nil {
		m = make(map[string]interface{})
	}
	for _, k := range s.defaultFields {
		m[k] = nil
	}
	return json.Marshal(m)
}