package models

import (
	"encoding/json"
	"time"
)

// Settings contains the elements to update the DNS record
// nolint: maligned
type Settings struct {
	Domain      string
	Host        string
	Provider    Provider
	IPMethod    IPMethod
	IPVersion   IPVersion
	Delay       time.Duration
	NoDNSLookup bool
	// Provider dependent fields
	Password       string // Namecheap, Infomaniak, DDNSS and NoIP only
	Key            string // GoDaddy, Dreamhost and Cloudflare only
	Secret         string // GoDaddy only
	Token          string // Cloudflare and DuckDNS only
	Email          string // Cloudflare only
	UserServiceKey string // Cloudflare only
	ZoneIdentifier string // Cloudflare only
	Identifier     string // Cloudflare only
	Proxied        bool   // Cloudflare only
	TTL            uint   // Cloudflare only
	Username       string // NoIP, Infomaniak, DDNSS only
}

func (settings *Settings) String() string {
	b, _ := json.Marshal(
		struct {
			Domain   string `json:"domain"`
			Host     string `json:"host"`
			Provider string `json:"provider"`
		}{
			settings.Domain,
			settings.Host,
			string(settings.Provider),
		},
	)
	return string(b)
}

// BuildDomainName builds the domain name from the domain and the host of the settings
func (settings *Settings) BuildDomainName() string {
	switch settings.Host {
	case "@":
		return settings.Domain
	case "*":
		return "any." + settings.Domain
	default:
		return settings.Host + "." + settings.Domain
	}
}
