// Package store defines the UpdateStore interface for version persistence.
package store

import (
	"github.com/briandorsey/partisci/version"
	"time"
)

// UpdateStore defines an interface for persisting application version information.
type UpdateStore interface {
	Update(v version.Version) (err error)
	App(AppId string) (as version.AppSummary, ok bool)
	Apps() (as []version.AppSummary, err error)
	Host(Host string) (hs version.HostSummary, ok bool)
	Hosts() (hs []version.HostSummary, err error)
	Versions(AppId string, Host string, Ver string) (
		vs []version.Version, err error)
	Clear() (err error)

	Trim(t time.Time) (c uint64, err error)
}
