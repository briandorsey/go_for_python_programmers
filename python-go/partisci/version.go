// Package version provides types and functions for manipulating version data.
package version

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// START VERSION OMIT
type Version struct {
	AppID       string    `json:"app_id,omitempty"`  // HLapp
// END VERSION OMIT
	App         string    `json:"app"`
	Ver         string    `json:"ver"`
	Host        string    `json:"host"`
	Instance    uint16    `json:"instance"`
	HostIP      string    `json:"host_ip,omitempty"`
	LastUpdate  int64     `json:"last_update,omitempty"`
	ExactUpdate time.Time `json:"-"`
}

// Key returns a suitable unique id for storing in a database.
// It is calculated using AppID, Host & Instance so later version changes
// will result in updates.
func (v *Version) Key() string {
	return v.AppID + v.Host + fmt.Sprintf("%d", v.Instance)
}

// Prepare readies a Version for use by calculating fields.
// Prepare *must* be called after populating fields and before passing to a store.
func (v *Version) Prepare() {
	v.AppID = appIDToID(v.App) // HL
	if v.LastUpdate == 0 {  // OMIT
		v.ExactUpdate = time.Now()  // OMIT
		v.LastUpdate = v.ExactUpdate.Unix()  // OMIT
	}  // OMIT
}

type AppSummary struct {
	AppID      string `json:"app_id"`
	App        string `json:"app"`
	LastUpdate int64  `json:"last_update"`
	HostCount  int32  `json:"host_count"`
}

type HostSummary struct {
	Host       string `json:"host"`
	LastUpdate int64  `json:"last_update"`
	AppCount   int32  `json:"app_count"`
}

func safeRunes(r rune) rune {
	if '0' <= r && r <= '9' {
		return r
	}
	if 'a' <= r && r <= 'z' {
		return r
	}
	return '_'
}

func appIDToID(app string) string {  // HL
    id := strings.ToLower(app)
	id = strings.Map(safeRunes, id)
	return id
}

// START JSON OMIT
func ParsePacket(host string, b []byte) (v Version, err error) {
	v = *new(Version)
	v.HostIP = host
	err = json.Unmarshal(b[:len(b)], &v)
	if err != nil {
		return
	}
// END JSON OMIT
	v.Prepare()

	// ensure minimal values were given
	if len(v.App) == 0 ||
		len(v.Ver) == 0 {
		err = errors.New("value for app & ver must be specified")
		return
	}
	return
}
